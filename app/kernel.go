package app

import (
	"context"
	"fmt"
	"github.com/acat/bootstrap"
	"github.com/acat/core/config"
	"github.com/acat/core/util/kprocess"
	"github.com/acat/core/vars"
	"github.com/acat/utils/logging"
	"github.com/robfig/cron/v3"
	"net/http"
	"os"
	"strconv"
	"time"
)

const localAddr = "0.0.0.0:"

func Run(application *vars.WebApplication)  {
	if application.Name == "" {
		logging.Fatal("Application name can't not be empty")
	}

	application.Type = vars.AppTypeWeb
	vars.App = application
	if err := RunApplication(application); err != nil {
		logging.Fatalf("App exit over: %v\n", err)
	}
	logging.Info("App exit over")
}

func RunApplication(Web *vars.WebApplication) error {
	// 1. load config
	var err error
	if err := config.LoadDefaultConfig(); err != nil {
		return nil
	}
	if Web.LoadConfig != nil {
		if err := Web.LoadConfig(); err != nil {
			return err
		}
	}
	// 2. init application
	if err := bootstrap.InitApplication(Web.Application); err != nil {
		return err
	}
	// 3. setup vars
	if err := setupWebVars(); err != nil {
		return err
	}
	if Web.SetupVars != nil {
		if err := Web.SetupVars(); err != nil {
			return fmt.Errorf("App.SetupVars err: %v", err)
		}
	}
	//4.  setup server monitor


	// 5 run task
	if Web.RegisterTasks != nil {
		cronTask := Web.RegisterTasks()
		if len(cronTask) != 0 {
			cn := cron.New(cron.WithSeconds())
			for i := 0; i < len(cronTask); i++ {
				if cronTask[i].TaskFunc != nil {
					_, err = cn.AddFunc(cronTask[i].Cron, cronTask[i].TaskFunc)
					if err != nil {
						logging.Fatalf("App run cron task err: %v", err)
					}
				}
			}
			cn.Start()
			logging.Info("App run cron task")
		}
	}

	// 6. set init service port
	var (
		addr     string
		pidFile  string
		network = "tcp"
	)

	if Web.Port != 0 {
		addr = localAddr + strconv.Itoa(Web.Port)
	} else if vars.ServerSetting.Port != 0 {
		addr = localAddr + strconv.Itoa(vars.ServerSetting.Port)
	}
	// 7. run http server
	if Web.RegisterHttpRoute == nil {
		logging.Fatalf("App RegisterHttpRoute nil !!")
	}
	wd, _ := os.Getwd()

	if vars.ServerSetting.PIDFile != "" {
		pidFile = vars.ServerSetting.PIDFile
	} else {
		pidFile = fmt.Sprintf("%s/%s.pid", wd, Web.Name)
	}

	kp := new(kprocess.KProcess)
	if vars.ServerSetting != nil && vars.ServerSetting.NetWork != "" {
		network = vars.ServerSetting.NetWork
	}
	ln, err := kp.Listen(network, addr, pidFile)
	if err != nil {
		logging.Fatalf("App kprocess listen err: %v", err)
	}
	ginEngine := Web.RegisterHttpRoute()
	serve := http.Server{
		Handler: ginEngine,
		ReadTimeout: time.Duration(vars.ServerSetting.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(vars.ServerSetting.WriteTimeout) * time.Second,
		IdleTimeout: time.Duration(vars.ServerSetting.IdleTimeout) * time.Second,
	}
	go func() {
		if err := serve.Serve(ln); err != nil {
			logging.Fatalf("App run Serve: %v", err)
		}
	}()
	<-kp.Exit()

	bootstrap.AppPrepareForceExit()
	if err := serve.Shutdown(context.Background()); err != nil {
		logging.Fatalf("App server Shutdown: %v", err)
	}
	logging.Info("App server Shutdown ok")
	err = bootstrap.AppShutDown(Web.Application)

	return err
}

// setupWebVars ...
func setupWebVars() error {
	if err := bootstrap.SetupCommonVars(); err != nil {
		return err
	}
	return nil
}

