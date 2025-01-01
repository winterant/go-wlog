package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/winterant/wlog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func GetLogger() *slog.Logger {
	writers := io.MultiWriter(&lumberjack.Logger{
		Filename:   "./log/main.log", // 日志文件的位置
		MaxSize:    128,              // 文件最大大小（单位MB）
		MaxBackups: 0,                // 保留的最大旧文件数量
		MaxAge:     90,               // 保留旧文件的最大天数
		Compress:   false,            // 是否压缩/归档旧文件
		LocalTime:  true,             // 使用本地时间创建时间戳
	}, os.Stdout)

	handler := wlog.NewPrettyHandler(wlog.WithWriter(writers), wlog.WithLever(slog.LevelDebug))
	return slog.New(handler).With("key", "display_in_each_log")
}

func main() {
	slogger := GetLogger()
	ctx := wlog.ContextWithArgs(context.Background(), "taskId", "tsk-thisisatask")
	slogger.Log(ctx, slog.LevelDebug, "process is starting...")
	slogger.Log(ctx, slog.LevelInfo, fmt.Sprintf("My name is %s.", "Winterant"), "money", "9999999")
}
