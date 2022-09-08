package validators


// ValidatePasswordConfirm 自定义规则，检查两次密码是否一样
func ValidatePasswordConfirm(password, PasswordConfirm string, errs map[string][]string) map[string][]string {

	if password != PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配")
	}

	return errs
}


// TODO『手机、邮箱验证码』