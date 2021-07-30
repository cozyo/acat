package Validates

import "github.com/astaxie/beego/validation"

type LoginUserCodeArgs struct {
	UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	VerifyCode  string `form:"verify_code" json:"verify_code"`
}

func (l *LoginUserCodeArgs) Valid(v *validation.Validation) {
	if l.UserName == "" {
		v.SetError("UserName", "请输入后台账号")
		return
	} else if len(l.UserName) < 2 {
		v.SetError("UserName", "账号名不能少于2位")
		return
	}

	if l.Password == "" {
		v.SetError("Password", "请输入密码")
		return
	} else if len(l.Password) < 6 {
		v.SetError("Password", "密码长度不能少于6位")
		return
	}
	if l.VerifyCode == "" {
		v.SetError("VerifyCode", "请输入验证码")
		return
	} else if len(l.VerifyCode) != 4 {
		v.SetError("VerifyCode", "验证码长度错误")
		return
	}
}