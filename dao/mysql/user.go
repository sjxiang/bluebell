package mysql

import (
	"errors"

	// "github.com/sjxiang/bluebell/models"
	// "golang.org/x/crypto/bcrypt"
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


// // InsertUser 向数据库中插入一条新的记录
// func InsertUser(user *models.User) (err error) {
// 	// 对密码加密
	
// 	// GenerateFromPassword 的第二个参数时 cost 值，建议大于 12，数值越大耗费时间越长
// 	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
// 	if err != nil {
// 		return
// 	}
	
// 	// 执行 SQL 语句入库
// 	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
// 	_, err = db.Exec(sqlStr, user.UserID, user.Username, string(hash))
// 	return 
// }


// func QueryUserByUsername() {

// }

