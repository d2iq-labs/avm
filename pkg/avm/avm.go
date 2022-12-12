// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package avm

import (
	"fmt"
	"github.com/d2iq-labs/avm/pkg/sources/asdf"
	"os"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/pkg/config"
)

var _ AVM = new(avm)

// AVM provides an interface for managing the avm.
type AVM interface {
	// ListSources returns a list of installed sources.
	ListSources() []string
}

type avm struct {
	out     output.Output
	cfg     config.Config
	sources map[string]bool
}

func New(out output.Output) (AVM, error) {
	cfg := config.DefaultConfig()

	out.V(4).Info(fmt.Sprintf("initializing avm with config: %+v", cfg))

	// ensure all required directories exist
	if err := ensureDirectory(cfg.HomeDir(), out); err != nil {
		return nil, err
	}

	if err := ensureDirectory(cfg.DataDir(), out); err != nil {
		return nil, err
	}

	if err := ensureDirectory(cfg.SourcesDir(), out); err != nil {
		return nil, err
	}

	avm := &avm{out: out, cfg: cfg, sources: map[string]bool{}}

	// install sources, if not already installed. this is a no-op if the source is already installed.

	if err := asdf.Install(cfg, out); err != nil {
		out.Errorf(err, "failed to install asdf")
	}

	avm.sources["asdf"] = true

	return avm, nil
}

func (a *avm) ListSources() []string {
	var sources []string

	for source, installed := range a.sources {
		if installed {
			sources = append(sources, source)
		}
	}

	return sources
}

// ensureDirectory ensures that the given directory path exists.
func ensureDirectory(path string, out output.Output) error {
	out.V(6).Info(fmt.Sprintf("ensuring directory exists: %s", path))

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}
