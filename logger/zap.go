package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var lg *zap.Logger

// InitLogger 初始化lg
func InitLogger() (err error) {
	writeSyncer := getLogWriter("easy-cron-log", 200, 30, 7)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte("debug"))
	if err != nil {
		return
	}
	var core zapcore.Core

	// 进入开发模式，日志输出到终端
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	core = zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, l),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
	)

	lg = zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
