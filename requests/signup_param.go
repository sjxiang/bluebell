package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)



type ParamSignup struct {
	Username        string `json:"username"         valid:"username"` 
	Password        string `json:"password"         valid:"password"`
	PasswordConfirm string `json:"password_confirm" valid:"password_confirm"`
}





func Signup(data interface{}, ctx *gin.Context) map[string][]string {

	// 1. 自定义规则
	rules := govalidator.MapData{
		"username": []string{"required", "alpha_num", "between:3,20"},
		"password": []string{"required", "min:6"},
		"password_confirm": []string{"required", "min:6"},		
	}

	// 2. 自定义验证出错时的提示
	messages := govalidator.MapData{
		"username": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码为必填项",
		},
	}

	// 3. 开始验证
	errs := validate(data, rules, messages)
	return errs
}

