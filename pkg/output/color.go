package output

import (
	"os"

	"github.com/fatih/color"
)

// SetupColorMode will set the output color mode.
func SetupColorMode() {
	if val, envset := os.LookupEnv("FORCE_COLOR"); envset && val == "true" {
		color.NoColor = false
	}
}
