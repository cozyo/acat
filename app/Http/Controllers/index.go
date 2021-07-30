package Controllers

import (
	"github.com/acat/app/Http"
	"github.com/acat/app/Http/Validates"
	"github.com/acat/app/Respsitory/captcha"
	"github.com/acat/app/Service/Admin"
	"github.com/acat/core/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取验证码
func GetVerifyCodeApi(c *gin.Context)  {
	var (
		lenCode  = 4
		width    = 90
		height   = 40
	)

	// 保存session
	sessionName, imgByte := captcha.CreateCode(lenCode, width, height)
	if err := captcha.SetCode(c, sessionName); err != nil {
		msgData := make(map[string]interface{})
		msgData["err"] = err
		msgData["msg"] = "初始化session不成功"

		Http.JsonResponse(c, http.StatusOK, code.ErrorImageCaptch, msgData)
		return
	}

	c.Header("Content-Type", "image/png")
	if _, err := c.Writer.WriteString(string(imgByte)); err != nil {
		msgData := make(map[string]interface{})
		msgData["err"] = err
		msgData["msg"] = "图形验证码生成不正确"

		Http.JsonResponse(c, http.StatusOK, code.ErrorImageCaptch, msgData)
		return
	}
}

// 登录接口
func LoginAdminApi(c *gin.Context)  {
	var (
		form Validates.LoginUserCodeArgs
		err  error
	)
	// 验证入参是否正确
	if err = Http.BindAndValid(c, &form); err != nil {
		Http.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	if token, ok := Admin.LoginUserWithVerifyCode(c, &form); ok {
		c.Header("token", token)
		Http.JsonResponse(c, http.StatusOK, code.SUCCESS, "token")
		return
	}

	Http.JsonResponse(c, http.StatusOK, code.ERROR, "验证码不正确")
}