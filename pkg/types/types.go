// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package types

// Source provides an interface for managing a plugin source.
type Source interface {
	// Name returns the name of the source.
	Name() string
}
