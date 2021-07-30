package captcha

import (
	"github.com/acat/utils/helper"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
)

var captchaSessionName = "ACAT-CAPTCHA"

func init() {
	ranTag := helper.ShortTag(helper.GetRandStr(2), 1)
	captchaSessionName += ranTag
}

// 校验验证码是否正确
func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	codeMd5 := helper.Md5ToString(strings.ToUpper(code))
	if sessionCode := session.Get(captchaSessionName); codeMd5 == sessionCode {
		session.Delete(captchaSessionName)
		return true
	}
	return false
}

func CreateCode(n int, size ...int) (text string, imgByte []byte) {
	var (
		width  = 180
		height = 60
	)

	text = helper.GetRandStr(n)
	textMd5 := helper.Md5ToString(strings.ToUpper(text))

	if len(size) >= 2 {
		width = size[0]
		height = size[1]
	}

	imgByte = ImgText(width, height, text)

	return textMd5, imgByte
}

// 设置code并储存至session中
func SetCode(c *gin.Context, sessionName string) error {
	session := sessions.Default(c)
	session.Set(captchaSessionName, sessionName)
	//记着调用save方法，写入session
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
