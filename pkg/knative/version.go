package knative

import (
	"github.com/wavesoftware/go-magetasks/pkg/environment"
	"github.com/wavesoftware/go-magetasks/pkg/git"
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

// VersionResolver is version resolver implementation directly targeting
// Knative project CI.
type VersionResolver struct {
	Git           *git.Resolver
	Environmental *environment.VersionResolver
}

func (v VersionResolver) Version() string {
	return v.resolver().Version()
}

func (v VersionResolver) IsLatest(versionRange string) (bool, error) {
	return v.resolver().IsLatest(versionRange)
}

func (v VersionResolver) resolver() version.Resolver {
	gitResolver := v.gitResolver()
	return version.OrderedResolver{Resolvers: []version.Resolver{
		version.CompositeResolver{
			VersionResolver:  v.environmentalResolver(),
			IsLatestResolver: gitResolver,
		},
		gitResolver,
	}}
}

func (v VersionResolver) gitResolver() version.Resolver {
	resolver := git.Resolver{}
	if v.Git != nil {
		resolver = *v.Git
	}
	return resolver
}

func (v VersionResolver) environmentalResolver() version.Resolver {
	resolver := environment.VersionResolver{
		VersionKey: "TAG",
		IsApplicable: []environment.Check{
			{Key: "TAG_RELEASE", Value: "1"},
			{Key: "TAG"},
		},
		LatestOne: []environment.Check{{
			Key: "is_auto_release", Value: "1",
		}},
	}
	if v.Environmental != nil {
		resolver = *v.Environmental
	}
	return resolver
}
