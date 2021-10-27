package version_test

import (
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/strings"
	"github.com/wavesoftware/go-magetasks/pkg/testing/errors"
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

func TestCompatibleRanges(t *testing.T) {
	tests := []struct {
		version             string
		tags                []string
		skipInvalidReleases bool
		want                []string
		err                 error
	}{{
		version: "v0.5.2-2-g8cc3513",
		want:    []string{},
	}, {
		version: "v0.5.3",
		want:    []string{"v0.5", "v0"},
	}, {
		version: "v0.5.3",
		tags:    []string{"v0.5.2", "v0.6.0"},
		want:    []string{"v0.5"},
	}, {
		version: "0.5.3",
		want:    []string{"0.5", "0"},
	}, {
		version: "af134dd",
		err:     version.ErrVersionIsNotValid,
	}}
	for _, tc := range tests {
		tr := version.StaticResolver{
			VersionString:       tc.version,
			Tags:                tc.tags,
			SkipInvalidReleases: tc.skipInvalidReleases,
		}
		t.Run(tr.String(), func(t *testing.T) {
			got, err := version.CompatibleRanges(tr)
			errors.Check(t, err, tc.err)
			if !equal(got, tc.want) {
				t.Fatalf("want: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func equal(a, b []string) bool {
	return strings.NewSet(a...).Equal(strings.NewSet(b...))
}
