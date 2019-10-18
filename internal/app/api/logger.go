package api

import (
	"caas-micro/internal/app/api/config"
	"caas-micro/pkg/logger"
	"caas-micro/pkg/util"
	"os"
	"path/filepath"
)

// InitLogger 初始化日志
func InitLogger() (func(), error) {
	logger.SetTraceIDFunc(util.MustUUID)

	c := config.GetGlobalConfig().Log
	logger.SetLevel(c.Level)
	logger.SetFormatter(c.Format)

	// 设定日志输出
	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				os.MkdirAll(filepath.Dir(name), 0777)

				f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logger.SetOutput(f)
				file = f
			}
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}
