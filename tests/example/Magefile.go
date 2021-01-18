// +build mage

package main

import (
	"log"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bins := []string{"dummy", "other"}
	for _, bin := range bins {
		config.Binaries = append(config.Binaries, config.Binary{Name: bin})
	}
	config.VersionVariablePath = "github.com/wavesoftware/go-magetasks/tests/example/internal.Version"
}
