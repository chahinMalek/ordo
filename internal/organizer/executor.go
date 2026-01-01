package organizer

import (
	"fmt"
	"os"
)

type Executor struct {
	DryRun  bool
	Verbose bool
}

func NewExecutor(dryRun, verbose bool) *Executor {
	return &Executor{
		DryRun:  dryRun,
		Verbose: verbose,
	}
}

func (e *Executor) Execute(plan *ActionPlan) error {
	if e.DryRun {
		e.printPlan(plan)
		return nil
	}

	for _, action := range plan.MkDirs {
		if e.Verbose {
			fmt.Printf("Creating directory: %s\n", action.Dir)
		}
		err := os.MkdirAll(action.Dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", action.Dir, err)
		}
	}

	for _, action := range plan.Moves {
		if e.Verbose {
			fmt.Printf("Moving: %s -> %s\n", action.SourcePath, action.TargetPath)
		}
		err := os.Rename(action.SourcePath, action.TargetPath)
		if err != nil {
			return fmt.Errorf("failed to move file %s to %s: %w", action.SourcePath, action.TargetPath, err)
		}
	}

	if e.Verbose {
		for _, skip := range plan.Skips {
			fmt.Printf("Skipped: %s (%s)\n", skip.Filename, skip.Reason)
		}
	}

	return nil
}

func (e *Executor) printPlan(plan *ActionPlan) {
	fmt.Println("--- Action Plan (Dry Run) ---")

	if len(plan.MkDirs) > 0 {
		fmt.Println("\nDirectories to create:")
		for _, action := range plan.MkDirs {
			fmt.Printf("  [+] %s\n", action.Dir)
		}
	}

	if len(plan.Moves) > 0 {
		fmt.Println("\nFiles to move:")
		for _, action := range plan.Moves {
			fmt.Printf("  [M] %s -> %s\n", action.SourcePath, action.TargetPath)
		}
	}

	if len(plan.Skips) > 0 {
		fmt.Println("\nFiles to skip:")
		for _, skip := range plan.Skips {
			fmt.Printf("  [S] %s (%s)\n", skip.Filename, skip.Reason)
		}
	}

	fmt.Println("-----------------------------")
}
