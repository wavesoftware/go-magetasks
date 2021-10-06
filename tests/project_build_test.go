package tests_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"testing"

	"gotest.tools/v3/assert"
)

func TestProjectBuild(t *testing.T) {
	c := exec.Command("./mage", "clean", "build")
	c.Dir = "./example"
	c.Stdout = os.Stdout
	err := c.Run()
	assert.NilError(t, err)
	c = exec.Command(fmt.Sprintf("./dummy-%s-%s",
		runtime.GOOS, runtime.GOARCH))
	c.Dir = "./example/build/_output/bin"
	c.Stdout = os.Stdout
	err = c.Run()
	assert.NilError(t, err)
}
