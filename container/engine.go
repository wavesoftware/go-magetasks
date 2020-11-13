package container

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/wavesoftware/go-magetasks/config"
)

var (
	errNoContainerEngineInstalled = errors.New(
		"can't find a installed container engine (podman or docker)",
	)
)

func containerFile(bin config.Binary) string {
	cf := getenv("CONTAINER_FILENAME", "Containerfile")
	return path.Join("cmd", bin.Name, cf)
}

func imageName(bin config.Binary) string {
	defBasename := fmt.Sprintf("localhost/%s", getenv("USER", "Anonymous"))
	basename := getenv("CONTAINER_BASENAME", defBasename)
	if !strings.HasSuffix(basename, "/") {
		basename += "/"
	}
	return basename + bin.Name
}

func resolveContainerEngine() (string, error) {
	candidates := []string{
		"podman", "docker",
	}
	for _, candidate := range candidates {
		path, err := exec.LookPath(candidate)
		if err == nil {
			return path, nil
		}
	}
	return "", errNoContainerEngineInstalled
}

func getenv(key, defValue string) string {
	val, found := os.LookupEnv(key)
	if !found {
		val = defValue
	}
	return val
}
