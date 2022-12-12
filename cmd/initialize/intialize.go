// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package initialize

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	avmpkg "github.com/d2iq-labs/avm/pkg/avm"
)

// NewCommand creates a new init command.
func NewCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize the avm configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			avm, err := avmpkg.New(out)
			if err != nil {
				return fmt.Errorf("failed to initialize avm: %w", err)
			}

			fmt.Printf("Initialized avm with sources: %s\n", avm.ListSources())

			return nil
		},
	}
}
