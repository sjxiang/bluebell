package mysql

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/sjxiang/bluebell/models"
)

// 把每一步数据库操作封装成函数
// 待 logic 层根据业务需求调用

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) bool {

	// var users []models.User
	// DB.Model(user).Select("user_id").Where("username = ?", username).Count(&count)

	var count int64

	DB.Table("user").Select("user_id").Where("username = ?", username).Count(&count)

	return !(count > 0)
}


// InsertUser 向数据库中插入一条新的记录
func InsertUser(user *models.User) bool {

	DB.Create(user)
	return user.ID > 0
}

// 查询 username 所属的用户
func QueryUserByUsername(user *models.User) bool {

	DB.Where("username = ?", user.Username).First(user)

	return user.ID > 0
}


// 
func ComparePassword(plainText, hash string) bool {
	return _hash.BcryptCheck(plainText, hash)
}


// err := bcrypt.CompareHashAndPassword([]byte(plainText), []byte(hash))
	
// 	if err != nil {
// 		zap.L().Error("密码错误 hash BcryptCheck", zap.Error(err))
// 		return false
// 	}
	
// 	return true 


