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
