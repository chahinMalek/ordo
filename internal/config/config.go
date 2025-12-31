package config

import (
	"embed"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/chahinMalek/ordo/internal/rules"
)

//go:embed config.toml
var defaultConfigFile embed.FS

const (
	configFileName = "config.toml"
	appName        = "ordo"
)

type Config struct {
	FallbackDir string                `toml:"fallback_dir"`
	Rules       map[string]rules.Rule `toml:"rules"`
}

func Load() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, appName, configFileName)
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		return loadEmbedded(), nil
	}

	var cfg Config
	_, err = toml.DecodeFile(configPath, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func Init() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	ordoConfigDir := filepath.Join(configDir, appName)
	configPath := filepath.Join(ordoConfigDir, configFileName)

	// ensure the config directory exists
	err = os.MkdirAll(ordoConfigDir, 0755)
	if err != nil {
		return err
	}

	// Read default from embedded
	var data []byte
	data, err = defaultConfigFile.ReadFile(configFileName)
	if err != nil {
		return err
	}

	// Write the default one
	return os.WriteFile(configPath, data, 0644)
}

func loadEmbedded() *Config {
	var cfg Config
	data, err := defaultConfigFile.ReadFile(configFileName)
	if err != nil {
		return &Config{FallbackDir: "unclassified", Rules: make(map[string]rules.Rule)}
	}

	_, err = toml.Decode(string(data), &cfg)
	if err != nil {
		return &Config{FallbackDir: "unclassified", Rules: make(map[string]rules.Rule)}
	}

	return &cfg
}
