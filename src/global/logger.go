package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// InitLogger 初始化日志
func InitLogger() {
	logPath := Viper.GetString("logger.file_path")
	if logPath == "" {
		logPath = "./log/manager.log" // 如果未配置日志路径，则默认在项目根目录下创建log目录
	}
	// 设置日志文件的位置、文件名、最大大小、最大备份数量和压缩
	hook := lumberjack.Logger{
		Filename:   logPath, // 日志路径
		MaxSize:    128,     // MB
		MaxBackups: 30,
		MaxAge:     7, // days
		Compress:   true,
	}
	// 配置日志级别
	atomicLevel := zap.NewAtomicLevel()
	logLevel := Viper.GetInt32("logger.level")
	atomicLevel.SetLevel(zapcore.Level(logLevel))
	// 创建编码器
	// 设置日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// 创建core
	writer := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		atomicLevel,
	).With([]zap.Field{})
	// 初始化logger
	Logger = zap.New(writer)
}
