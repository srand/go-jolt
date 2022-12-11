package jolt

type outputAttribute struct {
	ListAttribute
}

func NewOutputAttribute(output []Input) *ListAttribute {
	attr := NewListAttribute()
	for _, o := range output {
		attr.Append(o.String())
	}
	return attr
}
