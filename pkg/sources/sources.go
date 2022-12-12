// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package sources

// Source provides an interface for managing a plugin source.
type Source interface {
	// Name returns the name of the source.
	Name() string
	// Installs a plugin for that source
	Install(args ...string) (string, error)
}
