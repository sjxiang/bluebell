package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/schema"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"go.uber.org/zap"

	"github.com/sjxiang/bluebell/settings"
)

var DB *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn :=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&&multiStatements=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,  // 解决单数表名，user
		},
		Logger: logger.Default.LogMode(logger.Info),  // SQL 语句输出
	})

	if err != nil {
		zap.L().Error("连接 DB 失败", zap.Error(err))
		return
	}

	// 连接池配置
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := DB.DB()
	if err != nil {
		zap.L().Error("获取 sql.DB 失败", zap.Error(err))
		return
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(cfg.MaxOpenConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(cfg.MaxIdleConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	migration()

	return nil
}