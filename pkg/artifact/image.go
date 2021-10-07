package artifact

import (
	"errors"

	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/artifact/platform"
)

var errNotYetImplemented = errors.New("not yet implemented")

// Image is an OCI image that will be built from a binary.
type Image struct {
	config.Metadata
	Architectures []platform.Architecture
}

func (i Image) Build(name string) config.BuildResult {
	// TODO: not yet implemented
	return config.BuildResult{Error: errNotYetImplemented}
}

func ImageReferenceOf(_ Image) config.Resolver {
	// TODO: not yet implemented
	return func() string {
		return "tbd"
	}
}
