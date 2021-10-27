package git

import (
	"errors"

	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/cache"
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

// IsLatestStrategy is used to determine if current version is latest one.
type IsLatestStrategy func(Resolver) func(versionRange string) (bool, error)

// Remote represents a remote repository name and address.
type Remote struct {
	Name string
	URL  string
}

// Resolver implements version.Resolver for git SCM.
type Resolver struct {
	Cache cache.Cache
	IsLatestStrategy
	Repository
	*Remote
}

type cacheKey struct {
	typee string
}

func (r Resolver) Version() string {
	ver, err := r.cache().Compute(cacheKey{"version"}, func() (interface{}, error) {
		return r.repository().Describe()
	})
	ensure.NoError(err)
	return ver.(string)
}

func (r Resolver) IsLatest(versionRange string) (bool, error) {
	strategy := TagBasedIsLatestStrategy
	if r.IsLatestStrategy != nil {
		strategy = r.IsLatestStrategy
	}
	fn := strategy(r)
	latest, err := fn(versionRange)
	if err != nil {
		if !errors.Is(err, version.ErrVersionIsNotValid) {
			return false, err
		}
		return false, nil
	}
	return latest, nil
}

func (r Resolver) cache() cache.Cache {
	if r.Cache == nil {
		return config.Cache()
	}
	return r.Cache
}

func (r Resolver) repository() Repository {
	if r.Repository == nil {
		return installedGitBinaryRepo{r.remote()}
	}
	return r.Repository
}

func (r Resolver) remote() Remote {
	remote := Remote{Name: "origin"}
	if r.Remote != nil {
		remote = *r.Remote
	}
	return remote
}

func (r Resolver) resolveTags() []string {
	tt, err := r.cache().Compute(cacheKey{"tags"}, func() (interface{}, error) {
		return r.repository().Tags()
	})
	ensure.NoError(err)
	return tt.([]string)
}
