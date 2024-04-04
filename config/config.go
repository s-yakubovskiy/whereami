package config

import (
	"strings"
)

var Cfg AppConfig

type AppConfig struct {
	LogLevel        string          `mapstructure:"log_level"`
	Database        Database        `mapstructure:"database"`
	CrontabTasks    []CrontabTask   `mapstructure:"crontab_tasks"`
	MainProvider    string          `mapstructure:"main_provider"`
	ProviderConfigs ProviderConfigs `mapstructure:"provider_configs"`
}

type Database struct {
	Enabled bool   `mapstructure:"enabled"`
	Type    string `mapstructure:"type"`
	Path    string `mapstructure:"path"`
}

type ProviderConfigs struct {
	IpApi            ProviderConfig `mapstructure:"ipapi"`
	IpData           ProviderConfig `mapstructure:"ipdata"`
	IpQualityScore   ProviderConfig `mapstructure:"ipqualityscore"`
	PublicIpProvider string         `mapstructure:"public_ip_provider"`
}

type ProviderConfig struct {
	URL     string `mapstructure:"url"`
	APIKey  string `mapstructure:"api_key"`
	Enabled bool   `mapstructure:"enabled"`
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
