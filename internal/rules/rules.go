package rules

import (
	"path/filepath"
	"strings"
)

type Rule struct {
	TargetDir  string
	Extensions []string
}

type Resolver struct {
	extensionMap map[string]string
	fallbackDir  string
}

func NewResolver(rules []Rule, fallbackDir string) *Resolver {
	extMap := make(map[string]string)
	for _, rule := range rules {
		for _, ext := range rule.Extensions {
			cleanExt := normalizeExt(ext)
			extMap[cleanExt] = rule.TargetDir
		}
	}
	return &Resolver{
		extensionMap: extMap,
		fallbackDir:  fallbackDir,
	}
}

func (r *Resolver) Resolve(filename string) string {
	ext := filepath.Ext(filename)
	cleanExt := normalizeExt(ext)
	if cleanExt == "" {
		return r.fallbackDir
	}
	if target, found := r.extensionMap[cleanExt]; found {
		return target
	}
	return cleanExt
}

func normalizeExt(ext string) string {
	return strings.ToLower(strings.TrimPrefix(ext, "."))
}
