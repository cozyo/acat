package router

import "github.com/gin-gonic/gin"

func WebRouter(r *gin.Engine)  {
	r.GET("/test_web", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "web_test",
		})
	})
}
