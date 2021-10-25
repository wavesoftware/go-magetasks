//go:build mage
// +build mage

package main

import (
	"os"
	"strings"

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
		Metadata: config.Metadata{Name: "dummy"},
		Labels: map[string]config.Resolver{
			"description": config.StaticResolver("A dummy image description"),
		},
		Architectures: []platform.Architecture{
			platform.AMD64, platform.ARM64, platform.S390X, platform.PPC64LE,
		},
	}
	other := artifact.Binary{
		Metadata: config.Metadata{
			Name: "other",
			BuildVariables: config.NewBuildVariablesBuilder().
				ConditionallyAdd(referenceImageByDigest,
					metadata.ImagePath(), artifact.ImageReferenceOf(dummy)).
				Build(),
		},
		Platforms: []artifact.Platform{
			{OS: platform.Linux, Architecture: platform.AMD64},
			{OS: platform.Linux, Architecture: platform.ARM64},
			{OS: platform.Mac, Architecture: platform.AMD64},
			{OS: platform.Mac, Architecture: platform.ARM64},
			{OS: platform.Windows, Architecture: platform.AMD64},
		},
	}
	magetasks.Configure(config.Config{
		Version: &config.Version{
			Path: metadata.VersionPath(), Resolver: git.Resolver(),
		},
		Artifacts: []config.Artifact{
			dummy, other,
		},
		Checks: []config.Task{
			checks.GolangCiLint(),
			checks.Revive(),
			checks.Staticcheck(),
		},
		BuildVariables: map[string]config.Resolver{
			metadata.ImageBasenamePath(): func() string {
				return os.Getenv("IMAGE_BASENAME")
			},
		},
		Overrides: overrides.List,
	})
}

func skipImageReference() bool {
	if val, ok := os.LookupEnv("SKIP_IMAGE_REFERENCE"); ok {
		return strings.ToLower(val) == "true"
	}
	return false
}

func referenceImageByDigest() bool {
	return !skipImageReference()
}
