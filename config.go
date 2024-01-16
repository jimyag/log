package log

import (
	"github.com/rs/zerolog"
)

type Config struct {
	Filename   string        `toml:"filename"`     // 日志文件名
	Level      zerolog.Level `toml:"level"`        // 日志级别
	MaxSize    int           `toml:"max_size"`     // 单个日志文件的大小 MB
	MaxBackups int           `toml:"max_backups"`  // 保留的日志文件个数
	MaxAgeDays int           `toml:"max_age_days"` // 日志保留的最长时间: 天
	Compress   bool          `toml:"compress"`     // 是否压缩
	// 是否使用当前计算机的时间作为保存日志的名称
	// 默认使用UTC时间
	LocalTime bool `toml:"local_time"`
}
