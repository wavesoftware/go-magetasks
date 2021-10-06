package config

import (
	"context"

	"github.com/fatih/color"
)

// Resolver is a func that resolves to a string.
type Resolver func() string

// MageTag holds a mage tag.
type MageTag struct {
	Color color.Attribute
	Label string
}

type BuildResult struct {
	Error error
	Info  map[string]string
}

func (r BuildResult) Failed() bool {
	return r.Error != nil
}

type Artifact interface {
	Build(name string) BuildResult
}

type Version struct {
	Path string
	Resolver
}

type Metadata struct {
	Args map[string]Resolver
}

// Task is a custom function that will be used in the build.
type Task struct {
	Name      string
	Operation func() error
	Overrides []Configurator
}

// Config holds configuration information.
type Config struct {
	// ProjectDir holds a path to the project.
	ProjectDir string

	// BuildDirPath holds a build directory path.
	BuildDirPath []string

	// Version contains the version information.
	*Version

	// MageTag holds default mage tag settings.
	MageTag

	// Dependencies will hold additional Golang dependencies that needs to be
	// installed before running tasks.
	Dependencies Dependencies

	// Artifacts holds a list of artifacts to be built.
	Artifacts map[string]Artifact

	// Cleaning additional cleaning tasks.
	Cleaning []Task

	// Checks holds a list of checks to perform.
	Checks []Task

	// Overrides holds a list of overrides of this configuration.
	Overrides []Configurator

	// Default task to run.
	Default func()

	// context.Context is standard Golang context.
	context.Context
}

// Configurator will configure project.
type Configurator interface {
	Configure(cfg Configurable)
}

// Configurable will allow changes of the Config structure.
type Configurable interface {
	Config() *Config
}

// Configured represents a configured project.
type Configured interface {
	Default() func()
}
