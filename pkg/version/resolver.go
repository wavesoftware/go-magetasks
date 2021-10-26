package version

// Resolver will resolve version string, and tell is that the latest artifact.
type Resolver interface {
	// Version returns the version string.
	Version() string
	// IsLatest tells if the version is the latest one.
	IsLatest() bool
}
