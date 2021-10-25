package git

import (
	"context"
	"strings"

	"github.com/blang/semver"
	"github.com/magefile/mage/sh"
	"github.com/wavesoftware/go-ensure"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/version"
)

// Resolver returns a version.Resolver.
func Resolver() version.Resolver {
	return gitInfo{}
}

type cacheKey struct {
	typee string
}

type gitInfo struct{}

func (g gitInfo) Version() string {
	typee := "version-string"
	if ver, ok := fromContext(typee).(string); ok {
		return ver
	}
	ver, err := sh.Output("git", "describe",
		"--always", "--tags", "--dirty")
	ensure.NoError(err)
	saveInContext(typee, ver)
	return ver
}

func (g gitInfo) IsLatest() bool {
	v := g.Version()
	sv, err := semver.ParseTolerant(v)
	ensure.NoError(err)
	for _, t := range resolveTags() {
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

func resolveTags() []string {
	typee := "tags"
	if tags, ok := fromContext(typee).([]string); ok {
		return tags
	}
	tagsOut, err := sh.Output("git", "tag", "--list")
	ensure.NoError(err)
	tags := strings.Split(tagsOut, "\n")
	saveInContext(typee, tags)
	return tags
}

func saveInContext(typ string, version interface{}) {
	config.WithContext(func(ctx context.Context) context.Context {
		return context.WithValue(ctx, cacheKey{typ}, version)
	})
}

func fromContext(typ string) interface{} {
	return config.Actual().Context.Value(cacheKey{typ})
}
