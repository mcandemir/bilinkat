package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewHandlerOptions(level slog.Leveler) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: level,
	}
}

func NewLogger(structureType string, handlerOptions *slog.HandlerOptions, stream *os.File) *Logger {
	var handler slog.Handler
	switch structureType {
	case "json":
		handler = slog.NewJSONHandler(stream, handlerOptions)
	case "text":
		handler = slog.NewTextHandler(stream, handlerOptions)
	default:
		handler = slog.NewJSONHandler(stream, handlerOptions)
	}
	return &Logger{
		logger: slog.New(handler),
	}
}

func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}
	l.logger.InfoContext(ctx, msg, args...)
}

func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}
	l.logger.ErrorContext(ctx, msg, args...)
}

func (l *Logger) Debug(ctx context.Context, msg string, args ...any) {
	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}
	l.logger.DebugContext(ctx, msg, args...)
}

func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	if requestID := ctx.Value("request_id"); requestID != nil {
		args = append(args, "request_id", requestID)
	}
	l.logger.WarnContext(ctx, msg, args...)
}
