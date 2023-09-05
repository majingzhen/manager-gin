package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"manager-gin/src/global"
	"manager-gin/src/routers"
)

func init() {
	global.InitViper()
	global.InitLogger()
	global.InitDataSource()
	global.InitMongoClient()
}

func Run() {
	gin.SetMode(global.Viper.GetString("server.model"))
	// 关闭日志颜色
	gin.DisableConsoleColor()
	router := new(routers.Routers).InitRouter()
	_ = router.Run(fmt.Sprintf(":%d", global.Viper.GetInt("server.port")))
}
