package jolt

import "strings"

type MapAttribute struct {
	items map[string]string
	keys  []string
}

func NewMapAttribute() *MapAttribute {
	return &MapAttribute{
		items: map[string]string{},
	}
}

func (a *MapAttribute) Append(values ...string) {
	for _, val := range values {
		keyval := strings.Split(val, "=")
		if len(keyval) == 1 {
			a.Set(keyval[0], "")
		} else {
			a.Set(keyval[0], keyval[1])
		}
	}
}

func (a *MapAttribute) Clone() Attribute {
	m := NewMapAttribute()
	for k, v := range a.items {
		a.items[k] = v
	}
	return m
}

func (a *MapAttribute) Assign(values ...string) {
	a.items = map[string]string{}
	a.keys = []string{}
	a.Append(values...)
}

func (a *MapAttribute) Set(key, value string) {
	a.items[key] = value
	a.keys = append(a.keys, key)
}

func (a *MapAttribute) Prefix(pfx string) Attribute {
	return a.Join("=").Prefix(pfx)
}

func (a *MapAttribute) Suffix(sfx string) Attribute {
	return a.Join("=").Suffix(sfx)
}

func (a *MapAttribute) String() string {
	return a.Join("=").Join(" ").String()
}

func (a *MapAttribute) Join(sep string) Attribute {
	strs := []string{}
	for _, key := range a.keys {
		val := a.items[key]
		if val != "" {
			strs = append(strs, key+sep+val)
		} else {
			strs = append(strs, key)
		}
	}
	return &ListAttribute{items: strs}
}
