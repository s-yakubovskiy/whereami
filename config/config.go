package config

import (
	"strings"
	"time"
)

var Cfg AppConfig

type AppConfig struct {
	LogLevel        string          `mapstructure:"log_level" yaml:"log_level"`
	Database        Database        `mapstructure:"database"`
	CrontabTasks    []CrontabTask   `mapstructure:"crontab_tasks" yaml:"crontab_tasks"`
	MainProvider    string          `mapstructure:"main_provider" yaml:"main_provider"`
	ProviderConfigs ProviderConfigs `mapstructure:"provider_configs" yaml:"provider_configs"`
	GPSConfig       GPSConfig       `mapstructure:"gps" yaml:"gps"`
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
	PublicIpProvider string         `mapstructure:"public_ip_provider" yaml:"public_ip_provider"`
}

type ProviderConfig struct {
	URL     string `mapstructure:"url"`
	APIKey  string `mapstructure:"api_key" yaml:"api_key"`
	Enabled bool   `mapstructure:"enabled"`
}

type GPSConfig struct {
	Enabled    bool          `mapstructure:"enabled"`
	Timeout    time.Duration `mapstructure:"timeout"`
	GpsdSocket string        `mapstructure:"gpsd_socket"`
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
