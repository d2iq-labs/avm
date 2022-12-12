// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package asdf

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "embed"

	"github.com/magefile/mage/sh"

	"github.com/mesosphere/dkp-cli-runtime/core/output"

	"github.com/d2iq-labs/avm/pkg/config"
	"github.com/d2iq-labs/avm/pkg/types"
)

var (
	//go:embed bin/entrypoint.sh
	entrypoint string

	// ensure that asdf implements the Source interface
	_ types.Source = new(Asdf)
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

func (a *Asdf) GetPlugin(name string) (*types.Plugin, error) {
	plugins, err := a.ListPlugins()
	if err != nil {
		return nil, err
	}

	plugin, ok := plugins[name]
	if !ok {
		return nil, fmt.Errorf("plugin %s not found", name)
	}

	return plugin, nil
}

func (a *Asdf) ListPlugins() (map[string]*types.Plugin, error) {
	plugins := make(map[string]*types.Plugin)

	out, err := a.execute("plugin", "list", "--urls")

	// no need to check for err here, asdf returns an error if no plugins are installed
	if out == "No plugins installed" {
		return plugins, nil
	}

	if err != nil {
		return nil, err
	}

	plugin := bufio.NewScanner(strings.NewReader(out))

	for plugin.Scan() {
		line := strings.TrimSpace(plugin.Text())

		tokens := strings.Fields(line)

		if len(tokens) != 2 {
			return nil, fmt.Errorf("unexpected output from asdf plugin list: %+v", tokens)
		}

		name := tokens[0]
		url := tokens[1]

		plugins[name] = &types.Plugin{
			Name:   name,
			URL:    url,
			Source: a,
		}
	}

	return plugins, nil
}

func (a *Asdf) ListPluginVersions(req *types.ListPluginVersionsRequest) ([]*types.PluginVersion, error) {
	var versions []*types.PluginVersion

	args := []string{"list"}

	if req.AllVersions {
		args = append(args, "all")
	}

	args = append(args, req.Name)

	out, err := a.execute(args...)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(out))

	for scanner.Scan() {
		version := strings.TrimSpace(scanner.Text())

		if version == "" {
			continue
		}

		versions = append(versions, &types.PluginVersion{Version: version})
	}

	return versions, nil
}

func (a *Asdf) InstallPluginVersion(req *types.InstallPluginVersionRequest) error {
	plugins, err := a.ListPlugins()
	if err != nil {
		return err
	}

	if _, ok := plugins[req.Name]; !ok {
		_, err := a.execute("plugin", "add", req.Name, req.URL)
		if err != nil {
			return fmt.Errorf("failed to install plugin %s: %w", req.Name, err)
		}
	}

	_, err = a.execute("install", req.Name, req.Version)
	if err != nil {
		return fmt.Errorf("failed to install plugin %s version %s: %w", req.Name, req.Version, err)
	}

	return nil
}

func (a *Asdf) execute(args ...string) (string, error) {
	bufOut := &bytes.Buffer{}
	bufErr := &bytes.Buffer{}

	_, err := sh.Exec(a.env, bufOut, bufErr, filepath.Join(a.path, "bin", "entrypoint.sh"), args...)
	if err != nil {
		return strings.TrimSuffix(bufErr.String(), "\n"), err
	}

	return strings.TrimSuffix(bufOut.String(), "\n"), err
}
