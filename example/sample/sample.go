package main

import "github.com/jimyag/log"

func main() {
	cfg := log.Config{
		Filename:   "sample.log",
		Level:      log.LevelDebug,
		MaxSize:    10,
		Compress:   false,
		LocalTime:  true,
		MaxBackups: 10,
		MaxAgeDays: 10,
	}
	logger := log.NewLog(&cfg)
	logger.Debug().Str("foo", "bar").Msg("debug")
	logger.Info().Str("foo", "bar").Msg("info")
	logger.Error().Str("foo", "bar").Msg("error")
	// make
	// make run
	// make clean

	log.Debug().Str("std", "foo bar").Msg("this is std log")
}
