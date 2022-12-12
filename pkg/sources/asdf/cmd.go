// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package asdf

import (
	"path/filepath"

	"github.com/magefile/mage/sh"
)

func (a *Asdf) execute(args ...string) (string, error) {
	return sh.OutputWith(a.env, filepath.Join(a.path, "bin", "entrypoint.sh"), args...)
}
