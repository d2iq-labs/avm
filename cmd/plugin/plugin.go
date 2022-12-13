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
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs a tool for a specific plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			avm, err := avmpkg.New(out)
			if err != nil {
				return fmt.Errorf("failed to initialize avm: %w", err)
			}

			defaultSource := avm.GetDefaultSource()

			pluginName, err := cmd.Flags().GetString("name")
			if err != nil {
				return fmt.Errorf("Error parsing plugin name: %v", err)
			}

			pluginVersion, err := cmd.Flags().GetString("version")
			if err != nil {
				return fmt.Errorf("Error parsing plugin version: %v", err)
			}

			pluginURL, err := cmd.Flags().GetString("url")
			if err != nil {
				return fmt.Errorf("Error parsing plugin url: %v", err)
			}

			err = defaultSource.InstallPluginVersion(
				&types.InstallPluginVersionRequest{
					Name:    pluginName,
					Version: pluginVersion,
					URL:     pluginURL,
				},
			)
			if err != nil {
				return fmt.Errorf("failed to install plugin: %w", err)
			}

			fmt.Printf("installed plugin %s with version %s\n", pluginName, pluginVersion)

			return nil
		},
	}

	var name string
	var version string
	var url string

	cmd.Flags().StringVar(&name, "name", "", "name of the plugin")
	cmd.Flags().StringVar(&version, "version", "", "version of the plugin")
	cmd.Flags().StringVar(&url, "url", "", "url of the plugin")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("version")

	return cmd
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
