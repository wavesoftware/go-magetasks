package magetasks

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"github.com/wavesoftware/go-magetasks/pkg/ldflags"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Test will execute regular unit tests.
func Test() {
	mg.Deps(Check, files.EnsureBuildDir)
	t := tasks.StartMultiline("✅", "Testing")
	cmd := "richgo"
	if color.NoColor {
		cmd = "go"
	}
	args := []string{
		"test", "-v", "-covermode=count",
		fmt.Sprintf("-coverprofile=%s/coverage.out", files.BuildDir()),
	}
	args = append(appendVersion(args), "./...")
	err := sh.RunV(cmd, args...)
	t.End(err)
}

func appendVersion(args []string) []string {
	version := config.Actual().Version
	if version != nil {
		args = ldflags.NewBuilder().
			Add(version.Path, version.Resolver).
			Build(args)
	}
	return args
}
