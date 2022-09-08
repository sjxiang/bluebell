package logic

import (

	"golang.org/x/crypto/bcrypt"

	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/models"
	"github.com/sjxiang/bluebell/pkg/snowflake"
	"github.com/sjxiang/bluebell/requests"
)

// 存放业务逻辑的代码，拼装


func Signup(p *requests.ParamSignup) (err error) {
	
	// 1. 判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. 生成 uid
	userID := snowflake.GetID()
	
	// 3. 密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)  // 参数 2，cost 值，建议大于 12，数值越大耗费时间越长
	if err != nil {
		return
	}

	// 4. 保存进数据库
	u := models.User{  // 构造 1 个 user 实例
		UserID: userID,
		Username: p.Username,
		Password: string(hash),
	}

	err = mysql.InsertUser(&u)

	return
}


func Login(p *requests.ParamLogin) (err error) {
	
	// 1. 查询请求登录的用户
	user := models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.QueryUserByUsername(&user); err != nil {
		return err
	}

	// 2. 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password)); err != nil {
		return err
	}


	// 3. JWT
	return nil
}