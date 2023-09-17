package logging

import (
	"context"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func getZapLogger() *zap.Logger {
	level, err := zapcore.ParseLevel(config.Configuration.Logging.Level)
	if err != nil {
		panic(fmt.Sprintf("can't parse log level: %v", err))
	}

	cfg := &zap.Config{
		Level:         zap.NewAtomicLevelAt(level),
		Development:   true,
		DisableCaller: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	zapLogger := zap.Must(cfg.Build())
	_ = zapLogger.Sync()
	return zapLogger
}

func setUp(ctx context.Context) *Logger {
	zapLogger := getZapLogger()
	log = &Logger{
		loggingBackend: zapLogger.Sugar(),
		gorm: &gormLoggingConfig{
			IgnoreRecordNotFoundError: true,
			TraceStr:                  "%s\n[%.3fms] [rows:%v] %s",
			TraceWarnStr:              "%s %s\n[%.3fms] [rows:%v] %s",
			TraceErrStr:               "%s %s\n[%.3fms] [rows:%v] %s",
			LogLevel:                  zapLogger.Level(),
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
