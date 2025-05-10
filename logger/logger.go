package logger

import (
	"log/slog"
)

var Log = slog.Default()

func init() {
	// todo 从环境变量获取日志输出级别
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func Error(msg string, err error) {

	Log.Error(msg, "err", err.Error())
}

func Info(msg string, args ...interface{}) {
	Log.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	Log.Debug(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	Log.Warn(msg, args...)
}
