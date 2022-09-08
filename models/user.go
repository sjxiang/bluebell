package models


import (
	"time"
)

type Base struct {
	ID       uint64     `gorm:"column:id;primaryKey;autoIncrement"` 
}

// 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`
}


// 用户表
type User struct {
	Base 
	
	UserID   int64  `gorm:"type:varchar(64);not null"`
	Username string `gorm:"type:varchar(64);not null"`
	Password string `gorm:"type:varchar(64)"`
	Email    string `gorm:"type:varchar(64)"`
	Gender   int

	CommonTimestampsField
}

