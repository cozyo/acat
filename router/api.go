package router

import (
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.Engine)  {

	r.GET("/test_api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test_api",
		})
	})
}
