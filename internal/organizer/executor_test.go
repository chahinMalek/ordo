package organizer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExecutor(t *testing.T) {

	// setup mock directory
	tmpDir := t.TempDir()
	createFile(t, tmpDir, "document.pdf")
	createFile(t, tmpDir, "image.png")

	plan := &ActionPlan{
		MkDirs: []MkDirAction{
			{Dir: filepath.Join(tmpDir, "docs")},
		},
		Moves: []MoveAction{
			{
				SourcePath: filepath.Join(tmpDir, "document.pdf"),
				TargetPath: filepath.Join(tmpDir, "docs", "document.pdf"),
			},
		},
		Skips: []SkipAction{
			{Filename: "image.png", Reason: "no target directory"},
		},
	}

	exec := NewExecutor(false, false)
	if err := exec.Execute(plan); err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	for _, action := range plan.MkDirs {
		if _, err := os.Stat(action.Dir); err != nil {
			t.Fatalf("Directory %s does not exist", action.Dir)
		}
	}

	for _, action := range plan.Moves {
		if _, err := os.Stat(action.TargetPath); err != nil {
			t.Fatalf("File %s does not exist", action.TargetPath)
		}
	}

	for _, action := range plan.Skips {
		if _, err := os.Stat(action.Filename); err == nil {
			t.Fatalf("File %s should have been skipped", action.Filename)
		}
	}
}
