package main

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"caas-micro/internal/app/api"
	"caas-micro/pkg/logger"
	"caas-micro/pkg/util"
)

// VERSION 版本号，
// 可以通过编译的方式指定版本号：go build -ldflags "-X main.VERSION=x.x.x"
var VERSION = "1.0.0"

var (
	configFile string
	swaggerDir string
)

func init() {
	//flag.StringVar(&configFile, "c", "", "配置文件(.json,.yaml,.toml)")
	//flag.StringVar(&swaggerDir, "swagger", "", "swagger目录")
}

func main() {
	//flag.Parse()
	configFile = "/api.toml"
	if configFile == "" {
		panic("请使用-c指定配置文件")
	}

	var state int32 = 1
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	ctx := logger.NewTraceIDContext(context.Background(), util.MustUUID())
	span := logger.StartSpanWithCall(ctx)

	call := api.Init(ctx,
		api.SetConfigFile(configFile),
		api.SetSwaggerDir(swaggerDir),
		api.SetVersion(VERSION))

	select {
	case sig := <-sc:
		atomic.StoreInt32(&state, 0)
		span().Printf("获取到退出信号[%s]", sig.String())
	}

	if call != nil {
		call()
	}
	span().Printf("服务退出")

	os.Exit(int(atomic.LoadInt32(&state)))
}
