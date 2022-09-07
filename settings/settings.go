package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)


// Conf 全局变量，用来保存应用的所有配置信息
var Conf = new(Config)


type Config struct {
	*AppConfig         `mapstructure:"app"`
	*LogConfig         `mapstructure:"log"`
	*MySQLConfig       `mapstructure:"mysql"`
	*RedisConfig       `mapstructure:"redis"`
}


type AppConfig struct {
	Name   		string `mapstructure:"name"`
	Mode   		string `mapstructure:"mode"`
	Port    	int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	MachineID   string `mapstructure:"machine_id"`
	StartTime   string `mapstructure:"start_time"`
}


type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}


type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}


type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}



func Init() (err error) {
	
	viper.SetConfigFile("./settings/config.yaml")
	
	if err = viper.ReadInConfig(); err != nil {  
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	if err = viper.Unmarshal(Conf); err != nil {
		// 反序列化到结构体，失败
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改了 ... ")
		
		if err = viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
			return
		}

	})

	return
}

