package config

import (
	"context"

	"github.com/fatih/color"
)

// FillInDefaultValues in provided config and returns a filled one.
func FillInDefaultValues(cfg Config) Config {
	if len(cfg.BuildDirPath) == 0 {
		cfg.BuildDirPath = []string{"build", "_output"}
	}
	empty := &MageTag{}
	if cfg.MageTag == *empty {
		cfg.MageTag = MageTag{
			Color: color.FgCyan,
			Label: "[MAGE]",
		}
	}
	if cfg.Dependencies == nil {
		cfg.Dependencies = NewDependencies("github.com/kyoh86/richgo")
	}
	if cfg.Context == nil {
		cfg.Context = context.TODO()
	}
	if cfg.Artifacts == nil {
		cfg.Artifacts = make(map[string]Artifact)
	}
	return cfg
}
