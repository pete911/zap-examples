package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	nameContextKey contextKey = "request"
	idContextKey   contextKey = "request-id"
)

type contextKey string

func NewZapConfig(level zapcore.Level) zap.Config {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.FunctionKey = "fn"
	config.Level.SetLevel(level)
	return config
}

func GetLoggerFields(ctx context.Context, fields ...zap.Field) []zap.Field {
	var fieldsWithContext []zap.Field
	for _, k := range []contextKey{idContextKey, nameContextKey} {
		if v := ctx.Value(k); v != nil {
			fieldsWithContext = append(fieldsWithContext, zap.String(string(k), v.(string)))
		}
	}
	return append(fieldsWithContext, fields...)
}

func GetRequestContext(ctx context.Context, name string) context.Context {
	requestContext := context.WithValue(ctx, nameContextKey, name)
	return context.WithValue(requestContext, idContextKey, generateRequestId())
}

func GetRequestContextId(ctx context.Context) string {
	if v := ctx.Value(idContextKey); v != nil {
		return v.(string)
	}
	return ""
}

func generateRequestId() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%X-%X", b[0:4], b[4:])
}
