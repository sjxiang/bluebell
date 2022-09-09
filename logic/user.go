package logic

import (
	"errors"

	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/models"
	"github.com/sjxiang/bluebell/pkg/snowflake"
	"github.com/sjxiang/bluebell/requests"
)


var (
	ErrorUserExist   = errors.New("用户已存在")
	ErrorUserNoExist = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorEncryptFailed = errors.New("加密失败")
)

// 存放业务逻辑的代码，拼装


func Signup(p *requests.ParamSignup) (err error) {
	
	// 1. 判断用户存不存在（或者是否已经注册）
	if ok := mysql.CheckUserExist(p.Username); !ok {
		return errors.New("用户已存在")
	}

	// 2. 生成 uid
	userID := snowflake.GetID()
	
	// 3. 密码加密
	hash, err := mysql.BcryptPassword(p.Password)
	if err != nil {
		return errors.New("加密失败")
	}

	// 4. 保存进数据库
	u := models.User{  // 构造 1 个 user 实例
		UserID: userID,
		Username: p.Username,
		Password: hash,
	}

	if ok := mysql.InsertUser(&u); !ok {
		return errors.New("创建失败")
	}

	return
}


func Login(p *requests.ParamLogin) (err error) {
	
	// 1. 查询请求登录的用户 `username`
	var user models.User
	user.Username = p.Username

	if ok := mysql.QueryUserByUsername(&user); !ok {
		return errors.New("用户不存在")
	}

	// 2. 判断密码是否正确 
	if ok := mysql.ComparePassword(p.Password, user.Password); !ok {
		return errors.New("密码错误")
	}


	// 3. JWT
	return nil
}

