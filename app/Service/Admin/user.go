package Admin

import (
	"github.com/acat/app/Http/Validates"
	"github.com/acat/app/Respsitory/captcha"
	"github.com/gin-gonic/gin"
)

func LoginUserWithVerifyCode(c *gin.Context, req *Validates.LoginUserCodeArgs) (string, bool) {
	var token string

	if ok := captcha.CaptchaVerify(c, req.VerifyCode); !ok {
		return token, ok
	}
	return token, true
}