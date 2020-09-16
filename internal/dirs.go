package internal

import (
	"os"
	"path"

	"github.com/wavesoftware/go-ensure"
)

// EnsureBuildDir creates a build directory.
func EnsureBuildDir() {
	d := path.Join(BuildDir(), "bin")
	ensure.NoError(os.MkdirAll(d, os.ModePerm))
}
