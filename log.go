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
