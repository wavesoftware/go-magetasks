package version

import (
	"github.com/wavesoftware/go-magetasks/pkg/environment"
)

// Check is used to verify environment values.
type Check environment.Pair

// EnvironmentResolver is used to resolve version information solely on
// environment variables.
type EnvironmentResolver struct {
	VersionKey   environment.Key
	IsApplicable []Check
	LatestOne    []Check
}

func (e EnvironmentResolver) Version() string {
	values := environment.Current()
	if !e.isApplicable(values) {
		return ""
	}
	return string(values[e.VersionKey])
}

func (e EnvironmentResolver) IsLatest() bool {
	values := environment.Current()
	if !e.isApplicable(values) {
		return false
	}
	for _, check := range e.LatestOne {
		if !check.test(values) {
			return false
		}
	}
	return len(e.LatestOne) > 0
}

func (e EnvironmentResolver) isApplicable(values environment.Values) bool {
	for _, check := range e.IsApplicable {
		if !check.test(values) {
			return false
		}
	}
	return true
}

func (c Check) test(values environment.Values) bool {
	if c.Value == "" {
		_, ok := values[c.Key]
		return ok
	}
	val, ok := values[c.Key]
	return ok && val == c.Value
}
