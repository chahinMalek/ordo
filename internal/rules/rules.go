package rules

import (
	"path/filepath"
	"strings"
)

type Rule struct {
	Extensions []string `toml:"extensions"`
}

type Resolver struct {
	extensionMap map[string]string
	fallbackDir  string
}

func NewResolver(rules map[string]Rule, fallbackDir string) *Resolver {
	extMap := make(map[string]string)
	for targetDir, rule := range rules {
		for _, ext := range rule.Extensions {
			cleanExt := normalizeExt(ext)
			extMap[cleanExt] = targetDir
		}
	}
	return &Resolver{
		extensionMap: extMap,
		fallbackDir:  fallbackDir,
	}
}

func (r *Resolver) Resolve(filename string, useGroups bool, useFallback bool) string {
	ext := filepath.Ext(filename)
	cleanExt := normalizeExt(ext)
	if cleanExt == "" {
		if useFallback {
			return r.fallbackDir
		}
		return ""
	}
	if useGroups {
		if target, found := r.extensionMap[cleanExt]; found {
			return target
		}
	}
	return cleanExt
}

func normalizeExt(ext string) string {
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}
