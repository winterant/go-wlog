# go-wslog

---

`go-wlog` is a logging library based on `slog` that adds a Handler capable of printing logs in a format that is easy to
browse.

# 🔨 Installation

```bash
go get -u github.com/winterant/wlog
```

## 🪤 Examples

### Use default logger

The log will be output to `os.Stdout` and the file `./log/main.log`.

```go
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
```

log:

```
2024-10-02 12:21:32.365340 DEBUG /Users/jinglong/Projects/github/myslog/main.go:12 [taskId=tsk-thisisataskid] [tag=mytag] process is starting...
2024-10-02 12:21:32.365816 INFO  /Users/jinglong/Projects/github/myslog/main.go:15 [taskId=tsk-thisisataskid] [tag=mytag] My name is Winterant.
```

### Custom config

```go
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
```

log:

```
2024-10-02 11:42:17.227797 DEBUG /Users/jinglong/Projects/github/myslog/main.go:34 [taskId=tsk-thisisataskid] [tag=mytag] process is starting...
2024-10-02 11:42:17.228035 INFO  /Users/jinglong/Projects/github/myslog/main.go:37 [taskId=tsk-thisisataskid] [tag=mytag] My name is Winterant.
```

### Get slog.Logger

```go
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
```

log:

```
2024-10-01 21:05:59.713409 DEBUG /Users/jinglong/Projects/github/myslog/main.go:35 [key=display_in_each_log] [taskId=tsk-thisisatask] process is starting...
2024-10-01 21:05:59.714219 INFO  /Users/jinglong/Projects/github/myslog/main.go:38 [key=display_in_each_log] [taskId=tsk-thisisatask] [money=9999999] My name is Winterant.
```

## 🚛 Appendix

### filebeat日志收集配置

`filebeat.yaml`:

```yaml
filebeat.inputs:
  - type: log
    paths:
      - './log/*.log'
    multiline.pattern: '^\d{4}-\d{2}-\d{2}'
    multiline.negate: true
    multiline.match: after

processors:
  - drop_event:
      when:
        regexp:
          message: 'FILEBEAT_EXCLUDE'  # 排除包含FILEBEAT_EXCLUDE的日志

output.elasticsearch:
  hosts: [ "10.10.10.10:8200" ]
  username: "myusername"
  password: "mypassword"
  index: "my_project_online"

setup.ilm.enabled: false
setup.template.name: "my_project_online"
setup.template.pattern: "my_project_online-*"
setup.template.overwrite: false
```
