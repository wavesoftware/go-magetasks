//go:build mage

package main

import (
	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/config/buildvars"
	"github.com/wavesoftware/go-magetasks/pkg/artifact"
	artifactimage "github.com/wavesoftware/go-magetasks/pkg/artifact/image"
	"github.com/wavesoftware/go-magetasks/pkg/artifact/platform"
	"github.com/wavesoftware/go-magetasks/pkg/checks"
	"github.com/wavesoftware/go-magetasks/pkg/image"
	"github.com/wavesoftware/go-magetasks/pkg/knative"
	"github.com/wavesoftware/go-magetasks/tests/example/build/overrides"
	"github.com/wavesoftware/go-magetasks/tests/example/pkg/metadata"
)

// Default target is set to Build
//
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Build

func init() { //nolint:gochecknoinits
	archs := []platform.Architecture{
		platform.AMD64, platform.ARM64, platform.S390X, platform.PPC64LE,
	}
	dummy := artifact.Image{
		Metadata: config.Metadata{Name: "dummy"},
		Labels: map[string]config.Resolver{
			"description": config.StaticResolver("A dummy image description"),
		},
		Architectures: archs,
	}
	sampleimage := artifact.Image{
		Metadata:      config.Metadata{Name: "sampleimage"},
		Architectures: archs,
	}
	other := artifact.Binary{
		Metadata: config.Metadata{
			Name: "other",
			BuildVariables: buildvars.Assemble([]buildvars.Operator{
				image.InfluenceableReference{
					Path:        metadata.ImagePath(metadata.DummyImage),
					EnvVariable: "MAGETASKS_EXAMPLE_DUMMY_IMAGE",
					Image:       dummy,
				},
				image.InfluenceableReference{
					Path:        metadata.ImagePath(metadata.SampleImage),
					EnvVariable: "MAGETASKS_EXAMPLE_SAMPLE_IMAGE",
					Image:       sampleimage,
				},
			}),
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
			Path: metadata.VersionPath(), Resolver: knative.NewVersionResolver(),
		},
		Artifacts: []config.Artifact{
			dummy, sampleimage, other,
		},
		Checks: []config.Task{
			checks.GolangCiLint(func(opts *checks.GolangCiLintOptions) {
				opts.Version = "v1.55.2"
			}),
			checks.Revive(),
			checks.Staticcheck(),
		},
		BuildVariables: map[string]config.Resolver{
			metadata.ImageBasenamePath():          artifactimage.BaseName,
			metadata.ImageBasenameSeparatorPath(): artifactimage.BaseNameSeparator,
		},
		Overrides: overrides.List,
	})
}
