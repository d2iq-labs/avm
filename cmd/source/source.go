// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package source

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"
)

func NewCommand(out output.Output) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "source",
	}

	// Subcommands
	cmd.AddCommand(ListCommand(out))

	return cmd
}

// ListCommand creates a new command to list all source
func ListCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Lists a source",
		RunE: func(cmd *cobra.Command, args []string) error {
			out.V(6).Info(fmt.Sprintf("args: %v", args))
			return nil
		},
	}
}
