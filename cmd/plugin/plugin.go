// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package plugin

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"
)

func NewCommand(out output.Output) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "plugin",
		Short: "Plugin",
	}

	// Subcommands
	cmd.AddCommand(InstallCommand(out))
	cmd.AddCommand(AddCommand(out))
	cmd.AddCommand(RemoveCommand(out))
	cmd.AddCommand(ListCommand(out))

	return cmd
}

// InstallCommand creates a new command to install
func InstallCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "install",
		Short: "Installs a tool for a specific plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			out.V(6).Info(fmt.Sprintf("args: %v", args))
			return nil
		},
	}
}

// AddCommand creates a new command to add a plugin
func AddCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Adds a plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			out.V(6).Info(fmt.Sprintf("args: %v", args))
			return nil
		},
	}
}

// RemoveCommand creates a new command to remove a plugin
func RemoveCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "remove",
		Short: "Removes a plugin",
		RunE: func(cmd *cobra.Command, args []string) error {
			out.V(6).Info(fmt.Sprintf("args: %v", args))
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
			out.V(6).Info(fmt.Sprintf("args: %v", args))
			return nil
		},
	}
}
