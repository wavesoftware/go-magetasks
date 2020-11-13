// +build mage

package main

import (
	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"

	// mage:import
	_ "github.com/wavesoftware/go-magetasks/container"
)

// Default target is set to Binary
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Binary

func init() {
	config.Binaries = append(config.Binaries, config.Binary{
		Name: "dummy",
	})
	config.VersionVariablePath = "github.com/wavesoftware/go-magetasks/tests/example/internal.Version"
}
