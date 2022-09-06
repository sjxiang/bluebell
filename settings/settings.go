package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)


func Init() (err error) {
	
	viper.SetConfigName("config")       // 指定配置文件名称
	viper.SetConfigType("yaml")         // 指定配置文件类型
	viper.AddConfigPath("./settings")   // 指定查找配置文件的路径 （相对路径）
	err = viper.ReadInConfig()      // 读取配置信息

	if err != nil {  // 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改了 ... ")
	})

	return
}


/*

示例

	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
	}

	fmt.Println(viper.GetString("app.name"))


缺陷
	写错了，就是空 ""，很难发现错误

*/

