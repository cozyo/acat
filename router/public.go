package router


import (
	"github.com/acat/app/Http/Controllers"
	"github.com/gin-gonic/gin"
)

func PublicRouter(r *gin.Engine)  {
	r.GET("/", Controllers.IndexApi)  // home
	r.GET("/ping", Controllers.PingApi) // ping
	r.GET("/image/verify_code", Controllers.GetVerifyCodeApi)  // 获取图形验证码
	r.POST("/admin/login", Controllers.LoginAdminApi)  // 后台登录


	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}