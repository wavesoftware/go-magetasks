package environment_test

import (
	"strconv"
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/environment"
	"github.com/wavesoftware/go-magetasks/pkg/testing/errors"
	"github.com/wavesoftware/go-magetasks/pkg/version"
	"gotest.tools/v3/assert"
)

func TestVersionResolver(t *testing.T) {
	tests := []testCase{{
		environment: environment.New(),
	}, {
		environment: environment.New("TAG=v4.6.23", "TAG_RELEASE=1"),
		version:     "v4.6.23",
	}, {
		environment: environment.New("TAG=v6.23.0", "TAG_RELEASE=1", "is_auto_release=1"),
		version:     "v6.23.0",
	}}
	for i, tc := range tests {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			resolver := tc.resolver()
			assert.Equal(t, resolver.Version(), tc.version)
			_, err := resolver.IsLatest(version.AnyVersion)
			errors.Check(t, err, environment.ErrNotSupported)
		})
	}
}

type testCase struct {
	environment environment.Values
	version     string
}

func (tc testCase) resolver() version.Resolver {
	return environment.VersionResolver{
		VersionKey: "TAG",
		IsApplicable: []environment.Check{{
			Key: "TAG_RELEASE", Value: "1",
		}},
		ValuesSupplier: func() environment.Values {
			return tc.environment
		},
	}
}
