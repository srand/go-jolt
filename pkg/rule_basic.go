package jolt

type BasicRule struct {
	Aggregate bool
	Mandatory bool
	Message   string
	Phony     bool
}

func (r *BasicRule) IsAggregate() bool {
	return r.Aggregate
}

func (r *BasicRule) IsMandatory() bool {
	return r.Mandatory
}

func (r *BasicRule) IsPhony() bool {
	return r.Phony
}
