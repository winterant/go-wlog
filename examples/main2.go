package main

import (
	"context"
	"io"
	"log/slog"
	"os"

	"github.com/winterant/wlog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	// 自行指定日志输出目标
	writers := io.MultiWriter(&lumberjack.Logger{
		Filename:   "./log/main.log", // 日志文件的位置
		MaxSize:    128,              // 文件最大大小（单位MB）
		MaxBackups: 0,                // 保留的最大旧文件数量
		MaxAge:     90,               // 保留旧文件的最大天数
		Compress:   false,            // 是否压缩/归档旧文件
		LocalTime:  true,             // 使用本地时间创建时间戳
	}, os.Stdout)
	wlog.InitDefaultLogger(writers, slog.LevelDebug)
}

func main() {
	ctx := context.Background()
	ctx = wlog.ContextWithArgs(ctx, "taskId", "tsk-thisisataskid", "tag", "mytag") // 利用context确保每一条都输出某些信息
	wlog.Debug(ctx, "process is starting...")
	wlog.Info(ctx, "My name is %s.", "Winterant")
}
