package container

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/internal"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Images builds a OCI images of binaries.
func Images() {
	mg.Deps(magetasks.Binary)

	if len(config.Binaries) > 0 {
		t := tasks.StartMultiline("📦", "Packaging OCI images")
		errs := make([]error, 0)
		for _, binary := range config.Binaries {
			cf := containerFile(binary)
			if internal.DontExists(cf) {
				continue
			}
			args := []string{
				"build", "-f", cf, "-t", imageName(binary), ".",
			}
			err := sh.RunV(containerEngine(), args...)
			errs = append(errs, err)
		}
		t.End(errs...)
	}
}

func containerEngine() string {
	ce, err := resolveContainerEngine()
	ensure.NoError(err)
	return ce
}
