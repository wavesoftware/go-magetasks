package container

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/internal"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Publish publishes built images to a remote registry.
func Publish() {
	mg.Deps(Images)

	if len(config.Binaries) > 0 {
		t := tasks.StartMultiline("📤", "Publishing OCI images")
		errs := make([]error, 0)
		for _, binary := range config.Binaries {
			cf := containerFile(binary)
			if internal.DontExists(cf) {
				continue
			}
			args := []string{"push", imageName(binary)}
			err := sh.RunV(containerEngine(), args...)
			errs = append(errs, err)
		}
		t.End(errs...)
	}
}
