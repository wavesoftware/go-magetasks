package magetasks

import (
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/output"
)

// Configure will set up a project to be built.
func Configure(cfg config.Config) {
	cfg = config.FillInDefaultValues(cfg)
	p := project{cfg: &cfg}
	output.SetupColorMode()
	config.Configure(p)
}

type project struct {
	cfg *config.Config
}

func (p project) Config() *config.Config {
	return p.cfg
}
