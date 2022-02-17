package metadata_test

import (
	"testing"

	"github.com/wavesoftware/go-magetasks/tests/example/pkg/metadata"
	"gotest.tools/v3/assert"
)

func TestImagePath(t *testing.T) {
	p := metadata.ImagePath()
	assert.Check(t, len(p) > 2)
}
