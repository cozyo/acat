package Controllers

import (
	"github.com/acat/app/Http"
	"github.com/acat/core/code"
	"github.com/acat/core/vars"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context)  {
	Http.JsonResponse(c, http.StatusOK, code.SUCCESS, "Welcome to " + vars.App.Name)
}

func PingApi(c *gin.Context)  {
	Http.JsonResponse(c, http.StatusOK, code.SUCCESS, "")
}
