package config

import (
	"strings"
)

var Cfg AppConfig

type AppConfig struct {
	LogLevel     string        `mapstructure:"log_level"`
	Database     Database      `mapstructure:"database"`
	CrontabTasks []CrontabTask `mapstructure:"crontab_tasks"`
}

type Database struct {
	Enabled bool   `mapstructure:"enabled"`
	Type    string `mapstructure:"type"`
	Path    string `mapstructure:"path"`
}

type CrontabTask struct {
	Schedule string `mapstructure:"schedule"`
}

func Init() {
	Cfg = *NewLoader().
		WithConfigType("yaml").
		WithDefaults(map[string]interface{}{}).
		WithAutomaticEnv(true).
		WithEnvKeyReplacer(strings.NewReplacer(".", "_")).
		Load(&AppConfig{})
}
