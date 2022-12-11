package jolt

type Attribute interface {
	Append(val ...string)
	Assign(val ...string)
	Clone() Attribute
	Join(sep string) Attribute
	Prefix(pfx string) Attribute
	Suffix(sfx string) Attribute
	String() string
}
