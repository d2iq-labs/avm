// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"fmt"

	"github.com/adrg/xdg"
)

type Config struct {

	// Home directory for the avm configuration.
	homeDir string

	// Data directory for the avm.
	dataDir string
}

// DefaultConfig returns the default configuration for avm.
func DefaultConfig() Config {
	return Config{
		homeDir: fmt.Sprintf("%s/avm", xdg.ConfigHome),
		dataDir: fmt.Sprintf("%s/avm", xdg.DataHome),
	}
}

func (c Config) HomeDir() string {
	return c.homeDir
}

func (c Config) DataDir() string {
	return c.dataDir
}

func (c Config) SourcesDir() string {
	return fmt.Sprintf("%s/sources", c.dataDir)
}
