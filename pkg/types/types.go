// Copyright 2022 D2iQ, Inc. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package types

// Source provides an interface for managing a plugin source.
type Source interface {
	// Name returns the name of the source.
	Name() string

	// GetPlugin returns the plugin with the given name.
	GetPlugin(name string) (*Plugin, error)

	// ListPlugins returns a list installed plugins.
	ListPlugins() (map[string]*Plugin, error)

	// ListPluginVersions returns a list of versions of the plugin according to the given request.
	ListPluginVersions(req *ListPluginVersionsRequest) ([]*PluginVersion, error)

	// InstallPluginVersion install the given plugin with specified version.
	InstallPluginVersion(req *InstallPluginVersionRequest) error
}

// Plugin represents a source plugin.
type Plugin struct {
	// Name is the name of the plugin.
	Name string

	// URL is the URL of the plugin.
	URL string

	// Source is the source of the plugin.
	Source Source
}

// Versions returns installed versions of the plugin.
func (p *Plugin) Versions() ([]*PluginVersion, error) {
	return p.Source.ListPluginVersions(
		&ListPluginVersionsRequest{
			Name:        p.Name,
			AllVersions: false,
		},
	)
}

// PluginVersion represents a version of a plugin.
type PluginVersion struct {
	// Version is the version of the plugin.
	Version string
}

// ListPluginVersionsRequest is the request for listing plugin versions.
type ListPluginVersionsRequest struct {
	// Name of the plugin to list versions for.
	Name string

	// List all versions of the plugin.
	AllVersions bool
}

// InstallPluginVersionRequest is the request to install a plugin with a given version.
type InstallPluginVersionRequest struct {
	// Name of the plugin to install.
	Name string

	// Version of the plugin to install. Depending on the source, this may be a optional.
	Version string

	// URL of the plugin to install. Depending on the source, this may be optional.k
	URL string
}
