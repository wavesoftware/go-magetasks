package checks

import (
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/internal"
)

// GolangCiLintOptions contains options for GolangCi Lint.
type GolangCiLintOptions struct {
	// New when set will check only new code.
	New bool

	// Fix when set will try to fix issues.
	Fix bool
}

// GolangCiLint will configure golangci-lint in the build.
func GolangCiLint() {
	opts := GolangCiLintOptions{}
	GolangCiLintWithOptions(opts)
}

// GolangCiLintWithOptions will configure golangci-lint in the build with
// options.
func GolangCiLintWithOptions(opts GolangCiLintOptions) {
	config.Checks = append(config.Checks, config.CustomTask{
		Name: "golangci-lint",
		Task: func() error {
			return golangCiLint(opts)
		},
	})
}

func golangCiLint(opts GolangCiLintOptions) error {
	c := path.Join(internal.RepoDir(), ".golangci.yaml")
	if internal.DontExists(c) {
		return nil
	}
	args := []string{"run"}
	if opts.Fix {
		args = append(args, "--fix")
	}
	if opts.New {
		args = append(args, "--new")
	}
	args = append(args, "./...")
	return sh.RunV("golangci-lint", args...)
}
