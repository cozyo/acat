package route

import (
	"github.com/acat/app/Http/Middleware"
	"github.com/acat/core/vars"
	"github.com/acat/router"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func InitRouter(accessInfoLogger, accessErrLogger io.Writer) *gin.Engine {

	gin.DefaultWriter = io.MultiWriter(os.Stdout, accessInfoLogger)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, accessErrLogger)

	if vars.ServerSetting == nil || vars.ServerSetting.Mode == "" {
		// 默认生产
		vars.ServerSetting.Mode = gin.ReleaseMode
	}
	gin.SetMode(vars.ServerSetting.Mode)

	r := gin.New()
	r.Use(Middleware.Cors())
	boot(r)

	log.Println("监听端口", "http://127.0.0.1:52001")

	return r
}

// boot Define your route model
func boot(Route *gin.Engine)  {
	router.PublicRouter(Route)
	router.WebRouter(Route)
	router.ApiRouter(Route)
}