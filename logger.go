package log

import (
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	var goModName string
	info, ok := debug.ReadBuildInfo()
	if ok {
		goModName = info.Main.Path
	} else {
		goModName = "unknown"
	}

	// set caller
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		file = strings.TrimPrefix(file, goModName+"/")
		return file + ":" + strconv.Itoa(line)
	}
}

var _ Log = (*logger)(nil)

type logger struct {
	rawLogger *zerolog.Logger
}

func (l *logger) Debug() *zerolog.Event {
	return l.rawLogger.Debug().Timestamp().Caller(1)
}

func (l *logger) Info() *zerolog.Event {
	return l.rawLogger.Info().Timestamp().Caller(1)
}

func (l *logger) Error(errs ...error) *zerolog.Event {
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

func (l *logger) Panic(errs ...error) *zerolog.Event {
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

func (l *logger) RawLogger() *zerolog.Logger {
	return l.rawLogger
}

func NewLog(cfg *Config) Log {
	writer := lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize, // megabytes
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAgeDays,
		Compress:   cfg.Compress,
		LocalTime:  cfg.LocalTime,
	}
	zLog := zerolog.New(&writer).With().Logger()
	zerolog.SetGlobalLevel(toZerologLevel(cfg.Level))
	logger := logger{
		rawLogger: &zLog,
	}
	return &logger
}
