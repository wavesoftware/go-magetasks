package version_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/wavesoftware/go-magetasks/pkg/version"
)

func TestCompatibleRanges(t *testing.T) {
	tests := []struct {
		version string
		want    []string
		err     error
	}{{
		version: "v0.5.2-2-g8cc3513",
		want:    []string{},
	}, {
		version: "v0.5.3",
		want:    []string{"v0.5", "v0"},
	}, {
		version: "0.5.3",
		want:    []string{"0.5", "0"},
	}, {
		version: "af134dd",
		err:     version.ErrVersionIsNotSemantic,
	}}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.version, func(t *testing.T) {
			r := testResolver{tc.version}
			got, err := version.CompatibleRanges(r)
			checkError(t, err, tc.err)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want: %#v, got: %#v", tc.want, got)
			}
		})
	}
}

func checkError(tb testing.TB, got, want error) {
	tb.Helper()
	if want != nil {
		if !errors.Is(got, want) {
			tb.Fatalf("want %v, got %v", want, got)
		}
	} else {
		if got != nil {
			tb.Fatal("got unexpected:", got)
		}
	}
}

type testResolver struct {
	version string
}

func (r testResolver) Version() string {
	return r.version
}

func (r testResolver) IsLatest() bool {
	return false
}
