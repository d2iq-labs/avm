// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package source

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	avmpkg "github.com/d2iq-labs/avm/pkg/avm"
	"github.com/d2iq-labs/avm/pkg/config"
	"github.com/d2iq-labs/avm/pkg/types"
)

func NewCommand(out output.Output) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Manage plugin sources",
	}

	// Subcommands
	cmd.AddCommand(ListCommand(out))
	cmd.AddCommand(AddComannd(out))

	return cmd
}

// ListCommand creates a new command to list all source
func ListCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List installed plugin sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			avm, err := avmpkg.New(out)
			if err != nil {
				return fmt.Errorf("failed to initialize avm: %w", err)
			}
			// we only have one source for now, so we can just print it.
			fmt.Printf("%s\n", avm.GetDefaultSource().Name())

			return nil
		},
	}
}

var sourceAlias map[string]string = map[string]string{"go": "golang"}
var defaultVersion map[string]string = map[string]string{"go": "1.19.3"}

// ListCommand creates a new command to list all source
func AddComannd(out output.Output) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Installs a source",
		RunE: func(cmd *cobra.Command, args []string) error {
			avm, err := avmpkg.New(out)
			if err != nil {
				return fmt.Errorf("failed to initialize avm: %w", err)
			}

			defaultSource := avm.GetDefaultSource()

			pluginName, err := cmd.Flags().GetString("name")
			if err != nil {
				return fmt.Errorf("Error parsing source name: %v", err)
			}

			err = defaultSource.InstallPluginVersion(
				&types.InstallPluginVersionRequest{
					Name:    sourceAlias[pluginName],
					Version: defaultVersion[pluginName],
				},
			)

			if err != nil {
				return fmt.Errorf("failed to install plugin: %w", err)
			}

			cfg := config.DefaultConfig()
			symlinkPath := filepath.Join(cfg.SourcesDir(), pluginName)
			sourcePath := filepath.Join(cfg.SourcesDir(), "asdf", "installs", sourceAlias[pluginName], defaultVersion[pluginName], pluginName)

			// Remove symlink if exists, in case we update the version
			if _, err := os.Lstat(symlinkPath); err == nil {
				os.Remove(symlinkPath)
			}

			err = os.Symlink(sourcePath, symlinkPath)
			if err != nil {
				return fmt.Errorf("failed to create symlink: %w", err)
			}

			return nil
		},
	}

	var name string
	var version string
	cmd.Flags().StringVar(&name, "name", "", "name of the source")
	cmd.Flags().StringVar(&version, "version", "", "version of the source")
	cmd.MarkFlagRequired("name")

	return cmd
}
