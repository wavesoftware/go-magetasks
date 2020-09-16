package magetasks

import (
	"os"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/internal"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Check will run all lints checks.
func Check() {
	t := tasks.StartMultiline("🔍", "Checking")
	mg.Deps(revive, staticcheck)
	t.End(nil)
}

func revive() error {
	mg.Deps(internal.BuildDeps)
	reviveConfig := path.Join(internal.RepoDir(), "revive.toml")
	if fileExists(reviveConfig) {
		return sh.RunV("revive", "-config", "revive.toml", "-formatter", "stylish", "./...")
	}
	return nil
}

func staticcheck() error {
	mg.Deps(internal.BuildDeps)
	return sh.RunV("staticcheck", "-f", "stylish", "./...")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
