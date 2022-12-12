// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package initialize

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/pkg/config"
	"github.com/d2iq-labs/avm/pkg/sources/asdf"
)

// NewCommand creates a new init command.
func NewCommand(out output.Output) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize the avm configuration",
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg := config.DefaultConfig()

			if err := ensureDirectory(cfg.HomeDir(), out); err != nil {
				return err
			}

			if err := ensureDirectory(cfg.DataDir(), out); err != nil {
				return err
			}

			if err := ensureDirectory(cfg.SourcesDir(), out); err != nil {
				return err
			}

			// install sources
			err := asdf.Install(cfg, out)
			if err != nil {
				out.Errorf(err, "failed to install asdf")
			}

			out.Info("avm initialized")

			return nil
		},
	}
}

// ensureDirectory ensures that the given directory path exists.
func ensureDirectory(path string, out output.Output) error {
	out.V(6).Info(fmt.Sprintf("ensuring directory exists: %s", path))

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}
