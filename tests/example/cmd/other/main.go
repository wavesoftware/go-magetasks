package main

import (
	"fmt"

	"github.com/wavesoftware/go-magetasks/tests/example/pkg/metadata"
)

func main() {
	fmt.Printf("Version: %s\nImage: %s\n",
		metadata.Version, metadata.Image)
}
