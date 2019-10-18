package api

import (
	"caas-micro/internal/app/api/config"
	"caas-micro/pkg/logger"
	"caas-micro/proto"
	"context"
	"os"

	"go.uber.org/dig"
)

type options struct {
	ConfigFile string
	ModelFile  string
	WWWDir     string
	SwaggerDir string
	Version    string
}

// Option 定义配置项
type Option func(*options)

// SetConfigFile 设定配置文件
func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

// SetSwaggerDir 设定swagger目录
func SetSwaggerDir(s string) Option {
	return func(o *options) {
		o.SwaggerDir = s
	}
}

// SetVersion 设定版本号
func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Init 应用初始化
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	err := config.LoadGlobalConfig(o.ConfigFile)
	handleError(err)

	cfg := config.GetGlobalConfig()

	logger.Printf(ctx, "服务启动，运行模式：%s，版本号：%s，进程号：%d", cfg.RunMode, o.Version, os.Getpid())

	if v := o.SwaggerDir; v != "" {
		cfg.Swagger = v
	}

	loggerCall, err := InitLogger()
	handleError(err)

	// 创建依赖注入容器
	container, containerCall := BuildContainer()

	// 初始化HTTP服务
	httpCall := InitHTTPServer(ctx, container)
	return func() {
		if httpCall != nil {
			httpCall()
		}
		if containerCall != nil {
			containerCall()
		}
		if loggerCall != nil {
			loggerCall()
		}
	}
}

// BuildContainer 创建依赖注入容器
func BuildContainer() (*dig.Container, func()) {
	// 创建依赖注入容器
	container := dig.New()

	err := proto.Inject(container)
	handleError(err)

	return container, nil
}
