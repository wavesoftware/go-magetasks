package magetasks

import (
	"context"

	"github.com/magefile/mage/mg"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Build all expected artifacts.
func Build() {
	mg.Deps(Test, files.EnsureBuildDir)
	if len(config.Actual().Artifacts) > 0 {
		t := tasks.Start("🔨", "Building")
		errs := make([]error, 0)
		for name, artifact := range config.Actual().Artifacts {
			result := artifact.Build(name)
			if result.Failed() {
				errs = append(errs, result.Error)
			} else {
				config.WithContext(func(ctx context.Context) context.Context {
					return context.WithValue(ctx, buildKey{artifact}, result)
				})
			}
		}
		t.End(errs...)
	}
}

type buildKey struct {
	artifact config.Artifact
}
