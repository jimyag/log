# log

Wraps zerolog so that it can be used in other projects.

## How to use

```shell
go get -u github.com/jimyag/log
```

[Other examples](./example)

```go
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
 log := log.NewLog(&cfg)
 log.Debug().Str("foo", "bar").Msg("debug")
 log.Info().Str("foo", "bar").Msg("info")
 log.Error().Str("foo", "bar").Msg("error")
}

```
