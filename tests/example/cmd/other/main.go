package main

import (
	"log"

	"github.com/wavesoftware/go-magetasks/tests/example/pkg/metadata"
)

func main() {
	log.Printf("Version: %s\nImage: %s\n",
		metadata.Version, metadata.Image)
}
