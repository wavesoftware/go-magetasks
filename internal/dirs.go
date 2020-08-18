package internal

import (
	"os"
	"path"

	"github.com/wavesoftware/go-ensure"
)

// ensureBuildDir creates a build directory
func ensureBuildDir() {
	d := path.Join(buildDir, "bin")
	ensure.NoError(os.MkdirAll(d, os.ModePerm))
}
