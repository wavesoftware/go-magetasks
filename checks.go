package magetasks

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/internal"
	"github.com/wavesoftware/go-magetasks/internal/tasks"
)

// Check will run all lints checks
func Check() {
	t := tasks.StartMultiline("🔍", "Checking")
	mg.Deps(revive, staticcheck)
	t.End(nil)
}

func revive() error {
	mg.Deps(internal.BuildDeps)
	return sh.RunV("revive", "-config", "revive.toml", "-formatter", "stylish", "./...")
}

func staticcheck() error {
	mg.Deps(internal.BuildDeps)
	return sh.RunV("staticcheck", "-f", "stylish", "./...")
}
