package organizer

import (
	"os"
	"path/filepath"

	"github.com/chahinMalek/ordo/internal/rules"
)

type MkDirAction struct {
	Dir string
}

type MoveAction struct {
	SourcePath string
	TargetPath string
}

type SkipAction struct {
	Filename string
	Reason   string
}

type ActionPlan struct {
	MkDirs []MkDirAction
	Moves  []MoveAction
	Skips  []SkipAction
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
	resolver *rules.Resolver,
) (*ActionPlan, error) {

	plan := &ActionPlan{
		MkDirs: make([]MkDirAction, 0),
		Moves:  make([]MoveAction, 0),
	}
	seenDirs := make(map[string]bool)

	for _, filename := range filenames {
		targetDirName := resolver.Resolve(filename, useGroups)
		if targetDirName == "" {
			plan.Skips = append(plan.Skips, SkipAction{
				Filename: filename,
				Reason:   "no target directory",
			})
			continue
		}

		// check if the target directory already exists
		targetDir := filepath.Join(baseDir, targetDirName)
		if _, err := os.Stat(targetDir); err != nil {
			if !seenDirs[targetDir] {
				plan.MkDirs = append(plan.MkDirs, MkDirAction{Dir: targetDir})
				seenDirs[targetDir] = true
			}
		}

		// check if the target directory is a file
		if info, err := os.Stat(targetDir); err == nil && !info.IsDir() {
			plan.Skips = append(plan.Skips, SkipAction{
				Filename: filename,
				Reason:   "target directory is a file",
			})
			continue
		}

		sourcePath := filepath.Join(baseDir, filename)
		targetPath := filepath.Join(targetDir, filename)

		// skip if the move action doesn't change anything
		if sourcePath == targetPath {
			plan.Skips = append(plan.Skips, SkipAction{
				Filename: filename,
				Reason:   "source and target paths are the same",
			})
			continue
		}

		// check if the target file already exists
		if _, err := os.Stat(targetPath); err == nil {
			plan.Skips = append(plan.Skips, SkipAction{
				Filename: filename,
				Reason:   "target file already exists in destination",
			})
			continue
		}

		plan.Moves = append(plan.Moves, MoveAction{
			SourcePath: sourcePath,
			TargetPath: targetPath,
		})
	}
	return plan, nil
}
