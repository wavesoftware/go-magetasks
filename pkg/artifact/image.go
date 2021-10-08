package artifact

import (
	"errors"
	"fmt"

	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/artifact/platform"
	"github.com/wavesoftware/go-magetasks/pkg/output"
	"github.com/wavesoftware/go-magetasks/pkg/output/color"
)

const imageReferenceKey = "oci.image.reference"

var errNotYetImplemented = errors.New("not yet implemented")

// Image is an OCI image that will be built from a binary.
type Image struct {
	config.Metadata
	Architectures []platform.Architecture
}

// KoBuilder builds images with Google's KO.
type KoBuilder struct{}

func (kb KoBuilder) Accepts(artifact config.Artifact) bool {
	_, ok := artifact.(Image)
	return ok
}

func (kb KoBuilder) Build(artifact config.Artifact, notifier config.Notifier) config.Result {
	// TODO: not yet implemented
	return config.Result{
		Error: fmt.Errorf("%w: ko builder", errNotYetImplemented)}
}

// ImageReferenceOf will try to fetch an image reference from image build result.
func ImageReferenceOf(img Image) config.Resolver {
	return func() string {
		result, ok := config.Actual().Context.Value(BuildKey(img)).(config.Result)
		if !ok || result.Failed() {
			return noImageReference(img)
		}
		ref, ok := result.Info[imageReferenceKey]
		if !ok {
			return noImageReference(img)
		}
		return ref
	}
}

func noImageReference(artifact config.Artifact) string {
	output.Println(color.Yellow("WARNING"),
		" can't resolve image reference for: ", artifact.GetName())
	return ""
}
