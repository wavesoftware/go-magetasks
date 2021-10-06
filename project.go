package magetasks

import (
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/output"
)

// Project will set up a project to be built.
func Project(cfg config.Config) config.Configured {
	cfg = config.FillInDefaultValues(cfg)
	p := project{cfg: &cfg}
	output.SetupColorMode()
	return config.Configure(p)
}

type project struct {
	cfg *config.Config
}

func (p project) Config() *config.Config {
	return p.cfg
}
