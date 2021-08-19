package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerContextFields = []string{"request"}

func NewZapConfig(level zapcore.Level) zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level.SetLevel(level)
	return config
}

func GetContext(ctx context.Context) []zap.Field {
	var fields []zap.Field
	for _, loggerField := range loggerContextFields {
		if v := ctx.Value(loggerField); v != nil {
			fields = append(fields, zap.String(loggerField, fmt.Sprintf("%s", v)))
		}
	}
	return fields
}
