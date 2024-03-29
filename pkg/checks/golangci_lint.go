package checks

import (
	"fmt"
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
)

const golangciLintName = "golangci-lint"

type GolangCiLintParam func(*GolangCiLintOptions)

// GolangCiLintOptions contains options for GolangCi Lint.
type GolangCiLintOptions struct {
	// New when set will check only new code.
	New bool

	// Fix when set will try to fix issues.
	Fix bool

	// Version is a version of golangci-lint to use.
	Version string
}

func (o GolangCiLintOptions) effectiveVersion() string {
	if o.Version != "" {
		return o.Version
	}
	return "latest"
}

// GolangCiLint will configure golangci-lint in the build.
func GolangCiLint(params ...GolangCiLintParam) config.Task {
	opts := GolangCiLintOptions{}
	for _, p := range params {
		p(&opts)
	}
	return GolangCiLintWithOptions(opts)
}

// GolangCiLintWithOptions will configure golangci-lint in the build with
// options.
func GolangCiLintWithOptions(opts GolangCiLintOptions) config.Task {
	return config.Task{
		Name: golangciLintName,
		Operation: func(notifier config.Notifier) error {
			return golangCiLint(opts, notifier)
		},
		Overrides: []config.Configurator{
			config.NewBinaries("golangci/golangci-lint@" + opts.effectiveVersion()),
		},
	}
}

func golangCiLint(opts GolangCiLintOptions, notifier config.Notifier) error {
	configFiles := []string{".golangci.yaml", ".golangci.yml"}
	if configFilesMissing(configFiles) {
		skipBecauseOfMissingConfig(notifier, configFiles)
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
	cmd := fmt.Sprintf("%s/tools/golangci-lint", files.BuildDir())
	return sh.RunV(cmd, args...)
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
