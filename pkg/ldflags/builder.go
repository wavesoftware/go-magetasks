package ldflags

import (
	"fmt"
	"strings"

	"github.com/wavesoftware/go-magetasks/config"
)

// Builder builds the LD flags by adding values resolvers.
type Builder interface {
	// Add a name and a resolver to the builder.
	Add(name string, resolver config.Resolver) Builder
	// Build onto the args.
	Build(args []string) []string
}

// NewBuilder creates a new builder.
func NewBuilder() Builder {
	return &defaultBuilder{}
}

type defaultBuilder struct {
	resolvers map[string]config.Resolver
}

func (d *defaultBuilder) Add(name string, resolver config.Resolver) Builder {
	d.resolvers[name] = resolver
	return d
}

func (d *defaultBuilder) Build(args []string) []string {
	collected := make([]string, 1, len(d.resolvers)+1)
	collected[0] = "-ldflags=-X"
	for name, resolver := range d.resolvers {
		collected = append(collected, fmt.Sprintf("%s=%s", name, resolver()))
	}
	return append(args, strings.Join(collected, " "))
}
