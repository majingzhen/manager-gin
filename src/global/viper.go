package global

import (
	"github.com/spf13/viper"
)

// InitViper 初始化配置
func InitViper() {
	GVA_VP = viper.New()
	GVA_VP.AddConfigPath(".")           // 添加配置文件搜索路径，点号为当前目录
	GVA_VP.AddConfigPath("./config")    // 添加多个搜索目录
	GVA_VP.SetConfigType("yml")         // 如果配置文件没有后缀，可以不用配置
	GVA_VP.SetConfigName("application") // 文件名，没有后缀
	// v.SetConfigFile("configs/app.yml")
	// 读取配置文件
	if err := GVA_VP.ReadInConfig(); err == nil {
		Logger.Info("use config file -> " + GVA_VP.ConfigFileUsed())
	}
}
