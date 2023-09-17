package logging

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"time"
)

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.loggingBackend.Level() > zap.DebugLevel {
		return
	}
	elapsed := time.Since(begin)
	switch { //TODO: must be debug logs
	case err != nil && l.loggingBackend.Level() == zap.ErrorLevel && (!errors.Is(err, gorm.ErrRecordNotFound) || !l.gorm.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			l.Errorf(ctx, l.gorm.TraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Errorf(ctx, l.gorm.TraceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.gorm.SlowThreshold && l.gorm.SlowThreshold != 0 && l.loggingBackend.Level() >= zap.WarnLevel:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.gorm.SlowThreshold)
		if rows == -1 {
			l.Warnf(ctx, l.gorm.TraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Warnf(ctx, l.gorm.TraceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.loggingBackend.Level() == zap.InfoLevel:
		sql, rows := fc()
		if rows == -1 {
			l.Infof(ctx, l.gorm.TraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Infof(ctx, l.gorm.TraceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}

func (l *Logger) Debug(_ context.Context, args ...interface{}) {
	l.loggingBackend.Debug(args...)
}
func (l *Logger) Debugf(_ context.Context, template string, args ...interface{}) {
	l.loggingBackend.Debugf(template, args...)
}
func (l *Logger) Debugln(_ context.Context, args ...interface{}) {
	l.loggingBackend.Info(args...)
}
func (l *Logger) Debugw(_ context.Context, msg string, keysAndValues ...interface{}) {
	l.loggingBackend.Infow(msg, keysAndValues...)
}
func (l *Logger) Error(_ context.Context, s string, args ...interface{}) {
	if len(args) > 0 {
		l.loggingBackend.Error(s, args)
		return
	}
	l.loggingBackend.Error(args...)
}
func (l *Logger) Errorf(_ context.Context, template string, args ...interface{}) {
	l.loggingBackend.Errorf(template, args...)
}
func (l *Logger) Errorln(_ context.Context, args ...interface{}) {
	l.loggingBackend.Error(args...)
}
func (l *Logger) Errorw(_ context.Context, msg string, keysAndValues ...interface{}) {
	l.loggingBackend.Errorw(msg, keysAndValues...)
}
func (l *Logger) Fatal(_ context.Context, args ...interface{}) {
	l.loggingBackend.Fatal(args...)
}
func (l *Logger) Fatalf(_ context.Context, template string, args ...interface{}) {
	l.loggingBackend.Fatalf(template, args...)
}
func (l *Logger) Fatalln(_ context.Context, args ...interface{}) {
	l.loggingBackend.Fatal(args...)
}
func (l *Logger) Fatalw(_ context.Context, msg string, keysAndValues ...interface{}) {
	l.loggingBackend.Fatalw(msg, keysAndValues...)
}
func (l *Logger) Info(_ context.Context, s string, args ...interface{}) {
	if len(args) > 0 {
		l.loggingBackend.Info(s, args)
		return
	}
	l.loggingBackend.Info(args...)
}
func (l *Logger) Infof(_ context.Context, template string, args ...interface{}) {
	l.loggingBackend.Infof(template, args...)
}
func (l *Logger) Infoln(_ context.Context, args ...interface{}) {
	l.loggingBackend.Info(args...)
}
func (l *Logger) Infow(_ context.Context, msg string, keysAndValues ...interface{}) {
	l.loggingBackend.Infow(msg, keysAndValues...)
}
func (l *Logger) Panic(_ context.Context, args ...interface{}) {
	l.loggingBackend.Panic(args...)
}
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.loggingBackend.Panicf(template, args...)
}
func (l *Logger) Panicln(_ context.Context, args ...interface{}) {
	l.loggingBackend.Panic(args...)
}
func (l *Logger) Panicw(_ context.Context, msg string, keysAndValues ...interface{}) {
	l.loggingBackend.Panicw(msg, keysAndValues...)
}
func (l *Logger) Sync() error {
	return l.loggingBackend.Sync()
}
func (l *Logger) Warn(_ context.Context, s string, args ...interface{}) {
	if len(args) > 0 {
		l.loggingBackend.Warn(s, args)
		return
	}
	l.loggingBackend.Warn(args...)
}
func (l *Logger) Warnf(_ context.Context, template string, args ...interface{}) {
	l.loggingBackend.Warnf(template, args...)
}
func (l *Logger) Warnln(_ context.Context, args ...interface{}) {
	l.loggingBackend.Warnln(args...)
}
func (l *Logger) Warnw(_ context.Context, msg string, keysAndValues ...interface{}) {
	l.loggingBackend.Warnw(msg, keysAndValues...)
}
