package deps

import (
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
)

// Install install build dependencies.
func Install() error {
	for _, dep := range config.Actual().Dependencies.Installs() {
		err := sh.RunWith(map[string]string{"GO111MODULE": "off"}, "go", "get", dep)
		if err != nil {
			return err
		}
	}
	return nil
}
