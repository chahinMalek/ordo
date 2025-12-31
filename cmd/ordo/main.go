package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chahinMalek/ordo/internal/config"
	"github.com/chahinMalek/ordo/internal/rules"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}
	resolver := rules.NewResolver(cfg.Rules, cfg.FallbackDir)

	dummy_files := []string{
		"test.jpg",
		"test.env",
		"test.txt",
	}

	for _, file := range dummy_files {
		resolved_dir := resolver.Resolve(file, true, true)
		target_path := filepath.Join(resolved_dir, file)
		fmt.Printf("Resolving %s -> %s\n", file, target_path)
	}
}
