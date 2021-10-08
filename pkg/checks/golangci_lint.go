package checks

import (
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"github.com/wavesoftware/go-magetasks/pkg/output"
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
	configFiles := []string{".golangci.yaml", ".golangci.yml"}
	if configFilesMissing(configFiles) {
		output.Printlnf("%s file(s) missing. Skipping.", configFiles)
		return nil
	}
	if !files.ExecutableAvailable(golangciLintName) {
		output.Printlnf("%s executable isn't available on system PATH's."+
			" Skipping.", golangciLintName)
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

func configFilesMissing(configFiles []string) bool {
	for _, file := range configFiles {
		c := path.Join(files.ProjectDir(), file)
		if !files.DontExists(c) {
			return false
		}
	}
	return true
}
