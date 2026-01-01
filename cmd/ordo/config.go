package main

import (
	"fmt"
	"os"

	"github.com/chahinMalek/ordo/internal/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var resetConfigCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset configuration to factory defaults",
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := config.GetConfigPath()
		if err != nil {
			return err
		}

		err = os.Remove(path)
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		err = config.Init()
		if err != nil {
			return err
		}

		fmt.Println("Configuration reset to defaults.")
		return nil
	},
}

func init() {
	configCmd.AddCommand(resetConfigCmd)
	rootCmd.AddCommand(configCmd)
}
