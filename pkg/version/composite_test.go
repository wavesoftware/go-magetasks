package version_test

import (
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/testing/errors"
	"github.com/wavesoftware/go-magetasks/pkg/version"
	"gotest.tools/v3/assert"
)

func TestCompositeResolver(t *testing.T) {
	resolver := version.CompositeResolver{
		VersionResolver:  version.StaticResolver{VersionString: "v1.0.1"},
		IsLatestResolver: version.StaticResolver{Tags: []string{"v1.0.0"}},
	}
	assert.Equal(t, resolver.Version(), "v1.0.1")
	latest, err := resolver.IsLatest(version.AnyVersion)
	errors.Check(t, err, nil)
	assert.Equal(t, latest, true)
}
