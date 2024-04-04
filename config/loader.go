package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
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
	v.SetConfigName("config") // Name of config file (without extension)
	v.SetConfigType("yaml")   // Config file type

	// default configuration paths
	configPaths := []string{
		"/etc/whereami/",
		"$HOME/.whereami/",
		"$HOME/.config/whereami/",
		// "./config/",
	}
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, attempting to create a default one: %v\n", err)
		if !loader.checkConfigExists(configPaths, "config.yaml") {
			loader.createDefaultConfig("$HOME/.config/whereami/config.yaml")
		}

		// Try reading the config again after creating default
		if err := v.ReadInConfig(); err != nil {
			panic(fmt.Errorf("Fatal error reading config file: %w", err))
		}
	}

	if err := v.Unmarshal(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}

// aux functions to check if config not existed
func (loader *Loader) checkConfigExists(paths []string, fileName string) bool {
	for _, path := range paths {
		expandedPath := os.ExpandEnv(path) // Expand environment variables
		configFilePath := filepath.Join(expandedPath, fileName)
		if _, err := os.Stat(configFilePath); err == nil {
			// Config file exists
			return true
		}
	}
	// Config file does not exist in any of the paths
	return false
}

// aux functions to create default config from file
func (loader *Loader) createDefaultConfigFromFile(defaultPath string) {
	// Path to the default config file template
	defaultConfigPath := "./config/_default_config.yaml"

	// Read the default configuration from the template
	defaultConfig, err := os.ReadFile(defaultConfigPath)
	if err != nil {
		panic(fmt.Errorf("Failed to read default config template: %w", err))
	}

	// Expand environment variables in the path, if any
	expandedPath := os.ExpandEnv(defaultPath)
	// Ensure the directory exists
	dir := filepath.Dir(expandedPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Errorf("Failed to create directory for default config: %w", err))
	}

	// Create the default config file at the specified path
	file, err := os.Create(expandedPath)
	if err != nil {
		panic(fmt.Errorf("Failed to create default config file: %w", err))
	}
	defer file.Close()

	// Write the default configuration to the new file
	if _, err := file.Write(defaultConfig); err != nil {
		panic(fmt.Errorf("Failed to write to default config file: %w", err))
	}

	fmt.Printf("Created default config file at %s\n", expandedPath)
}

// aux functions to create default config from struct dump
func (loader *Loader) createDefaultConfig(defaultPath string) {
	defaultAppConfig := AppConfig{
		LogLevel:     "debug",
		MainProvider: "ipdata",
		ProviderConfigs: ProviderConfigs{
			IpApi: ProviderConfig{
				URL:     "http://ip-api.com/json/",
				APIKey:  "",
				Enabled: true,
			},
			IpData: ProviderConfig{
				URL:     "https://api.ipdata.co",
				APIKey:  "",
				Enabled: true,
			},
			IpQualityScore: ProviderConfig{
				URL:     "https://ipqualityscore.com/api/json/ip",
				APIKey:  "",
				Enabled: true,
			},
			PublicIpProvider: "ifconfig.me",
		},
		Database: Database{
			Enabled: true,
			Type:    "sqlite",
			Path:    "~/work/common/whereami/whereami_db.sqlite",
		},
		CrontabTasks: []CrontabTask{
			{Schedule: "@every 1h"},
		},
	}

	defaultConfigBytes, err := yaml.Marshal(defaultAppConfig)
	if err != nil {
		panic(fmt.Errorf("Failed to marshal default config: %w", err))
	}

	// Expand environment variables in the path, if any
	expandedPath := os.ExpandEnv(defaultPath)
	// Replace "~" with the user home directory for cross-platform compatibility
	if strings.HasPrefix(expandedPath, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Errorf("Failed to get user home directory: %w", err))
		}
		expandedPath = filepath.Join(home, expandedPath[2:])
	}

	// Ensure the directory exists
	dir := filepath.Dir(expandedPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(fmt.Errorf("Failed to create directory for default config: %w", err))
	}

	// Create the default config file
	file, err := os.Create(expandedPath)
	if err != nil {
		panic(fmt.Errorf("Failed to create default config file: %w", err))
	}
	defer file.Close()

	// Write the default configuration to the new file
	if _, err := file.Write(defaultConfigBytes); err != nil {
		panic(fmt.Errorf("Failed to write to default config file: %w", err))
	}

	fmt.Printf("Created default config file at %s\n", expandedPath)
}
