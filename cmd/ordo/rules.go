package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/chahinMalek/ordo/internal/config"
	"github.com/chahinMalek/ordo/internal/rules"
	"github.com/spf13/cobra"
)

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Manage organization rules",
}

var listRulesCmd = &cobra.Command{
	Use:   "list",
	Short: "List current rules",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if len(cfg.Rules) == 0 {
			fmt.Println("No rules defined.")
			return nil
		}

		fmt.Println("Current Rules:")
		keys := make([]string, 0, len(cfg.Rules))
		for k := range cfg.Rules {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			fmt.Printf("  %s: %s\n", k, strings.Join(cfg.Rules[k].Extensions, ", "))
		}
		return nil
	},
}

var addRuleCmd = &cobra.Command{
	Use:   "add <group> <extensions...>",
	Short: "Add or update a custom rule",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		group := args[0]
		exts := args[1:]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		rule := cfg.Rules[group]
		rule.Extensions = append(rule.Extensions, exts...)

		// Deduplicate and normalize
		uniqueExts := make(map[string]bool)
		for _, e := range rule.Extensions {
			uniqueExts[strings.ToLower(strings.TrimPrefix(e, "."))] = true
		}

		rule.Extensions = make([]string, 0, len(uniqueExts))
		for e := range uniqueExts {
			rule.Extensions = append(rule.Extensions, e)
		}
		sort.Strings(rule.Extensions)

		if cfg.Rules == nil {
			cfg.Rules = make(map[string]rules.Rule)
		}
		cfg.Rules[group] = rule

		err = cfg.Save()
		if err != nil {
			return err
		}

		fmt.Printf("Rule '%s' updated with extensions: %s\n", group, strings.Join(rule.Extensions, ", "))
		return nil
	},
}

var deleteRuleCmd = &cobra.Command{
	Use:   "delete <group>",
	Short: "Delete a rule group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		group := args[0]

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		if _, ok := cfg.Rules[group]; !ok {
			return fmt.Errorf("rule group '%s' not found", group)
		}

		delete(cfg.Rules, group)
		err = cfg.Save()
		if err != nil {
			return err
		}

		fmt.Printf("Rule group '%s' deleted.\n", group)
		return nil
	},
}

var removeRuleCmd = &cobra.Command{
	Use:   "remove <group> <extension>",
	Short: "Remove an extension from a rule group",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		group := args[0]
		ext := strings.ToLower(strings.TrimPrefix(args[1], "."))

		cfg, err := config.Load()
		if err != nil {
			return err
		}

		rule, ok := cfg.Rules[group]
		if !ok {
			return fmt.Errorf("rule group '%s' not found", group)
		}

		newExts := make([]string, 0)
		found := false
		for _, e := range rule.Extensions {
			if strings.ToLower(strings.TrimPrefix(e, ".")) == ext {
				found = true
				continue
			}
			newExts = append(newExts, e)
		}

		if !found {
			return fmt.Errorf("extension '%s' not found in group '%s'", ext, group)
		}

		rule.Extensions = newExts
		cfg.Rules[group] = rule

		err = cfg.Save()
		if err != nil {
			return err
		}

		fmt.Printf("Extension '%s' removed from group '%s'.\n", ext, group)
		return nil
	},
}

func init() {
	rulesCmd.AddCommand(listRulesCmd)
	rulesCmd.AddCommand(addRuleCmd)
	rulesCmd.AddCommand(removeRuleCmd)
	rulesCmd.AddCommand(deleteRuleCmd)
	rootCmd.AddCommand(rulesCmd)
}
