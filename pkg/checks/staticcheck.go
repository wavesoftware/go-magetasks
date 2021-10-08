package checks

import (
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"github.com/wavesoftware/go-magetasks/pkg/output"
)

// Staticcheck will configure staticcheck in the build.
func Staticcheck() config.Task {
	return config.Task{
		Name:      "staticcheck",
		Operation: staticcheck,
		Overrides: []config.Configurator{
			config.NewDependencies("honnef.co/go/tools/cmd/staticcheck"),
		},
	}
}

func staticcheck() error {
	configFile := "staticcheck.conf"
	c := path.Join(files.ProjectDir(), configFile)
	if files.DontExists(c) {
		output.Printlnf("%s file don't exists. Skipping.", configFile)
		return nil
	}
	return sh.RunV("staticcheck", "-f", "stylish", "./...")
}
