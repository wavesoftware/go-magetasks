// +build mage

package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wavesoftware/go-magetasks/pkg/checks"

	// mage:import
	"github.com/wavesoftware/go-magetasks"
	"github.com/wavesoftware/go-magetasks/config"

	// mage:import
	_ "github.com/wavesoftware/go-magetasks/container"
)

// Default target is set to Binary
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Binary // nolint:deadcode,unused

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	addBinary(config.Binary{
		Name: "dummy",
		ImageArgs: map[string]string{
			"DESC": "v0.6.9-Final",
		},
	})
	addBinary(config.Binary{Name: "other"})
	config.VersionVariablePath = "github.com/wavesoftware/go-magetasks/tests/example/internal.Version"
	checks.GolangCiLint()
}

func addBinary(bin config.Binary) {
	config.Binaries = append(config.Binaries, bin)
}
