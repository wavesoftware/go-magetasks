package container

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Images builds a OCI images of binaries
func Images() {
	mg.Deps(magetasks.Binary)

	ce, err := resolveContainerEngine()
	ensure.NoError(err)
	if len(config.Binaries) > 0 {
		t := tasks.StartMultiline("📦", "Packaging OCI images")
		errs := make([]error, 0)
		for _, binary := range config.Binaries {
			args := []string{
				"build",
				"-f", containerFile(binary),
				"-t", imageName(binary),
				".",
			}
			args = append(args)
			err := sh.RunV(ce, args...)
			errs = append(errs, err)
		}
		t.End(errs...)
	}
}
