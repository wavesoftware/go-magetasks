package version

// StaticResolver is just static version information given up upfront.
type StaticResolver struct {
	VersionString string
	Latest        bool
}

func (s StaticResolver) Version() string {
	return s.VersionString
}

func (s StaticResolver) IsLatest() bool {
	return s.Latest
}
