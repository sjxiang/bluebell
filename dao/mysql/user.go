package mysql

import (

	"go.uber.org/zap"
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

	return !(count > 0)  // Tips：取反
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




// ===

// 明文密码加密
func BcryptPassword(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	if err != nil {
		zap.L().Error("加密失败_mysql_BcryptPassword", zap.Error(err))  // 参数 2 为 cost 值，建议大于 12，数值越大耗费时间越长
		return "", err
	}
	return string(hash), nil 
} 


// 校验明文密码和哈希值
func ComparePassword(plainText, hash string) bool {
	// 前哈希值，后明文密码
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainText)); err != nil {
		zap.L().Error("密码错误_mysql_ComparePassword", zap.Error(err))
		return false
	}
	return true
}