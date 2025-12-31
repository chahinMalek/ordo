package config

import (
	"github.com/BurntSushi/toml"
	"github.com/chahinMalek/ordo/internal/rules"
	"os"
	"path/filepath"
)

type Config struct {
	FallbackDir string       `toml:"fallback_dir"`
	Rules       []rules.Rule `toml:"rules"`
}

func DefaultConfig() *Config {
	return &Config{
		FallbackDir: "unclassified",
		Rules: []rules.Rule{
			{
				TargetDir:  "images",
				Extensions: []string{"jpg", "jpeg", "png", "gif", "webp"},
			},
			{
				TargetDir:  "videos",
				Extensions: []string{"mp4", "avi", "mkv", "mov", "wmv"},
			},
			{
				TargetDir:  "documents",
				Extensions: []string{"pdf", "doc", "docx", "xls", "xlsx", "ppt", "pptx"},
			},
			{
				TargetDir:  "music",
				Extensions: []string{"mp3", "wav", "flac", "aac", "ogg"},
			},
			{
				TargetDir:  "archives",
				Extensions: []string{"zip", "tar", "gz", "bz2", "7z"},
			},
			{
				TargetDir:  "text",
				Extensions: []string{"txt", "md", "markdown"},
			},
			{
				TargetDir:  "notebooks",
				Extensions: []string{"ipynb"},
			},
		},
	}
}

func Load() (*Config, error) {
	cfg := DefaultConfig()
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(configDir, "ordo", "config.toml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return cfg, nil
	}
	if _, err := toml.DecodeFile(configPath, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
