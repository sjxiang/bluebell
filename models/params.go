package models


// Param
type ParamSignup struct {
	Username        string `json:"username"         valid:"username"` 
	Password        string `json:"password"         valid:"password"`
	PasswordConfirm string `json:"password_confirm" valid:"password_confirm"`
}


