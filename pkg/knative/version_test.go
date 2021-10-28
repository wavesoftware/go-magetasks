package knative_test

import (
	"fmt"
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/cache"
	"github.com/wavesoftware/go-magetasks/pkg/environment"
	"github.com/wavesoftware/go-magetasks/pkg/git"
	"github.com/wavesoftware/go-magetasks/pkg/knative"
	"github.com/wavesoftware/go-magetasks/pkg/strings"
	"github.com/wavesoftware/go-magetasks/pkg/testing/errors"
	"github.com/wavesoftware/go-magetasks/pkg/version"
	"gotest.tools/v3/assert"
)

func TestVersionResolver(t *testing.T) {
	tests := []testCase{{
		environment: environment.New(),
	}, {
		environment:  environment.New(),
		versionRange: "< 9999.9.9",
	}, {
		environment: environment.New("TAG=v4.6.23", "TAG_RELEASE=1"),
		version:     "v4.6.23",
	}, {
		// FIXME: This is not working because the CompositeResolver is buggy, and
		//        GitResolver resolves "" as version instead of using value resolved
		//        from environment.
		environment: environment.New("TAG=v6.23.0", "TAG_RELEASE=1"),
		version:     "v6.23.0",
		latest:      true,
	}}
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			resolver := tc.resolver()
			assert.Equal(t, resolver.Version(), tc.version)
			if tc.version != "" {
				latest, err := resolver.IsLatest(tc.versionRange)
				errors.Check(t, err, tc.err)
				assert.Equal(t, latest, tc.latest)
			}
		})
	}
}

type testCase struct {
	environment  environment.Values
	describe     string
	tags         strings.Set
	versionRange string
	version      string
	latest       bool
	err          error
}

func (tc testCase) resolver() version.Resolver {
	resolver := knative.VersionResolver{
		Git: &git.Resolver{
			Cache: cache.NoopCache{},
			Repository: git.StaticRepository{
				DescribeString: tc.describe,
				TagsSet:        tc.tags,
			},
		},
		Environmental: &environment.VersionResolver{
			VersionKey: "TAG",
			IsApplicable: []environment.Check{{
				Key: "TAG_RELEASE", Value: "1",
			}},
			LatestOne: []environment.Check{{
				Key: "is_auto_release", Value: "1",
			}},
			ValuesSupplier: func() environment.Values {
				return tc.environment
			},
		},
	}

	return resolver
}
