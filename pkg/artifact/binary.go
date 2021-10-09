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

// ResultBinaries is used to cache results of building a binaries.
const ResultBinaries = "binaries"

var (
	// ErrGoBuildFailed when go fails to build the project.
	ErrGoBuildFailed = errors.New("go build failed")

	// ErrInvalidArtifact when invalid type of artifact is given.
	ErrInvalidArtifact = errors.New("invalid artifact")
)

// Binary represents a binary that will be built.
type Binary struct {
	config.Metadata
	Platforms []Platform
}

func (b Binary) GetType() string {
	return "📦"
}

// BinaryBuilder is a regular binary Golang builder.
type BinaryBuilder struct{}

func (bb BinaryBuilder) Accepts(artifact config.Artifact) bool {
	_, ok := artifact.(Binary)
	return ok
}

func (bb BinaryBuilder) Build(artifact config.Artifact, notifier config.Notifier) config.Result {
	b, ok := artifact.(Binary)
	if !ok {
		return config.Result{Error: ErrInvalidArtifact}
	}
	info := make(map[string]string)
	var err error
	for _, platform := range b.Platforms {
		var bin string
		bin, err = b.buildForPlatform(platform, artifact.GetName(), notifier)
		if err == nil {
			info[fmt.Sprintf("%s-%s", platform.OS, platform.Architecture)] = bin
		} else {
			break
		}
	}
	return config.Result{Error: err, Info: info}
}

// BuildKey returns the config.ResultKey for a build command.
func BuildKey(artifact config.Artifact) config.ResultKey {
	return config.ResultKey(fmt.Sprintf("%s-%s-%s",
		artifact.GetType(), artifact.GetName(), ResultBinaries))
}

func (b Binary) buildForPlatform(
	platform Platform,
	name string,
	notifier config.Notifier,
) (string, error) {
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
	env := map[string]string{
		"GOOS":   string(platform.OS),
		"GOARCH": string(platform.Architecture),
	}
	notifier.Notify(fmt.Sprintf("go build (%s-%s)",
		platform.OS, platform.Architecture))
	err := sh.RunWithV(env, "go", args...)
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
