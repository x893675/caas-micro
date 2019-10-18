package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New for init zap log library
func New() (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte("debug"))
	if err != nil {
		return nil, err
	}

	cw := zapcore.Lock(os.Stdout)

	// file core 采用jsonEncoder
	cores := make([]zapcore.Core, 0, 1)

	// stdout core 采用 ConsoleEncoder
	ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	cores = append(cores, zapcore.NewCore(ce, cw, level))

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	zap.ReplaceGlobals(logger)

	return logger, err
}

//var ProviderSet = wire.NewSet(New, NewOptions)
