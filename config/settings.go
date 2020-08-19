package config

import "github.com/fatih/color"

var (
	RepoDir      string
	BuildDirPath = []string{"build", "_output"}
	MageTag      = MageTagStruct{
		Color: color.FgCyan,
		Label: "[MAGE]",
	}
	Dependencies = []string{
		"github.com/kyoh86/richgo",
		"github.com/mgechev/revive",
		"honnef.co/go/tools/cmd/staticcheck",
	}
	VersionVariablePath string
	Binaries            []Binary
	CleaningTasks       []CleaningTask
)
