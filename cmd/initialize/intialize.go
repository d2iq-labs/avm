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

	"github.com/d2iq-labs/avm/pkg/sources/asdf"
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

			// source plugin directory
			sourcePath := filepath.Join(dataHome, "sources")

			if err := os.MkdirAll(sourcePath, 0755); err != nil {
				return err
			}

			// install sources
			err := asdf.Install(sourcePath, out)
			if err != nil {
				out.Errorf(err, "failed to install asdf")
			}

			out.Info("avm initialized")

			return nil
		},
	}
}
