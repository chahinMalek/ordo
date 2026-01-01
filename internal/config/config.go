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
	Rules map[string]rules.Rule `toml:"rules"`
}

func Load() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

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

func (c *Config) Save() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	// ensure directory exists
	err = os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		return err
	}

	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(c)
}

func GetConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, appName, configFileName), nil
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
		return &Config{Rules: make(map[string]rules.Rule)}
	}

	_, err = toml.Decode(string(data), &cfg)
	if err != nil {
		return &Config{Rules: make(map[string]rules.Rule)}
	}

	return &cfg
}
