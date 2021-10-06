package git

import (
	"context"

	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks/config"
)

type cacheKey struct {
	ref interface{}
	key string
}

// Version returns a git version string.
func Version() string {
	if version, ok := fromContext(); ok {
		return version
	}
	version, err := sh.Output("git", "describe",
		"--always", "--tags", "--dirty")
	ensure.NoError(err)
	saveInContext(version)
	return version
}

func gitVersionCacheKey() cacheKey {
	return cacheKey{
		ref: config.Actual(), key: "ctx.gitVersion",
	}
}

func saveInContext(version string) {
	config.WithContext(func(ctx context.Context) context.Context {
		key := gitVersionCacheKey()
		return context.WithValue(ctx, key, version)
	})
}

func fromContext() (string, bool) {
	key := gitVersionCacheKey()
	ver, ok := config.Actual().Context.Value(key).(string)
	return ver, ok
}
