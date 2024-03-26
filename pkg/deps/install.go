package deps

import (
	"context"
	"fmt"

	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/dotenv"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"github.com/wavesoftware/go-magetasks/pkg/output"
	"github.com/wavesoftware/go-magetasks/pkg/targets"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Install install build dependencies.
func Install(ctx context.Context) error {
	targets.Deps(ctx, dotenv.Load, output.Setup, files.EnsureBuildDir)
	deps := config.Actual().Dependencies
	dest := fmt.Sprintf("%s/tools", files.BuildDir())
	t := tasks.Start("🔧", "Installing tools",
		deps.Golang.Count()+deps.Binaries.Count() > 0)
	if err := deps.Golang.Install(ctx, t, dest); err != nil {
		return err
	}
	if err := deps.Binaries.Install(ctx, t, dest); err != nil {
		return err
	}
	t.End()
	return nil
}
