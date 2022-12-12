// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package asdf

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/pkg/config"
)

// Install installs the asdf source plugin to given destination path
func Install(cfg config.Config, out output.Output) error {
	version := "v0.10.2"

	asdfPath := filepath.Join(cfg.SourcesDir(), "asdf")

	if _, err := os.Stat(asdfPath); os.IsNotExist(err) {
		out.V(6).Info(fmt.Sprintf("installing asdf version %s to %s", version, asdfPath))

		cmd := exec.Command("git", "clone", "--branch", version, "https://github.com/asdf-vm/asdf.git", asdfPath)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = os.Environ()

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to clone asdf: %w", err)
		}
	}

	return nil
}
