package git_test

import (
	"fmt"
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/git"
	"github.com/wavesoftware/go-magetasks/pkg/testing/errors"
	"github.com/wavesoftware/go-magetasks/pkg/version"
	"gotest.tools/v3/assert"
)

func TestResolver(t *testing.T) {
	tests := []testCase{{
		version:      "v1.5.2-2-g8cc3513",
		versionRange: version.AnyVersion,
	}, {
		version:      "v1.5.3",
		versionRange: version.AnyVersion,
		latest:       true,
	}, {
		version:      "v1.5.3",
		tags:         []string{"v1.5.2", "v1.6.0"},
		versionRange: version.AnyVersion,
	}, {
		version:      "v1.5.3",
		tags:         []string{"wrong", "v1.5.2", "v1.5.4", "v1.6.0"},
		versionRange: version.AnyVersion,
	}, {
		version:      "1.5.3",
		versionRange: version.AnyVersion,
		latest:       true,
	}, {
		version:      "af134dd",
		versionRange: version.AnyVersion,
	}}
	for _, tc := range tests {
		tc := tc
		resolver := git.Resolver{
			Cache: noopCache{},
			Repository: staticRepository{
				describe: tc.version,
				tags:     tc.tags,
			},
		}
		t.Run(tc.String(), func(t *testing.T) {
			assert.Equal(t, resolver.Version(), tc.version)
			latest, err := resolver.IsLatest(tc.versionRange)
			errors.Check(t, err, tc.err)
			assert.Equal(t, latest, tc.latest)
		})
	}
}

type testCase struct {
	version      string
	tags         []string
	versionRange string
	latest       bool
	err          error
}

func (tc testCase) String() string {
	name := tc.version
	if len(tc.tags) > 0 {
		name = fmt.Sprintf("%s-%v", name, tc.tags)
	}
	if tc.versionRange != version.AnyVersion {
		name = fmt.Sprintf("%s-%s", name, tc.versionRange)
	}
	return name
}

type noopCache struct{}

func (n noopCache) Compute(_ interface{}, provider func() (interface{}, error)) (interface{}, error) {
	return provider()
}

func (n noopCache) Drop(_ interface{}) interface{} {
	return nil
}

type staticRepository struct {
	describe string
	tags     []string
}

func (s staticRepository) Describe() (string, error) {
	return s.describe, nil
}

func (s staticRepository) Tags() ([]string, error) {
	return s.tags, nil
}
