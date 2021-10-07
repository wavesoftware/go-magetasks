package tasks

import (
	"github.com/fatih/color"
	"github.com/wavesoftware/go-magetasks/config"
)

func mageTag() string {
	mt := config.Actual().MageTag
	return color.New(mt.Color).Sprint(mt.Label)
}
