package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"time"
)

var log *Logger

type gormLoggingConfig struct {
	*logger.Config
	IgnoreRecordNotFoundError bool
	TraceStr                  string
	TraceWarnStr              string
	TraceErrStr               string
	SlowThreshold             time.Duration
	LogLevel                  zapcore.Level
}

type Logger struct {
	loggingBackend *zap.SugaredLogger
	gorm           *gormLoggingConfig
}

func (l *Logger) LogMode(_ logger.LogLevel) logger.Interface {
	newLogger := *l
	return &newLogger
}
