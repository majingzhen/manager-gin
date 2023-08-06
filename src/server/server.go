package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"manager-gin/src/global"
	"manager-gin/src/routers"
	"os"
)

func init() {
	global.InitLogger()
	global.InitViper()
	global.InitDataBase()
}

func Run() {
	gin.SetMode(global.GVA_VP.GetString("server.model"))
	// 关闭日志颜色
	gin.DisableConsoleColor()
	log, _ := os.Create(global.GVA_VP.GetString("logger.file_path"))
	gin.DefaultWriter = io.MultiWriter(log)
	router := new(routers.Routers).InitRouter()

	_ = router.Run(fmt.Sprintf(":%d", global.GVA_VP.GetInt("server.port")))
}
