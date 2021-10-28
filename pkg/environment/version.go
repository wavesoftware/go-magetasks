package environment

import (
	"log"

	"github.com/wavesoftware/go-magetasks/pkg/version"
)

// Check is used to verify environment values.
type Check Pair

// VersionResolver is used to resolve version information solely on
// environment variables.
type VersionResolver struct {
	VersionKey   Key
	IsApplicable []Check
	LatestOne    []Check
	ValuesSupplier
}

func (e VersionResolver) Version() string {
	values := e.environment()
	if !e.isApplicable(values) {
		return ""
	}
	return string(values[e.VersionKey])
}

func (e VersionResolver) IsLatest(versionRange string) (bool, error) {
	if versionRange == "" {
		versionRange = version.AnyVersion
	}
	if versionRange != version.AnyVersion {
		log.Printf("Ignoring version range %#v", versionRange)
	}
	values := e.environment()
	if !e.isApplicable(values) {
		return false, nil
	}
	for _, check := range e.LatestOne {
		if !check.test(values) {
			return false, nil
		}
	}
	return len(e.LatestOne) > 0, nil
}

func (e VersionResolver) environment() Values {
	supplier := Current
	if e.ValuesSupplier != nil {
		supplier = e.ValuesSupplier
	}
	return supplier()
}

func (e VersionResolver) isApplicable(values Values) bool {
	for _, check := range e.IsApplicable {
		if !check.test(values) {
			return false
		}
	}
	return true
}

func (c Check) test(values Values) bool {
	if c.Value == "" {
		_, ok := values[c.Key]
		return ok
	}
	val, ok := values[c.Key]
	return ok && val == c.Value
}
