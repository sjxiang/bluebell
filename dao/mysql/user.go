package mysql

import (
	"errors"

	"github.com/sjxiang/bluebell/models"
)

// 把每一步数据库操作封装成函数
// 待 logic 层根据业务需求调用

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {

	// var users []models.User
	// DB.Model(user).Select("user_id").Where("username = ?", username).Count(&count)

	var count int64

	err = DB.Table("user").Select("user_id").Where("username = ?", username).Count(&count).Error
	if err != nil {
		// 数据库查询出错
		return err
	}
	if count > 0 {
		// 用户已存在的错误
		return errors.New("用户已存在")
	}

	return nil
}


// InsertUser 向数据库中插入一条新的记录
func InsertUser(user *models.User) (err error) {

	// 执行 SQL 入库
	err = DB.Create(user).Error
	if err != nil {
		
		// 创建失败
		return  
	}

	return nil  
}


func QueryUserByUsername(user *models.User) (err error) {
	
	err = DB.Where("username = ?", user.Username).Find(user).Error

	if err != nil {
		return 
	}

	return nil 
}

