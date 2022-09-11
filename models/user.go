package models


// 用户表
type User struct {
	Base 
	
	UserID   int64  `gorm:"type:varchar(64);not null"`
	Username string `gorm:"type:varchar(64);not null;UNIQUE"`
	Password string `gorm:"type:varchar(64)"`
	Email    string `gorm:"type:varchar(64);UNIQUE"`
	Gender   int

	CommonTimestampsField
}

