package log

import "github.com/rs/zerolog"

var defaultLog = NewStdLog(LevelDebug)

func Debug() *zerolog.Event {
	return defaultLog.RawLogger().
		Debug().Timestamp().Caller(1)
}

func Info() *zerolog.Event {
	return defaultLog.RawLogger().
		Info().Timestamp().Caller(1)
}

func Warn(errs ...error) *zerolog.Event {
	e := defaultLog.RawLogger().
		Warn().Timestamp().Caller(1)
	switch len(errs) {
	case 0:
		return e
	case 1:
		return e.Err(errs[0])
	default:
		return e.Errs("errors", errs)
	}
}

func Error(errs ...error) *zerolog.Event {
	e := defaultLog.RawLogger().
		Error().Timestamp().Caller(1)
	switch len(errs) {
	case 0:
		return e
	case 1:
		return e.Err(errs[0])
	default:
		return e.Errs("errors", errs)
	}
}

func Panic(errs ...error) *zerolog.Event {
	e := defaultLog.RawLogger().
		Panic().Timestamp().Caller(1)
	switch len(errs) {
	case 0:
		return e
	case 1:
		return e.Err(errs[0])
	default:
		return e.Errs("errors", errs)
	}
}
