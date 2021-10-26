package git

import (
	"github.com/blang/semver/v4"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/cache"
)

// IsLatestStrategy is used to determine if current version is latest one.
type IsLatestStrategy func(Resolver) bool

// Remote represents a remote repository name and address.
type Remote struct {
	Name string
	URL  string
}

// Resolver implements version.Resolver for git SCM.
type Resolver struct {
	Cache cache.Cache
	IsLatestStrategy
	Info
	*Remote
}

type cacheKey struct {
	typee string
}

func (r Resolver) Version() string {
	ver, err := r.cache().Compute(cacheKey{"version"}, func() (interface{}, error) {
		return r.info().Description()
	})
	ensure.NoError(err)
	return ver.(string)
}

func (r Resolver) IsLatest() bool {
	strategy := defaultIsLatestStrategy
	if r.IsLatestStrategy != nil {
		strategy = r.IsLatestStrategy
	}
	return strategy(r)
}

func (r Resolver) cache() cache.Cache {
	if r.Cache == nil {
		return config.Cache()
	}
	return r.Cache
}

func (r Resolver) info() Info {
	if r.Info == nil {
		return shellInfo{r.remote()}
	}
	return r.Info
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
		return r.info().Tags()
	})
	ensure.NoError(err)
	return tt.([]string)
}

func defaultIsLatestStrategy(r Resolver) bool {
	v := r.Version()
	sv, err := semver.ParseTolerant(v)
	ensure.NoError(err)
	for _, t := range r.resolveTags() {
		tv, err := semver.ParseTolerant(t)
		if err != nil {
			continue
		}
		if tv.GT(sv) {
			return false
		}
	}
	return true
}
