// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package initialize

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"
)

// NewCommand creates a new init command.
func NewCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize the avm configuration",
		RunE: func(cmd *cobra.Command, args []string) error {

			configHome := filepath.Join(xdg.ConfigHome, "avm")

			out.V(6).Info(fmt.Sprintf("ensuring config directory exists: %s", configHome))

			if err := os.MkdirAll(configHome, 0755); err != nil {
				return err
			}

			dataHome := filepath.Join(xdg.DataHome, "avm")

			out.V(6).Info(fmt.Sprintf("ensuring data directory exists: %s", dataHome))

			if err := os.MkdirAll(dataHome, 0755); err != nil {
				return err
			}

			out.Info("avm initialized")

			return nil
		},
	}
}
