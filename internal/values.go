package internal

import (
	"path"
	"runtime"

	"github.com/fatih/color"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks"
)

var (
	repoDir             = currentFileDir()
	buildDir            = relativeToRepo(magetasks.BuildDirPath)
	git                 = sh.OutCmd("git")
	gitVerCache *string = nil
	MageTag             = color.New(magetasks.MageTagColor).Sprint(magetasks.MageTagLabel)
)

func currentFileDir() string {
	_, filename, _, _ := runtime.Caller(magetasks.CallerDepth)
	return path.Dir(filename)
}

func relativeToRepo(paths []string) string {
	fullpath := make([]string, len(paths)+1)
	fullpath[0] = repoDir
	for ix, elem := range paths {
		fullpath[ix+1] = elem
	}
	return path.Join(fullpath...)
}
