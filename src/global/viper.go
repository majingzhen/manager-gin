package global

import (
	"github.com/spf13/viper"
)

// InitViper 初始化配置
func InitViper() {
	Viper = viper.New()
	Viper.AddConfigPath(".")           // 添加配置文件搜索路径，点号为当前目录
	Viper.AddConfigPath("./config")    // 添加多个搜索目录
	Viper.SetConfigType("yml")         // 如果配置文件没有后缀，可以不用配置
	Viper.SetConfigName("application") // 文件名，没有后缀
	// 读取配置文件
	if err := Viper.ReadInConfig(); err != nil {
		panic("读取配置文件错误")
	}
}
