package organizer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/chahinMalek/ordo/internal/rules"
)

func TestPlan(t *testing.T) {

	// setup mock directory
	tmpDir := t.TempDir()
	createFile(t, tmpDir, "document.pdf")
	createFile(t, tmpDir, "image.png")
	createFile(t, tmpDir, "noext")

	err := os.MkdirAll(filepath.Join(tmpDir, "docs"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	// file collision scenario
	createFile(t, tmpDir, "docs/collision.pdf")
	createFile(t, tmpDir, "collision.pdf")

	// folder collision scenario
	createFile(t, tmpDir, "txt")
	createFile(t, tmpDir, "readme.txt")

	mockRules := map[string]rules.Rule{
		"docs": {Extensions: []string{"pdf"}},
	}
	resolver := rules.NewResolver(mockRules)

	files := []string{"document.pdf", "image.png", "noext", "collision.pdf", "readme.txt"}
	expectedMkDirs := []MkDirAction{
		{Dir: filepath.Join(tmpDir, "png")},
	}
	expectedMoves := []MoveAction{
		{
			SourcePath: filepath.Join(tmpDir, "document.pdf"),
			TargetPath: filepath.Join(tmpDir, "docs", "document.pdf"),
		},
		{
			SourcePath: filepath.Join(tmpDir, "image.png"),
			TargetPath: filepath.Join(tmpDir, "png", "image.png"),
		},
	}
	expectedSkips := []SkipAction{
		{Filename: "noext", Reason: "no target directory"},
		{Filename: "collision.pdf", Reason: "target file already exists in destination"},
		{Filename: "readme.txt", Reason: "target directory is a file"},
	}

	plan, err := Plan(tmpDir, files, true, resolver)
	if err != nil {
		t.Fatalf("Plan failed: %v", err)
	}

	// mkdir assertions
	if len(plan.MkDirs) != len(expectedMkDirs) {
		t.Errorf("Expected %d mkdir actions, got %d", len(expectedMkDirs), len(plan.MkDirs))
		return
	}
	for i, actualMkDir := range plan.MkDirs {
		expectedMkDir := expectedMkDirs[i]
		if actualMkDir != expectedMkDir {
			t.Errorf("MkDir %d: expected %v, got %v", i, expectedMkDir, actualMkDir)
		}
	}

	// move assertions
	if len(plan.Moves) != len(expectedMoves) {
		t.Errorf("Expected %d moves, got %d", len(expectedMoves), len(plan.Moves))
		t.Errorf("Expected moves: %v", expectedMoves)
		t.Errorf("Actual moves: %v", plan.Moves)
		return
	}
	for i, actualMove := range plan.Moves {
		expectedMove := expectedMoves[i]
		if actualMove != expectedMove {
			t.Errorf("Move %d: expected %v, got %v", i, expectedMove, actualMove)
		}
	}

	// skip assertions
	if len(plan.Skips) != len(expectedSkips) {
		t.Errorf("Expected %d skip actions, got %d", len(expectedSkips), len(plan.Skips))
		return
	}
	for i, actualSkip := range plan.Skips {
		expectedSkip := expectedSkips[i]
		if actualSkip != expectedSkip {
			t.Errorf("Skip %d: expected %v, got %v", i, expectedSkip, actualSkip)
		}
	}
}

func createFile(t *testing.T, base, path string) {
	t.Helper()
	fullPath := filepath.Join(base, path)
	err := os.WriteFile(fullPath, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file %s: %v", path, err)
	}
}
