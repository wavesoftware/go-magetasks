package version

import "fmt"

// StaticResolver just return values from values given upfront.
type StaticResolver struct {
	VersionString       string
	Tags                []string
	SkipInvalidReleases bool
}

func (r StaticResolver) String() string {
	return fmt.Sprintf("%v,tags=%v,skipInvalidReleases=%v",
		r.VersionString, r.Tags, r.SkipInvalidReleases)
}

func (r StaticResolver) Version() string {
	return r.VersionString
}

func (r StaticResolver) IsLatest(versionRange string) (bool, error) {
	return IsLatestGivenReleases(
		r.VersionString, versionRange, r.SkipInvalidReleases,
		func() []string {
			return r.Tags
		},
	)
}
