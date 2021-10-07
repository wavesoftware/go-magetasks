package artifact

import (
	"errors"
	"fmt"
	"path"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"github.com/wavesoftware/go-magetasks/pkg/ldflags"
)

// ErrGoBuildFailed when go fails to build the project.
var ErrGoBuildFailed = errors.New("go build failed")

// Binary represents a binary that will be built.
type Binary struct {
	config.Metadata
	Platforms []Platform
}

func (b Binary) Build(name string) config.BuildResult {
	info := make(map[string]string)
	var err error
	for _, platform := range b.Platforms {
		var bin string
		bin, err = b.buildForPlatform(platform, name)
		if err == nil {
			info[fmt.Sprintf("%s-%s", platform.OS, platform.Architecture)] = bin
		} else {
			break
		}
	}
	return config.BuildResult{Error: err, Info: info}
}

func (b Binary) buildForPlatform(platform Platform, name string) (string, error) {
	args := []string{
		"build",
	}
	version := config.Actual().Version
	if version != nil || len(b.Args) > 0 {
		builder := ldflags.NewBuilder()
		if version != nil {
			builder.Add(version.Path, version.Resolver)
		}
		for key, resolver := range b.Args {
			builder.Add(key, resolver)
		}
		args = builder.Build(args)
	}
	binary := fullBinaryName(platform, name)
	args = append(args, "-o", binary, fullBinaryDirectory(name))
	err := sh.RunV("go", args...)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrGoBuildFailed, err)
	}
	return binary, err
}

func fullBinaryName(platform Platform, name string) string {
	return path.Join(files.BuildDir(), "bin",
		fmt.Sprintf("%s-%s-%s", name, platform.OS, platform.Architecture))
}

func fullBinaryDirectory(name string) string {
	return path.Join(files.ProjectDir(), "cmd", name)
}
