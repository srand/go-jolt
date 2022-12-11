package jolt

import "strings"

type ListAttribute struct {
	items []string
	glob  bool
}

func NewListAttribute(val ...string) *ListAttribute {
	return &ListAttribute{items: val}
}

func NewGlobListAttribute(val ...string) *ListAttribute {
	return &ListAttribute{items: val, glob: true}
}

func (a *ListAttribute) Append(val ...string) {
	a.items = append(a.items, val...)
}

func (a *ListAttribute) Clone() Attribute {
	return NewListAttribute(a.items...)
}

func (a *ListAttribute) Assign(val ...string) {
	a.items = []string{}
	a.items = append(a.items, val...)
}

func (a *ListAttribute) Items() []string {
	return a.items
}

func (a *ListAttribute) Prefix(pfx string) Attribute {
	copy := &ListAttribute{}
	for _, item := range a.items {
		copy.items = append(copy.items, pfx+item)
	}
	return copy
}

func (a *ListAttribute) Suffix(sfx string) Attribute {
	copy := &ListAttribute{}
	for _, item := range a.items {
		copy.items = append(copy.items, item+sfx)
	}
	return copy
}

func (a *ListAttribute) String() string {
	return a.Join(" ").String()
}

func (a *ListAttribute) Join(sep string) Attribute {
	return &StringAttribute{value: strings.Join(a.items, sep)}
}
