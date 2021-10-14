package tests_test

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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
	c.Env = env(func(e string) bool {
		return e == "GOARCH" || e == "GOOS" || e == "GOARM"
	})
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	assert.NilError(tb, err)
}

func env(filter func(string) bool) []string {
	ret := make([]string, 0, len(os.Environ()))
	for _, e := range os.Environ() {
		envPair := strings.SplitN(e, "=", 2)
		key := envPair[0]
		if filter(key) {
			continue
		}

		ret = append(ret, e)
	}
	return ret
}
