package files

import (
	"os"
	"path"

	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks/config"
)

// EnsureBuildDir creates a build directory.
func EnsureBuildDir() {
	d := path.Join(BuildDir(), "bin")
	ensure.NoError(os.MkdirAll(d, os.ModePerm))
}

// BuildDir returns project build dir.
func BuildDir() string {
	return relativeToProjectRoot(config.Actual().BuildDirPath)
}

// ProjectDir returns project repo directory.
func ProjectDir() string {
	if config.Actual().ProjectDir != "" {
		return config.Actual().ProjectDir
	}
	repoDir, err := os.Getwd()
	ensure.NoError(err)
	return repoDir
}

func relativeToProjectRoot(paths []string) string {
	fullpath := make([]string, len(paths)+1)
	fullpath[0] = ProjectDir()
	for ix, elem := range paths {
		fullpath[ix+1] = elem
	}
	return path.Join(fullpath...)
}
