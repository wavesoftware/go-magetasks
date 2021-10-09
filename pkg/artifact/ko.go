package artifact

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/google/ko/pkg/commands"
	"github.com/google/ko/pkg/commands/options"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/files"
	"golang.org/x/mod/modfile"
)

const koImportPath = "ko.import-path"

// ErrKoBuildFailed when th Google's ko fails to build.
var ErrKoBuildFailed = errors.New("ko build failed")

// KoBuilder builds images with Google's KO.
type KoBuilder struct{}

func (kb KoBuilder) Accepts(artifact config.Artifact) bool {
	_, ok := artifact.(Image)
	return ok
}

func (kb KoBuilder) Build(artifact config.Artifact, notifier config.Notifier) config.Result {
	image, ok := artifact.(Image)
	if !ok {
		return config.Result{Error: ErrInvalidArtifact}
	}
	bo := &options.BuildOptions{}
	ctx := config.Actual().Context
	builder, err := commands.NewBuilder(ctx, bo)
	if err != nil {
		return resultErrKoBuildFailed(err)
	}
	importPath, err := imageImportPath(image)
	if err != nil {
		return resultErrKoBuildFailed(err)
	}
	result, err := builder.Build(ctx, importPath)
	if err != nil {
		return resultErrKoBuildFailed(err)
	}
	digest, err := result.Digest()
	if err != nil {
		return resultErrKoBuildFailed(err)
	}
	notifier.Notify(fmt.Sprintf("built image: %s", digest))
	return config.Result{Info: map[string]string{
		imageReferenceKey: digest.String(),
	}}
}

func resultErrKoBuildFailed(err error) config.Result {
	return config.Result{
		Error: fmt.Errorf("%w: %v", ErrKoBuildFailed, err),
	}
}

func imageImportPath(image Image) (string, error) {
	binDir := fullBinaryDirectory(image.GetName())
	rs, err := lookForGoModule(binDir)
	if err != nil {
		return "", err
	}
	importPath := rs.resolve(binDir)
	if resolver, ok := image.Args[koImportPath]; ok {
		importPath = resolver()
	}
	return importPath, nil
}

func lookForGoModule(dir string) (lookupGoModuleResult, error) {
	rs := lookupGoModuleResult{}
	for i := 0; i < 10_000; i++ {
		modFile := path.Join(dir, "go.mod")
		if files.DontExists(modFile) {
			dir = path.Dir(dir)
			rs.directoryDistance++
			continue
		}
		bytes, err := ioutil.ReadFile(modFile)
		if err != nil {
			return rs, err
		}
		file, err := modfile.Parse(modFile, bytes, nil)
		if err != nil {
			return rs, err
		}
		rs.module = file
		return rs, nil
	}
	return rs, fmt.Errorf("%w: can't find go module", ErrKoBuildFailed)
}

type lookupGoModuleResult struct {
	module            *modfile.File
	directoryDistance int
}

func (r lookupGoModuleResult) resolve(dir string) string {
	root := dir
	for i := 0; i < r.directoryDistance; i++ {
		root = path.Dir(root)
	}
	p := strings.Replace(dir, root, "", 1)
	return path.Join(r.module.Module.Mod.Path, p)
}
