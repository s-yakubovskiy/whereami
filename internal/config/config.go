package config

import (
	"fmt"
	"strings"
	"time"
)

var Cfg AppConfig

// Option is a function that applies a configuration setting
type Option func(*AppConfig)

// AppConfig holds the configuration settings for the application.
// It includes logging levels, database connections, scheduled tasks,
// provider details, and GPS configurations.
type AppConfig struct {
	LogLevel         string          `mapstructure:"log_level" yaml:"log_level"`
	Logging          LoggingConfig   `mapstructure:"logging"`
	Database         Database        `mapstructure:"database"`
	CrontabTasks     []CrontabTask   `mapstructure:"crontab_tasks" yaml:"crontab_tasks"`
	MainProvider     string          `mapstructure:"main_provider" yaml:"main_provider"`
	PublicIpProvider string          `mapstructure:"public_ip_provider" yaml:"public_ip_provider"`
	PublicIP         string          `mapstructure:"public_ip" yaml:"public_ip"`
	ProviderConfigs  ProviderConfigs `mapstructure:"provider_configs" yaml:"provider_configs"`
	GPSConfig        GPSConfig       `mapstructure:"gps" yaml:"gps"`
	Server           Server          `mapstructure:"server"`
}

type Server struct {
	HTTP HTTP `mapstructure:"http"`
	GRPC GRPC `mapstructure:"grpc"`
}

type HTTP struct {
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type GRPC struct {
	Address string        `mapstructure:"address"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func WithPublicIP(ip string) Option {
	return func(cfg *AppConfig) {
		cfg.PublicIP = ip
	}
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
	Output string `mapstructure:"output"`
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
	OpenWeather      ProviderConfig `mapstructure:"openweather"`
	Ifconfig         ProviderConfig `mapstructure:"ifconfig"`
	PublicIpProvider string         `mapstructure:"public_ip_provider" yaml:"public_ip_provider"`
}

func (c *ProviderConfigs) GetConfig(name string) (ProviderConfig, error) {
	switch name {
	case "ipapi":
		return c.IpApi, nil
	case "ipdata":
		return c.IpData, nil
	case "ipqualityscore":
		return c.IpQualityScore, nil
	case "openweather":
		return c.OpenWeather, nil
	default:
		return ProviderConfig{}, fmt.Errorf("unknown location service provider: %s", name)
	}
}

type ProviderConfig struct {
	Name    string `mapstructure:"name"`
	URL     string `mapstructure:"url"`
	APIKey  string `mapstructure:"api_key" yaml:"api_key"`
	Enabled bool   `mapstructure:"enabled"`
}

type GPSConfig struct {
	Enabled      bool          `mapstructure:"enabled"`
	Provider     string        `mapstructure:"provider"`
	Timeout      time.Duration `mapstructure:"timeout"`
	GpsdDumpFile string        `mapstructure:"gpsd_dump_file"`
	GpsdSocket   string        `mapstructure:"gpsd_socket"`
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

// func ApplyOptions(opts ...Option) *AppConfig {
func ApplyOptions(opts ...Option) {
	for _, opt := range opts {
		opt(&Cfg)
	}
}
