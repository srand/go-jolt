package jolt

import (
	"fmt"
)

type StringParam struct {
	Name     string
	Value    string
	Accepted []string
	Usage    string
}

func NewStringParam(name, value string, accepted []string, usage string) Parameter {
	p := &StringParam{
		Name:     name,
		Accepted: accepted,
		Usage:    usage,
	}
	if err := p.Set(value); value != "" && err != nil {
		return nil
	}
	return p
}

func (p *StringParam) validate(value string) error {
	if len(p.Accepted) > 0 {
		for _, acc := range p.Accepted {
			if value == acc {
				return nil
			}
		}
		return fmt.Errorf("Illegal value: %v", value)
	}
	return nil
}

func (p *StringParam) Set(value interface{}) error {
	var val string
	var ok bool

	if val, ok = value.(string); !ok {
		return fmt.Errorf("Invalid argument: %v", value)
	}

	if err := p.validate(val); err != nil {
		return err
	}

	p.Value = val

	return nil
}

func (p *StringParam) Get() interface{} {
	return p.Value
}

func (p *StringParam) IsSet() bool {
	return p.Value != ""
}
