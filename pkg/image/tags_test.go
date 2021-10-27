package image_test

import (
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/image"
	"github.com/wavesoftware/go-magetasks/pkg/strings"
	"github.com/wavesoftware/go-magetasks/pkg/testing/errors"
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

func TestTags(t *testing.T) {
	tests := []struct {
		version string
		tags    []string
		want    []string
		err     error
	}{{
		version: "v1.5.2-2-g8cc3513",
		want:    []string{"v1.5.2-2-g8cc3513"},
	}, {
		version: "v1.5.3",
		want:    []string{"v1.5.3", "v1.5", "v1", "latest"},
	}, {
		version: "v1.5.3",
		tags:    []string{"v1.5.2", "v1.6.0"},
		want:    []string{"v1.5", "v1.5.3"},
	}, {
		version: "v1.5.3",
		tags:    []string{"wrong", "v1.5.2", "v1.5.4", "v1.6.0"},
		want:    []string{"v1.5.3"},
	}, {
		version: "1.5.3",
		want:    []string{"1.5.3", "1.5", "1", "latest"},
	}, {
		version: "af134dd",
		want:    []string{"af134dd"},
	}}
	for _, tc := range tests {
		resolver := version.StaticResolver{
			VersionString:       tc.version,
			Tags:                tc.tags,
			SkipInvalidReleases: true,
		}
		t.Run(resolver.String(), func(t *testing.T) {
			got, err := image.Tags(resolver)
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
