package jolt

import (
	"path/filepath"
	"strings"
)

type Input interface {
	Abs() Input
	BareName() Input
	Clean() Input
	Dir() Input
	Ext() string
	Name() Input
	String() string
}

type File struct {
	Path string
}

func (i *File) Abs() Input {
	res, _ := filepath.Abs(i.Path)
	return &File{res}
}

func (i *File) Name() Input {
	return &File{filepath.Base(i.Path)}
}

func (i *File) BareName() Input {
	return &File{strings.TrimSuffix(i.Name().String(), i.Ext())}
}

func (i *File) Dir() Input {
	return &File{filepath.Dir(i.Path)}
}

func (i *File) Ext() string {
	return filepath.Ext(i.Path)
}

func (i *File) Clean() Input {
	return i.clean()
}

func (i *File) clean() *File {
	path := strings.Replace(i.Path, "..", "__", -1)
	path = filepath.Clean(path)
	return &File{path}
}

func (i *File) String() string {
	return i.clean().Path
}

type Rule interface {
	Build(env *Environment, input []Input) ([]Job, error)
	IsAggregate() bool
	IsMandatory() bool
	IsPhony() bool
	Outputs(env *Environment, input []Input) ([]Input, error)
}
