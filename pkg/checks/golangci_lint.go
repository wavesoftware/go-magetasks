package checks

import (
	"fmt"
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
)

const golangciLintName = "golangci-lint"

// GolangCiLintOptions contains options for GolangCi Lint.
type GolangCiLintOptions struct {
	// New when set will check only new code.
	New bool

	// Fix when set will try to fix issues.
	Fix bool
}

// GolangCiLint will configure golangci-lint in the build.
func GolangCiLint() config.Task {
	opts := GolangCiLintOptions{}
	return GolangCiLintWithOptions(opts)
}

// GolangCiLintWithOptions will configure golangci-lint in the build with
// options.
func GolangCiLintWithOptions(opts GolangCiLintOptions) config.Task {
	return config.Task{
		Name: golangciLintName,
		Operation: func() error {
			return golangCiLint(opts)
		},
	}
}

func golangCiLint(opts GolangCiLintOptions) error {
	configFile := ".golangci.yaml"
	c := path.Join(files.ProjectDir(), configFile)
	if files.DontExists(c) {
		fmt.Printf("%s file don't exists. Skipping.\n", configFile)
		return nil
	}
	if !files.ExecutableAvailable(golangciLintName) {
		fmt.Printf("%s executable isn't available on system PATH's."+
			" Skipping.\n", golangciLintName)
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
	return sh.RunV(golangciLintName, args...)
}
