package vars

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Version    = "1.0.0"
	AppTypeWeb = 1
)

// Application ...
type Application struct {
	Name         string
	Type         int32
	LoadConfig   func() error
	SetupVars    func() error
	StopFunc     func() error
}

// WebApplication ...
type WebApplication struct {
	*Application
	Port           int
	// 监控
	Mux *http.ServeMux
	// RegisterHttpRoute 定义HTTP router
	RegisterHttpRoute func() *gin.Engine
	// 系统定时任务
	RegisterTasks func() []CronTask
}

// ListenerApplication ...
type ListenerApplication struct {
	*Application
	Port           int
	NetWork        int
	ReadTimeOut    int
	WriteTimeOut   int
	// 监控
	Mux *http.ServeMux
	// RegisterHttpRoute 定义HTTP router
	EventHandler func(context.Context, []byte) ([]byte, error)
}
