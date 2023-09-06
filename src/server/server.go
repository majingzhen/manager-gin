package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"manager-gin/src/global"
	"manager-gin/src/routers"
	"manager-gin/src/utils"
	"net/http"
	"os"
	"os/signal"
	"time"
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

	port := global.Viper.GetInt("server.port")
	// 创建一个HTTP服务
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.Logger.Error("系统启动失败", zap.Error(err))
		}
		global.Logger.Info(fmt.Sprintf("系统启动成功,服务运行在 http://%s:%d", utils.GetIp(), port))
	}()

	// 创建一个信号接收器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// 等待接收到关闭信号
	<-quit
	log.Println("Shutting down server...")

	// 创建一个上下文对象
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭HTTP服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	global.CloseLogger()
	log.Println("Server exiting")
}
