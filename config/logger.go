package config

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
)

// NewLogger mengembalikan logger
func NewLogger() *zap.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	logLevelConfig := GetEnv("APP_LOG_LEVEL", "info")

	logLevel := zap.InfoLevel
	if logLevelConfig == "debug" {
		logLevel = zap.DebugLevel
	}

	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	return logger
}

// NewDatabaseLogger mengembalikan logger untuk database
func NewDatabaseLogger(logger *zap.Logger) *DatabaseLogger {
	return &DatabaseLogger{logger: logger}
}

type DatabaseLogger struct {
	logger *zap.Logger
}

func (g *DatabaseLogger) LogMode(logger.LogLevel) logger.Interface {
	return g
}
func (g *DatabaseLogger) Info(_ context.Context, msg string, data ...any) {
	g.logger.Sugar().Infof(msg, data)
}
func (g *DatabaseLogger) Warn(_ context.Context, msg string, data ...any) {
	g.logger.Sugar().Warnf(msg, data)
}
func (g *DatabaseLogger) Error(_ context.Context, msg string, data ...any) {
	g.logger.Sugar().Errorf(msg, data)
}
func (g *DatabaseLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin).String()
	sql, rows := fc()
	fields := []zapcore.Field{
		zap.Int64("rows", rows),
		zap.String("elapsed", elapsed),
	}
	if err != nil {
		fields = append(fields, zap.String("sql", sql))
	}
	log := g.logger.With(fields...)
	if log.Sugar().Debugf(sql); err != nil {
		log.Sugar().Errorf(err.Error())
	}
}
