package checks

import (
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/internal"
)

// GolangCiLint will configure golangci-lint in the build.
func GolangCiLint() {
	config.Checks = append(config.Checks, config.CustomTask{
		Name: "golangci-lint",
		Task: golangCiLint,
	})
}

func golangCiLint() error {
	c := path.Join(internal.RepoDir(), ".golangci.yaml")
	if internal.DontExists(c) {
		return nil
	}
	return sh.RunV("golangci-lint", "run", "./...")
}
