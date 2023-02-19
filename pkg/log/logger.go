package log

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger logr.Logger
)

type loggerCtxKeyType string

const (
	loggerCtxKey loggerCtxKeyType = "logger:sinan"
)

func init() {
	logger = Development(6, "console")
}

// Development initialize a development logger
func Development(logLevel int8, encoding string) logr.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(-zapcore.Level(logLevel))
	cfg.Encoding = encoding
	return buildLog(cfg)
}

// Production initialize a default logger to be used in production,
// it is used as the default logger.
func Production(logLevel int8, encoding string) logr.Logger {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(-zapcore.Level(logLevel))
	cfg.Encoding = encoding
	return buildLog(cfg)
}

func buildLog(cfg zap.Config) logr.Logger {
	zapLog, err := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return zapr.NewLogger(zapLog)
}

func SetLogger(l logr.Logger) {
	logger = l
}

func GetLogger() logr.Logger {
	return logger
}

func Enabled() bool { return logger.Enabled() }

func Info(msg string, keysAndValues ...interface{}) { logger.Info(msg, keysAndValues...) }

func Error(err error, msg string, keysAndValues ...interface{}) {
	logger.Error(err, msg, keysAndValues...)
}

func V(level int) logr.Logger { return logger.V(level) }

func WithValues(keysAndValues ...interface{}) logr.Logger { return logger.WithValues(keysAndValues...) }

func WithName(name string) logr.Logger { return logger.WithName(name) }

func FromCtx(ctx context.Context) logr.Logger {
	v := ctx.Value(loggerCtxKey)
	if v == nil {
		return logger
	}
	l, ok := v.(logr.Logger)
	if !ok {
		return logger
	}
	return l
}

func ToCtx(ctx context.Context, logger logr.Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, logger)
}
