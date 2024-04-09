// Package model contains the project structure definitions.
package model

// Root represents the root of the project structure.
type Root struct {
	Module string `yaml:"module"`
	Root   []*Dir `yaml:"root"`
}

// Dir represents a directory in the project structure.
type Dir struct {
	Name string `yaml:"name"`

	Files []*File `yaml:"files"`
	Dirs  []*Dir  `yaml:"dirs"`
}

// File represents a file in the project structure.
type File struct {
	Name string `yaml:"name"`
}
