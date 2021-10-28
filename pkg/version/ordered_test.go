package version_test

import (
	"fmt"
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/version"
	"gotest.tools/v3/assert"
)

type want struct {
	version  string
	isLatest bool
}

func TestOrderedResolver(t *testing.T) {
	tests := []struct {
		resolver     version.Resolver
		versionRange string
		want
	}{{
		resolver: version.OrderedResolver{},
		want:     want{version: "", isLatest: false},
	}, {
		resolver: version.OrderedResolver{Resolvers: []version.Resolver{
			version.StaticResolver{},
			version.StaticResolver{VersionString: "v3.4.5"},
		}},
		want: want{version: "v3.4.5", isLatest: true},
	}}
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tc.resolver.Version()
			assert.Equal(t, tc.resolver.Version(), tc.want.version)
			if tc.want.version != "" {
				latest, err := tc.resolver.IsLatest(tc.versionRange)
				assert.NilError(t, err)
				assert.Equal(t, latest, tc.want.isLatest)
			}
		})
	}
}
