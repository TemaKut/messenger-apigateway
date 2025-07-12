package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	logger *slog.Logger
}

func NewLogger(level LogLevel) (*Logger, error) {
	opts := &slog.HandlerOptions{}

	switch level {
	case LogLevelDebug:
		opts.Level = slog.LevelDebug
	case LogLevelInfo:
		opts.Level = slog.LevelInfo
	default:
		return nil, fmt.Errorf("error invalid log level %d", level)
	}

	return &Logger{
		logger: slog.New(slog.NewTextHandler(os.Stdout, opts)),
	}, nil
}

func (h *Logger) Infof(template string, args ...any) {
	h.logger.Info(fmt.Sprintf(template, args...))
}

func (h *Logger) Debugf(template string, args ...any) {
	h.logger.Debug(fmt.Sprintf(template, args...))
}

func (h *Logger) Errorf(template string, args ...any) {
	h.logger.Error(fmt.Sprintf(template, args...))
}
