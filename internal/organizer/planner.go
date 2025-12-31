package organizer

import (
	"github.com/chahinMalek/ordo/internal/rules"
	"os"
	"path/filepath"
)

type MoveAction struct {
	SourcePath string
	TargetPath string
	TargetDir  string
}

type ActionPlan struct {
	Moves []MoveAction
}

func ListFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}

func Plan(
	baseDir string,
	filenames []string,
	useGroups bool,
	useFallback bool,
	resolver *rules.Resolver,
) (*ActionPlan, error) {

	plan := &ActionPlan{
		Moves: make([]MoveAction, 0),
	}
	for _, filename := range filenames {
		targetDir := resolver.Resolve(filename, useGroups, useFallback)

		sourcePath := filepath.Join(baseDir, filename)
		targetPath := filepath.Join(baseDir, targetDir, filename)

		// todo: what if a file with the same name is already inside the target directory?
		if sourcePath == targetPath {
			continue
		}
		plan.Moves = append(plan.Moves, MoveAction{
			SourcePath: sourcePath,
			TargetPath: targetPath,
			TargetDir:  targetDir,
		})
	}
	return plan, nil
}
