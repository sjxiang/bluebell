package main

import (
	"fmt"
	"net/http"

	"github.com/sjxiang/bluebell/logger"
	"github.com/sjxiang/bluebell/middleware"
	"github.com/sjxiang/bluebell/pkg/snowflake"
	"github.com/sjxiang/bluebell/settings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)


func main() {
	
	// 1. 加载配置 
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
	}

	fmt.Println(viper.GetString("app.name"))
	
	if err := snowflake.Init(viper.GetString("app.startTime"), viper.GetInt64("app.machineID")); err != nil {
		fmt.Printf("init pkg snowflake failed, err:%v\n", err)
	}

	id := snowflake.GetID()
	fmt.Println(id)

	// 2. 初始化日志
	if err := logger.Init(); err != nil {
		zap.L().Debug("Logger init success")
	}
	// 3. 初始化 MySQL 连接
	// 4. 初始化 Redis 连接
	// 5. 注册路由
	// 6. 启动服务（优雅关机）
	zap.L().Debug("Logger init success")


	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery())

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Msg": "pong",
		})
	})

	r.Run(":8081")
}

