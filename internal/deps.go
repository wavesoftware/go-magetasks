package internal

import (
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks"
)

// BuildDeps install build dependencies
func BuildDeps() error {
	for _, dep := range magetasks.Dependencies {
		err := sh.RunWith(map[string]string{"GO111MODULE": "off"}, "go", "get", dep)
		if err != nil {
			return err
		}
	}
	return nil
}
