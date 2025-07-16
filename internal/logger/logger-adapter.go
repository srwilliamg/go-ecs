package logger

import (
	"context"

	"go.uber.org/zap"
)

type zapAdapter struct {
	l *zap.Logger
}

func NewZapAdapter(zapLogger *zap.Logger) Logger {
	return &zapAdapter{l: zapLogger}
}

func (a *zapAdapter) Enabled(level Level) bool {
	switch level {
	case DebugLevel:
		return a.l.Core().Enabled(zap.DebugLevel)
	case InfoLevel:
		return a.l.Core().Enabled(zap.InfoLevel)
	case WarnLevel:
		return a.l.Core().Enabled(zap.WarnLevel)
	case ErrorLevel:
		return a.l.Core().Enabled(zap.ErrorLevel)
	case FatalLevel:
		return a.l.Core().Enabled(zap.FatalLevel)
	default:
		return false
	}
}

func (a *zapAdapter) DebugCtx(ctx context.Context, msg string, fields ...Field) {
	a.l.Debug(msg, toZapFields(fields)...)
}

func (a *zapAdapter) InfoCtx(ctx context.Context, msg string, fields ...Field) {
	a.l.Info(msg, toZapFields(fields)...)
}

func (a *zapAdapter) WarnCtx(ctx context.Context, msg string, fields ...Field) {
	a.l.Warn(msg, toZapFields(fields)...)
}

func (a *zapAdapter) ErrorCtx(ctx context.Context, msg string, fields ...Field) {
	a.l.Error(msg, toZapFields(fields)...)
}

func (a *zapAdapter) FatalCtx(ctx context.Context, msg string, fields ...Field) {
	a.l.Fatal(msg, toZapFields(fields)...)
}

func (a *zapAdapter) Debug(msg string, fields ...Field) {
	a.l.Debug(msg, toZapFields(fields)...)
}

func (a *zapAdapter) Info(msg string, fields ...Field) {
	a.l.Info(msg, toZapFields(fields)...)
}

func (a *zapAdapter) Warn(msg string, fields ...Field) {
	a.l.Warn(msg, toZapFields(fields)...)
}

func (a *zapAdapter) Error(msg string, fields ...Field) {
	a.l.Error(msg, toZapFields(fields)...)
}

func (a *zapAdapter) Fatal(msg string, fields ...Field) {
	a.l.Fatal(msg, toZapFields(fields)...)
}

func (a *zapAdapter) With(fields ...Field) Logger {
	return &zapAdapter{l: a.l.With(toZapFields(fields)...)}
}

func (a *zapAdapter) Sync() error {
	return a.l.Sync()
}

func toZapFields(fields []Field) []zap.Field {
	zapFields := make([]zap.Field, len(fields))
	for i, f := range fields {
		zapFields[i] = zap.Any(f.Key, f.Value)
	}
	return zapFields
}
