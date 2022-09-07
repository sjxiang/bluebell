package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/sjxiang/bluebell/dao/mysql"
	"github.com/sjxiang/bluebell/dao/redis"
	"github.com/sjxiang/bluebell/logger"
	"github.com/sjxiang/bluebell/routes"
	"github.com/sjxiang/bluebell/settings"
)


func main() {
	
	// 1. 加载配置 
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()

	// 3. 初始化 MySQL 连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化 Redis 连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	// 5. 注册路由
	r := routes.Setup()

	// 6. 启动服务
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Conf.AppConfig.Port),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		zap.L().Fatal("监听端口 8081")
	}

}

