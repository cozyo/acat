package router


import (
	"github.com/acat/app/Http/Controllers"
	"github.com/gin-gonic/gin"
)

func PublicRouter(r *gin.Engine)  {
	r.GET("/", Controllers.IndexApi)
	r.GET("/ping", Controllers.PingApi)


	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}