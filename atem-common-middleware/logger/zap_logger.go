package logger

import (
	"fmt"
	"github/socloudng/atem-common/atem-common-middleware/fsystem"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap(cfg *ZapConfig) (logger *zap.Logger) {
	if ok, _ := fsystem.PathExists(cfg.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", cfg.Director)
		_ = os.Mkdir(cfg.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", cfg.Director), debugPriority, cfg),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", cfg.Director), infoPriority, cfg),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", cfg.Director), warnPriority, cfg),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", cfg.Director), errorPriority, cfg),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if cfg.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig(cfg *ZapConfig) (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  cfg.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoderFunc(cfg.Prefix),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case cfg.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case cfg.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case cfg.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case cfg.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler, cfg *ZapConfig) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName, cfg.LogInConsole) // 使用file-rotatelogs进行日志分割
	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(getEncoderConfig(cfg))
	} else {
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig(cfg))
	}
	return zapcore.NewCore(encoder, writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoderFunc(prefix string) func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(prefix + "2006/01/02 - 15:04:05.000"))
	}
}
