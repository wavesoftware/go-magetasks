package magetasks

import (
	"github.com/magefile/mage/mg"
	"github.com/wavesoftware/go-magetasks/config"
	"github.com/wavesoftware/go-magetasks/pkg/deps"
	"github.com/wavesoftware/go-magetasks/pkg/tasks"
)

// Check will run all lints checks.
func Check() {
	mg.Deps(deps.Install)
	t := tasks.StartMultiline("🔍", "Checking")
	for _, check := range config.Actual().Checks {
		p := t.Part(check.Name)
		ps := p.Starting()
		ps.Done(check.Operation())
	}
	t.End(nil)
}
