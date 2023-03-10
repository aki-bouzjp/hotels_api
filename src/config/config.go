package config

import (
	"github.com/BurntSushi/toml"
)

const (
	ERROR int = 0
	INFO  int = 1
	DEBUG int = 2
)

type LogConf struct {
	LogFilePath    string `toml:"log_filepath" env:"LOG_FILEPATH"`
	LogLevel       string `toml:"log_level" env:"LOG_LEVEL"`
	MaxSize        int    `toml:"max_size" env:"LOG_MAXSIZE"`
	MaxBackupTerms int    `toml:"max_backup_term" env:"LOG_BACKUP_TERM"`
	MaxAge         int    `toml:"max_age" env:"LOG_MAXAGE"`
}

type EventConf struct {
	EvengLogFilePath       string `toml:"eventlog_filepath" env:"EVENTLOG_FILEPATH"`
	EventLogOutputInterval string `toml:"eventlog_output_interval" env:"EVENTLOG_OUTPUT_INTERVAL"`
}

type Config struct {
	// System System `toml:"system" envPrefix:"SYSTEM_"`
	LogConf   LogConf   `toml:"log" envPrefix:"LOGCONF_"`
	EventConf EventConf `toml:"event" envPrefix:"EVENTCONF_"`
}

func New(path string) (*Config, error) {
	var conf Config
	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
