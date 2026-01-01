package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chahinMalek/ordo/internal/config"
	"github.com/chahinMalek/ordo/internal/organizer"
	"github.com/chahinMalek/ordo/internal/rules"
	"github.com/spf13/cobra"
)

var (
	targetPath string
	useGroups  bool
	dryRun     bool
	verbose    bool
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "ordo",
	Short: "Ordo is a fast, safe, and deterministic CLI tool for organizing files.",
	Long:  `Ordo brings order to chaos by automatically categorizing files into folders based on their extensions or user-defined grouping rules.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cfg *config.Config
		var err error
		if configPath != "" {
			// todo: load from specific path
			cfg, err = config.Load()
		} else {
			cfg, err = config.Load()
		}

		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		absPath, err := os.Getwd()
		if targetPath != "" {
			absPath, err = filepath.Abs(targetPath)
		}
		if err != nil {
			return err
		}

		files, err := organizer.ListFiles(absPath)
		if err != nil {
			return err
		}

		resolver := rules.NewResolver(cfg.Rules)
		plan, err := organizer.Plan(absPath, files, useGroups, resolver)
		if err != nil {
			return err
		}

		exec := organizer.NewExecutor(dryRun, verbose)
		if !dryRun {
			plan.SavePlan(absPath)
		}
		return exec.Execute(plan)
	},
}

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert the last organization plan",
	RunE: func(cmd *cobra.Command, args []string) error {
		// todo: implement
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&targetPath, "path", "p", "", "Target directory to organize (default: current directory)")
	rootCmd.PersistentFlags().BoolVarP(&useGroups, "groups", "g", true, "Enable group-based organization")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Preview changes without moving any files")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print detailed information during execution")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "Path to custom configuration file")

	rootCmd.AddCommand(revertCmd)
}
