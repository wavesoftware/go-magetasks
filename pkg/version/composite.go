package version

// CompositeResolver will use different resolvers for resolving version and is
// the latest information.
type CompositeResolver struct {
	VersionResolver  Resolver
	IsLatestResolver Resolver
}

func (c CompositeResolver) Version() string {
	return c.VersionResolver.Version()
}

func (c CompositeResolver) IsLatest(versionRange string) (bool, error) {
	return c.IsLatestResolver.IsLatest(versionRange)
}
