// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/mesosphere/dkp-cli-runtime/core/cmd/root"
	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/cmd/initialize"
	"github.com/d2iq-labs/avm/cmd/plugin"
	"github.com/d2iq-labs/avm/cmd/source"
)

func NewCommand(out, errOut io.Writer) (*cobra.Command, output.Output) {
	rootCmd, rootOpts := root.NewCommand(out, errOut)

	rootCmd.Use = filepath.Base(os.Args[0])

	// Subcommands
	rootCmd.AddCommand(initialize.NewCommand(rootOpts.Output))
	rootCmd.AddCommand(source.NewCommand(rootOpts.Output))
	rootCmd.AddCommand(plugin.NewCommand(rootOpts.Output))

	return rootCmd, rootOpts.Output
}

func Execute() {
	rootCmd, out := NewCommand(os.Stdout, os.Stderr)
	// disable cobra built-in error printing, we output the error with formatting.
	rootCmd.SilenceErrors = true

	if err := rootCmd.Execute(); err != nil {
		out.Error(err, "")
		os.Exit(1)
	}
}
