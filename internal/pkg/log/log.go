package log

import (
	"os"

	"pkg.zpf.com/golang/kit-scaffold/internal/pkg/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ErrorHandler 定义handler异常struct
type ErrorHandler struct {
	logger *zap.Logger
}

// InitLog 初始化 log
func InitLog(o *config.LogConfig) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)
	if err = level.UnmarshalText([]byte(o.Level)); err != nil {
		return nil, err
	}

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize,
		MaxBackups: o.MaxBackups,
		MaxAge:     o.MaxAge,
	})
	cw := zapcore.Lock(os.Stdout)

	// file core 采用jsonEncoder
	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	// stdout core 采用consoleEncoder
	if o.Stdout {
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return logger, err
}
