package git

import (
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

// TagBasedIsLatestStrategy is the default strategy, that uses a repository
// tags to determine if given version is latest within version range given.
func TagBasedIsLatestStrategy(r Resolver) func(versionRange string) (bool, error) {
	return func(versionRange string) (bool, error) {
		val, err := version.IsLatestGivenReleases(
			r.Version(), versionRange, true,
			r.resolveTags,
		)
		if err != nil {
			return false, err
		}
		return val, nil
	}
}
