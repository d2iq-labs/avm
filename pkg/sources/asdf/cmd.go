// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package asdf

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func (a *Asdf) execute(stdout, stderr io.Writer, args ...string) error {
	cmd := exec.Command(filepath.Join(a.path, "bin", "entrypoint.sh"), args...)

	// Set process outputs
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	// prepend the env to the existing env to allow overriding from the given env
	cmd.Env = append(os.Environ(), a.env...)

	fmt.Println("asdf executing: ", cmd.String())

	return cmd.Run()
}
