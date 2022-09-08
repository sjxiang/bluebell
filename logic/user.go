package logic

import (
	"github.com/sjxiang/bluebell/dao/mysql"
	// "github.com/sjxiang/bluebell/models"
	// "github.com/sjxiang/bluebell/pkg/snowflake"
	"github.com/sjxiang/bluebell/requests"
)

// 存放业务逻辑的代码，拼装


func Signup(p *requests.ParamSignup) (err error) {
	
	// 1. 判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 生成 uid
	// userID := snowflake.GetID()
	
	// 构造 1 个 user 实例
	// u := &models.User{
	// 	UserID: userID,
	// 	Username: p.Username,
	// 	Password: p.Password,
	// }


	// 3. 保存进数据库
	// err = mysql.InsertUser(u)

	return
}