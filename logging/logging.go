package logging

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func setUp(ctx context.Context) *Logger {
	cfg := zap.NewProductionConfig()
	cfg.Development = true
	cfg.DisableCaller = true
	//cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	zapLogger, err := cfg.Build()
	if err != nil {
		panic(fmt.Sprintf("can't initialize zap logger: %v", err))
	}
	defer func(log *zap.Logger) {
		_ = log.Sync()
	}(zapLogger) // flushes buffer, if any
	log = &Logger{
		loggingBackend: zapLogger.Sugar(),
		gorm: &gormLoggingConfig{
			IgnoreRecordNotFoundError: true,
			TraceStr:                  "%s\n[%.3fms] [rows:%v] %s",
			TraceWarnStr:              "%s %s\n[%.3fms] [rows:%v] %s",
			TraceErrStr:               "%s %s\n[%.3fms] [rows:%v] %s",
			LogLevel:                  zapcore.DebugLevel,
		},
	}
	log.Debug(ctx, "Logger instance initialized")

	return log
}
func GetLogger(ctx context.Context) *Logger {
	if log == nil {
		log = setUp(ctx)
	} else {
		log.Debug(ctx, "Logger instance reused")
	}
	return log
}
