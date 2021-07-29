package bootstrap

import (
	"fmt"
	"gitee.com/kelvins-io/common/log"
	"github.com/acat/core/vars"
	"github.com/acat/utils/logging"
	"os"
	"time"
)

const (
	DefaultLoggerRootPath = "./logs"
	DefaultLoggerLevel = "debug"
)

// 初始化application--日志部分
func InitApplication(application *vars.Application) error {
	var (
		rootPath = DefaultLoggerRootPath
		loggerLevel = DefaultLoggerLevel
	)

	if vars.LoggerSetting != nil && vars.LoggerSetting.RootPath != ""{
		rootPath = vars.LoggerSetting.RootPath
	}
	if vars.LoggerSetting != nil && vars.LoggerSetting.Level != "" {
		loggerLevel = vars.LoggerSetting.Level
	}

	if err := log.InitGlobalConfig(rootPath, loggerLevel, application.Name); err != nil {
		return fmt.Errorf("log.InitGlobalConfig: %v", err)
	}
	return nil
}

func AppShutDown(application *vars.Application) error {
	if application.StopFunc != nil {
		return application.StopFunc()
	}
	return nil
}

func AppPrepareForceExit()  {
	time.AfterFunc(10 * time.Second, func() {
		logging.Info("App server Shutdown timeout, force exit")
		os.Exit(1)
	})
}

// SetupCommonVars ...
func SetupCommonVars() error {
	if vars.ServerSetting != nil {
		vars.App.Port = vars.ServerSetting.Port
	}
	return nil
}