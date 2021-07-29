package main

import (
	"github.com/acat/app"
	"github.com/acat/core/startup"
	"github.com/acat/core/vars"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

const AppName = "micro-gm-api"

func main()  {
	application := &vars.WebApplication{
		Application : &vars.Application{
			Name: AppName,
			LoadConfig: startup.LoadConfig,
			SetupVars: startup.SetupVars,
			StopFunc: startup.SetStopFunc,

		},
		RegisterHttpRoute: startup.RegisterHttpRoute,
		RegisterTasks: startup.RegisterTasks,
	}
	app.Run(application)
}
