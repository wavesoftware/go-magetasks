package tests_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectBuild(t *testing.T) {
	c := exec.Command("./mage", "clean", "binary")
	c.Dir = "./example"
	err := c.Run()
	assert.NoError(t, err)
	c = exec.Command("./dummy")
	c.Dir = "./example/build/_output/bin"
	err = c.Run()
	assert.NoError(t, err)
}
