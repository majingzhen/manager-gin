package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var once sync.Once
var file *os.File
var ZpCore zapcore.Core

func InitLogger() {
	var err error
	file, err = os.Create(Viper.GetString("logger.file_path"))
	if err != nil {
		panic(err)
	}
	// 创建文件写入器
	fileWriteSyncer := zapcore.AddSync(file)

	// 创建console输出器
	consoleWriteSyncer := zapcore.AddSync(os.Stdout)

	// 配置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	// 创建编码器
	encoder := zap.NewDevelopmentEncoderConfig()
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 创建core
	writer := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(fileWriteSyncer, consoleWriteSyncer),
		atomicLevel,
	).With([]zap.Field{})
	// 初始化logger
	Logger = zap.New(writer)
	ZpCore = writer
	Logger.Sync()
}

// CloseLogger 关闭日志
func CloseLogger() {
	once.Do(func() {
		if file != nil {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}
	})
}
