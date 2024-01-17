package log

import "github.com/rs/zerolog"

type Log interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Error(...error) *zerolog.Event
	Panic(...error) *zerolog.Event
	RawLogger() *zerolog.Logger
}

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelError
	LevelPanic
	LevelDisabled
)

func toZerologLevel(level Level) zerolog.Level {
	switch level {
	case LevelDebug:
		return zerolog.DebugLevel
	case LevelInfo:
		return zerolog.InfoLevel
	case LevelError:
		return zerolog.ErrorLevel
	case LevelPanic:
		return zerolog.PanicLevel
	case LevelDisabled:
		return zerolog.Disabled
	default:
		return zerolog.InfoLevel
	}
}
