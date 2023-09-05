package main

import (
	"manager-gin/src/global"
	"manager-gin/src/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	done := make(chan bool)
	go func() {
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-osSignals:
			global.Logger.Info("系统信号中断")
			// 关闭mongo
			global.CloseMongoClient()
			// 关闭日志
			global.CloseLogger()
			done <- true
		}
	}()
	server.Run()
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-done:
			global.Logger.Info("系统退出")
			return
		default:
			continue
		}
	}
}
