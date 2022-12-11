package jolt

type BoolAttribute struct {
	value bool
}

func NewBoolAttribute(val bool) *BoolAttribute {
	return &BoolAttribute{val}
}

func (a *BoolAttribute) Append(values ...string) {
}

func (a *BoolAttribute) Assign(values ...string) {
	for _, value := range values {
		switch value {
		case "true":
			a.value = true
		case "false":
			a.value = false
		}
	}
}

func (a *BoolAttribute) Clone() Attribute {
	return NewBoolAttribute(a.value)
}

func (a *BoolAttribute) IsTrue() bool {
	return a.value
}

func (a *BoolAttribute) Prefix(pfx string) Attribute {
	return a
}

func (a *BoolAttribute) Suffix(sfx string) Attribute {
	return a
}

func (a *BoolAttribute) String() string {
	if a.value {
		return "true"
	} else {
		return "false"
	}
}

func (a *BoolAttribute) Join(sep string) Attribute {
	return a
}
