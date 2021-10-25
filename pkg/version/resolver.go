package version

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blang/semver"
)

// ErrVersionIsNotSemantic when given version is not semantic.
var ErrVersionIsNotSemantic = errors.New("version is not semantic")

// Resolver will resolve version string, and tell is that the latest artifact.
type Resolver interface {
	// Version returns the version string.
	Version() string
	// IsLatest tells if the version is the latest one.
	IsLatest() bool
}

// CompatibleRanges will resolve compatible ranges from a version resolver.
func CompatibleRanges(r Resolver) ([]string, error) {
	v := r.Version()
	prefix := ""
	if strings.HasPrefix(v, "v") {
		prefix = "v"
	}
	sv, err := semver.ParseTolerant(v)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrVersionIsNotSemantic, err)
	}
	if len(sv.Pre) == 0 && len(sv.Build) == 0 {
		return []string{
			fmt.Sprintf("%s%d.%d", prefix, sv.Major, sv.Minor),
			fmt.Sprintf("%s%d", prefix, sv.Major),
		}, nil
	}
	return []string{}, nil
}
