// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package avm

import (
	"fmt"
	"os"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/pkg/config"
	asdfpkg "github.com/d2iq-labs/avm/pkg/sources/asdf"
	"github.com/d2iq-labs/avm/pkg/types"
)

var _ AVM = new(avm)

// AVM provides an interface for managing the avm.
type AVM interface {
	// GetDefaultSource returns the default the source.
	GetDefaultSource() types.Source
}

type avm struct {
	out     output.Output
	cfg     config.Config
	sources map[string]types.Source
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

	avm := &avm{out: out, cfg: cfg, sources: make(map[string]types.Source)}

	// install sources, if not already installed. this is a no-op if the source is already installed.
	asdf, err := asdfpkg.New(cfg, out)
	if err != nil {
		out.Errorf(err, "failed to install asdf")
	}

	avm.sources[asdf.Name()] = asdf

	return avm, nil
}

func (a *avm) GetDefaultSource() types.Source {
	// we only have one source right now, so just return it
	return a.sources["asdf"]
}

// ensureDirectory ensures that the given directory path exists.
func ensureDirectory(path string, out output.Output) error {
	out.V(6).Info(fmt.Sprintf("ensuring directory exists: %s", path))

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}
