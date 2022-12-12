// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package asdf

import (
	"fmt"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/pkg/config"
	"github.com/d2iq-labs/avm/pkg/sources"
)

var (
	//go:embed bin/entrypoint.sh
	entrypoint string

	// ensure that asdf implements the Source interface
	_ sources.Source = new(Asdf)
)

// Asdf is a source plugin for asdf
type Asdf struct {
	// path for the asdf source plugin
	path string

	// env is a list of environment variables to set when executing asdf
	env map[string]string
}

// New creates a new asdf source plugin. If the source plugin is not installed, it will be installed.
func New(cfg config.Config, out output.Output) (*Asdf, error) {
	version := "v0.10.2"

	asdfPath := filepath.Join(cfg.SourcesDir(), "asdf")

	env := make(map[string]string)

	env["ASDF_DIR"] = asdfPath
	env["ASDF_DATA_DIR"] = asdfPath

	source := &Asdf{path: asdfPath, env: env}

	if _, err := os.Stat(asdfPath); os.IsNotExist(err) {
		out.V(6).Info(fmt.Sprintf("installing asdf version %s to %s", version, asdfPath))

		err := sh.RunV("git", "clone", "--branch", version, "https://github.com/asdf-vm/asdf.git", asdfPath)
		if err != nil {
			return nil, fmt.Errorf("failed to clone asdf: %w", err)
		}

		err = os.WriteFile(filepath.Join(asdfPath, "bin", "entrypoint.sh"), []byte(entrypoint), 0755)
		if err != nil {
			return nil, fmt.Errorf("failed to write entrypoint.sh: %w", err)
		}
	}

	version, err := source.execute("version")
	if err != nil {
		return nil, fmt.Errorf("failed to get asdf version: %w", err)
	}

	out.V(6).Info(fmt.Sprintf("asdf version %s installed to %s", version, asdfPath))

	return source, nil
}

func (a *Asdf) Name() string {
	return "asdf"
}
