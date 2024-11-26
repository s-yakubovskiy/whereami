package config

import (
	"strings"

	"github.com/google/wire"
)

// ProvideConfig loads the configuration.
func ProvideConfig() (*AppConfig, error) {
	Cfg = *NewLoader().
		WithConfigType("yaml").
		WithDefaults(map[string]interface{}{}).
		WithAutomaticEnv(true).
		WithEnvKeyReplacer(strings.NewReplacer(".", "_")).
		Load(&AppConfig{})

	return &Cfg, nil
}

// ProvideSoundConfig extracts the SoundConfig from the full Config struct.
func ProvideLoggingConfig(cfg *AppConfig) *LoggingConfig {
	return &cfg.Logging
}

func ProvideProviderConfigs(cfg *AppConfig) *ProviderConfigs {
	return &cfg.ProviderConfigs
}

func ProvideServerConfig(cfg *AppConfig) *Server {
	return &cfg.Server
}

var ProviderSet = wire.NewSet(
	ProvideConfig,
	ProvideProviderConfigs,
	ProvideLoggingConfig,
	ProvideServerConfig)
