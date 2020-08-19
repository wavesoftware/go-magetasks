package config

import "github.com/fatih/color"

type MageTagStruct struct {
	Color color.Attribute
	Label string
}

type Binary struct {
	Name string
}

type CleaningTask func() error
