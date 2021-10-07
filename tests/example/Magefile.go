//go:build mage
// +build mage

package main

import (

	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/artifact"
	"github.com/wavesoftware/go-magetasks/pkg/artifact/platform"
	"github.com/wavesoftware/go-magetasks/pkg/checks"
	"github.com/wavesoftware/go-magetasks/pkg/git"
	"github.com/wavesoftware/go-magetasks/tests/example/overrides"
	"github.com/wavesoftware/go-magetasks/tests/example/pkg/metadata"
)

// Default target is set to Build
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Build

func init() { //nolint:gochecknoinits
	dummy := artifact.Image{
		Metadata: config.Metadata{
			Args: map[string]config.Resolver{
				"DESC": func() string {
					return "A dummy image"
				},
			},
		},
		Architectures: []platform.Architecture{
			platform.AMD64, platform.ARM64, platform.S390X, platform.PPC64LE,
		},
	}
	other := artifact.Binary{
		Metadata: config.Metadata{
			Args: map[string]config.Resolver{
				metadata.ImagePath(): artifact.ImageReferenceOf(dummy),
			},
		},
		Platforms: []artifact.Platform{
			{OS: platform.Linux, Architecture: platform.AMD64},
			{OS: platform.Linux, Architecture: platform.ARM64},
			{OS: platform.Linux, Architecture: platform.S390X},
			{OS: platform.Linux, Architecture: platform.PPC64LE},
			{OS: platform.Mac, Architecture: platform.AMD64},
			{OS: platform.Mac, Architecture: platform.ARM64},
			{OS: platform.Windows, Architecture: platform.AMD64},
		},
	}
	magetasks.Configure(config.Config{
		Version: &config.Version{
			Path: metadata.VersionPath(), Resolver: git.Version,
		},
		Artifacts: map[string]config.Artifact{
			"dummy": dummy,
			"other": other,
		},
		Checks: []config.Task{
			checks.GolangCiLint(),
			checks.Revive(),
			checks.Staticcheck(),
		},
		Overrides: overrides.List,
	})
}
