package jolt

type StringAttribute struct {
	value string
}

func NewStringAttribute(val string) *StringAttribute {
	return &StringAttribute{val}
}

func (a *StringAttribute) Append(values ...string) {
	for _, val := range values {
		a.value = a.value + val
	}
}

func (a *StringAttribute) Assign(values ...string) {
	a.value = ""
	a.Append(values...)
}

func (a *StringAttribute) Clone() Attribute {
	return NewStringAttribute(a.value)
}

func (a *StringAttribute) Prefix(pfx string) Attribute {
	return &StringAttribute{value: pfx + a.value}
}

func (a *StringAttribute) Suffix(sfx string) Attribute {
	return &StringAttribute{value: a.value + sfx}
}

func (a *StringAttribute) String() string {
	return a.value
}

func (a *StringAttribute) Join(sep string) Attribute {
	return a
}
