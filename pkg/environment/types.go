package environment

import (
	"os"
	"strings"
)

// Key is an environment key.
type Key string

// Value is an environment value.
type Value string

// Pair holds a pair of environment key and value.
type Pair struct {
	Key
	Value
}

// Values holds environment values together with their keys.
type Values map[Key]Value

// Add a pair to environment values.
func (v Values) Add(pair Pair) {
	v[pair.Key] = pair.Value
}

// Current returns current environment values, from os.Environ method.
func Current() Values {
	vals := Values(map[Key]Value{})
	for _, pair := range os.Environ() {
		vals.Add(NewPair(pair))
	}
	return vals
}

// NewPair creates a pair from os.Environ style string.
func NewPair(environ string) Pair {
	parts := strings.SplitN(environ, "=", 1)
	return Pair{
		Key:   Key(parts[0]),
		Value: Value(parts[1]),
	}
}
