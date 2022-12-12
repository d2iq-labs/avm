// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	avmpkg "github.com/d2iq-labs/avm/pkg/avm"
	"github.com/d2iq-labs/avm/pkg/types"
)

func NewCommand(out output.Output) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Plugin",
	}

	// Subcommands
	cmd.AddCommand(InstallCommand(out))
	cmd.AddCommand(ListCommand(out))

	return cmd
}

// InstallCommand creates a new command to install
func InstallCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Installs a tool for a specific plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			avm, err := avmpkg.New(out)
			if err != nil {
				return fmt.Errorf("failed to initialize avm: %w", err)
			}

			defaultSource := avm.GetDefaultSource()

			err = defaultSource.InstallPluginVersion(
				&types.InstallPluginVersionRequest{
					Name:    "golang",
					Version: "1.19.3",
				},
			)
			if err != nil {
				return fmt.Errorf("failed to install plugin: %w", err)
			}

			fmt.Printf("installed plugin %s with version %s\n", "golang", "1.19.3")

			return nil
		},
	}
}

// ListCommand creates a new command to remove a plugin
func ListCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all plugins",
		RunE: func(cmd *cobra.Command, args []string) error {
			avm, err := avmpkg.New(out)
			if err != nil {
				return fmt.Errorf("failed to initialize avm: %w", err)
			}

			defaultSource := avm.GetDefaultSource()

			plugins, err := defaultSource.ListPlugins()
			if err != nil {
				return fmt.Errorf("failed to list plugins: %w", err)
			}

			for _, plugin := range plugins {
				fmt.Printf("%s\n", plugin.Name)

				versions, err := plugin.Versions()
				if err != nil {
					return fmt.Errorf("failed to list versions: %w", err)
				}

				for _, version := range versions {
					fmt.Printf("\t%s\n", version.Version)
				}
			}
			return nil
		},
	}
}
