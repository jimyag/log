package log

import (
	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _ Log = (*Logger)(nil)

type Logger struct {
	rawLogger *zerolog.Logger
}

func (l *Logger) Debug() *zerolog.Event {
	return l.rawLogger.Debug().Timestamp().Caller(1)
}

func (l *Logger) Info() *zerolog.Event {
	return l.rawLogger.Debug().Timestamp().Caller(1)
}

func (l *Logger) Error(errs ...error) *zerolog.Event {
	e := l.rawLogger.Error().Timestamp().Caller(1)
	switch len(errs) {
	case 0:
		return e
	case 1:
		return e.Err(errs[0])
	default:
		return e.Errs("errors", errs)
	}
}

func (l *Logger) Panic(errs ...error) *zerolog.Event {
	e := l.rawLogger.Panic().Timestamp().Caller(1)
	switch len(errs) {
	case 0:
		return e
	case 1:
		return e.Err(errs[0])
	default:
		return e.Errs("errors", errs)
	}
}

func (l *Logger) RawLogger() *zerolog.Logger {
	return l.rawLogger
}

func NewLogger(cfg *Config) *Logger {
	writer := lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAgeDays,
		Compress:   cfg.Compress,
		LocalTime:  cfg.LocalTime,
	}
	zLog := zerolog.New(&writer).With().Logger()
	logger := Logger{
		rawLogger: &zLog,
	}
	return &logger
}
