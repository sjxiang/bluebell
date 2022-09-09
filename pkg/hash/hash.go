package hash

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(plainText string) (string, error) {

	// 参数 2 为 cost 值，建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)

	if err != nil {
		zap.L().Error("加密失败 hash BcryptHash", zap.Error(err))
		return "", err
	}

	return string(bytes), nil
}


// BcryptCheck 对比明文密码 plainText 和数据库的哈希值 hash
func BcryptCheck(plainText, hash string) bool {
	
	err := bcrypt.CompareHashAndPassword([]byte(plainText), []byte(hash))
	
	if err != nil {
		zap.L().Error("密码错误 hash BcryptCheck", zap.Error(err))
		return false
	}
	
	return true 
}

