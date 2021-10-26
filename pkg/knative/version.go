package knative

import (
	"github.com/wavesoftware/go-magetasks/pkg/git"
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

// VersionResolver is version resolver implementation directly targeting
// Knative project CI.
type VersionResolver struct {
	Git           *git.Resolver
	Environmental *version.EnvironmentResolver
}

func (v VersionResolver) Version() string {
	return v.resolver().Version()
}

func (v VersionResolver) IsLatest() bool {
	return v.resolver().IsLatest()
}

func (v VersionResolver) resolver() version.Resolver {
	gr := git.Resolver{}
	if v.Git != nil {
		gr = *v.Git
	}
	er := version.EnvironmentResolver{
		VersionKey: "TAG",
		IsApplicable: []version.Check{
			{Key: "TAG_RELEASE", Value: "1"},
			{Key: "TAG"},
		},
		LatestOne: []version.Check{{
			Key: "is_auto_release", Value: "1",
		}},
	}
	if v.Environmental != nil {
		er = *v.Environmental
	}
	cr := version.CompositeResolver{
		VersionResolver:  er,
		IsLatestResolver: gr,
	}
	return version.OrderedResolver{
		Resolvers: []version.Resolver{cr, gr},
	}
}
