package tests_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectBuild(t *testing.T) {
	c := exec.Command("./mage", "clean", "binary")
	c.Dir = "./example"
	c.Stdout = os.Stdout
	err := c.Run()
	assert.NoError(t, err)
	c = exec.Command("./dummy")
	c.Dir = "./example/build/_output/bin"
	c.Stdout = os.Stdout
	err = c.Run()
	assert.NoError(t, err)
}
