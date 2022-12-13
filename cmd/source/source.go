// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package source

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	avmpkg "github.com/d2iq-labs/avm/pkg/avm"
)

func NewCommand(out output.Output) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Manage plugin sources",
	}

	// Subcommands
	cmd.AddCommand(ListCommand(out))

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
