package magetasks

import "github.com/fatih/color"

var (
	CallerDepth  = 2
	BuildDirPath = []string{"build", "_output"}
	MageTagColor = color.FgCyan
	MageTagLabel = "[MAGE]"
	Dependencies = []string{
		"github.com/kyoh86/richgo",
		"github.com/mgechev/revive",
		"honnef.co/go/tools/cmd/staticcheck",
	}
)
