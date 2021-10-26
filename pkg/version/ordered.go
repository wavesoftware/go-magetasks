package version

// OrderedResolver will try to resolve version information using a list of
// resolvers. First one that be successful will be used.
type OrderedResolver struct {
	Resolvers []Resolver
}

func (c OrderedResolver) Version() string {
	return c.resolver().Version()
}

func (c OrderedResolver) IsLatest() bool {
	return c.resolver().IsLatest()
}

func (c OrderedResolver) resolver() Resolver {
	for _, resolver := range c.Resolvers {
		ver := resolver.Version()
		if ver != "" {
			return resolver
		}
	}
	return StaticResolver{}
}
