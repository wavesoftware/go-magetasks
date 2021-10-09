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
	execCmd(t, "./example", "./mage", "clean", "build")
	execCmd(t, "./example/build/_output/bin", fmt.Sprintf("./other-%s-%s",
		runtime.GOOS, runtime.GOARCH))
}

func execCmd(tb testing.TB, dir, name string, args ...string) {
	tb.Helper()
	c := exec.Command(name, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	assert.NilError(tb, err)
}
