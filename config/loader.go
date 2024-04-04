package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Loader struct {
	defaults       map[string]interface{}
	automaticEnv   bool
	envPrefix      string
	replacer       *strings.Replacer
	configType     string
	embeddedConfig []byte
}

func NewLoader() *Loader {
	return &Loader{}
}

func (loader *Loader) WithDefaults(defaults map[string]interface{}) *Loader {
	loader.defaults = defaults
	return loader
}

func (loader *Loader) WithEmbeddedConfig(c []byte) *Loader {
	loader.embeddedConfig = c
	return loader
}

func (loader *Loader) WithAutomaticEnv(automaticEnv bool) *Loader {
	loader.automaticEnv = automaticEnv
	return loader
}

func (loader *Loader) WithEnvKeyReplacer(replacer *strings.Replacer) *Loader {
	loader.replacer = replacer
	return loader
}

func (loader *Loader) WithEnvPrefix(envPrefix string) *Loader {
	loader.envPrefix = envPrefix
	return loader
}

func (loader *Loader) WithConfigType(ct string) *Loader {
	loader.configType = ct
	return loader
}

func (loader *Loader) decorateViper(v *viper.Viper) {
	v.SetConfigType(loader.configType)

	if "" != loader.envPrefix {
		v.SetEnvPrefix(loader.envPrefix)
	}

	for key, value := range loader.defaults {
		v.SetDefault(key, value)
	}

	if loader.automaticEnv {
		v.AutomaticEnv()
	}

	if nil != loader.replacer {
		v.SetEnvKeyReplacer(loader.replacer)
	}

	if len(loader.embeddedConfig) > 0 {
		if err := v.ReadConfig(bytes.NewReader(loader.embeddedConfig)); err != nil {
			panic(err)
		}
	}

	// Bind alias here
	if err := v.BindEnv("data.vault.token", "VAULT_TOKEN"); err != nil {
		panic(err)
	}
}

func (loader *Loader) Load(appConfig *AppConfig) *AppConfig {
	// Load default .env file
	_ = godotenv.Load()

	// Load additional environment variables from a specific file
	_ = godotenv.Load("/etc/whereami/environment")

	v := viper.New()
	loader.decorateViper(v)
	v.SetConfigName("config")         // Name of config file (without extension)
	v.SetConfigType("yaml")           // Config file type
	v.AddConfigPath("/etc/whereami/") // Path to look for the config file in
	// v.AddConfigPath("$HOME/.whereami/")        // Call multiple times to add many search paths
	// v.AddConfigPath("$HOME/.config/whereami/") // Call multiple times to add many search paths
	v.AddConfigPath("./config/") // Optionally look for config in the working directory

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
	}

	if err := v.Unmarshal(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
