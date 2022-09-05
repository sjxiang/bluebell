package main

import (
	"fmt"

	"github.com/sjxiang/bluebell/conf"
	"github.com/spf13/viper"
)


func main() {
	
	// 1. 加载配置 
	if err := conf.Init(); err != nil {
		fmt.Printf("init conf failed, err:%v\n", err)
	}

	fmt.Println(viper.GetString("app.name"))
	
	
	// 2. 初始化日志
	// 3. 初始化 MySQL 连接
	// 4. 初始化 Redis 连接
	// 5. 注册路由
	// 6. 启动服务（优雅关机）
}

