//go:build e2e

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
	if testing.Short() {
		t.Skip("short tests only")
	}
	execCmd(t, "./example", "./mage", "clean", "build")
	execCmd(t, "./example/build/_output/bin", fmt.Sprintf("./other-%s-%s",
		runtime.GOOS, runtime.GOARCH))
}

func execCmd(tb testing.TB, dir, name string, args ...string) {
	tb.Helper()
	c := exec.Command(name, args...)
	c.Env = append(
		env(filterOutByName{names: []string{"GOOS"}}),
		"GOTRACEBACK=all",
	)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	assert.NilError(tb, c.Start())
	tb.Logf("Started `%q` with pid %d",
		append([]string{name}, args...),
		c.Process.Pid)
	assert.NilError(tb, c.Wait())
}

func env(filter envFilter) []string {
	ret := make([]string, 0, len(os.Environ()))
	for _, e := range os.Environ() {
		envPair := strings.SplitN(e, "=", 2)
		key := envPair[0]
		if !filter.include(key) {
			continue
		}

		ret = append(ret, e)
	}
	return ret
}

type filterOutByName struct {
	names []string
}

func (f filterOutByName) include(name string) bool {
	for _, n := range f.names {
		if name == n {
			return false
		}
	}
	return true
}

type envFilter interface {
	include(name string) bool
}
