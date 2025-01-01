package main

import (
	"context"

	"github.com/winterant/wlog"
)

func main() {
	ctx := wlog.ContextWithArgs(context.Background(), "taskId", "tsk-thisisataskid", "tag", "mytag")
	wlog.Debug(ctx, "(acquiescent myslog.Logger)process is starting...")
	wlog.Info(ctx, "My name is %s.", "Winterant")
}
