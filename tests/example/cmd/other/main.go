package main

import (
	"log"

	"github.com/wavesoftware/go-magetasks/tests/example/pkg/metadata"
)

func main() {
	log.Printf("Version: %s\n", metadata.Version)
	log.Printf("Image: %s\n", metadata.ResolveImage())
}
