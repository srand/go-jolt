package jolt

import "fmt"

type BoolParam struct {
	Name  string
	Value bool
	Usage string
}

func NewBoolParam(name string, value bool, usage string) Parameter {
	return &BoolParam{
		Name:  name,
		Value: value,
		Usage: usage,
	}
}

func (p *BoolParam) Set(value interface{}) error {
	if val, ok := value.(bool); ok {
		p.Value = val
		return nil
	}
	if val, ok := value.(string); ok {
		switch val {
		case "true":
			p.Value = true
		case "yes":
			p.Value = true
		case "false":
			p.Value = false
		case "no":
			p.Value = false
		default:
			return fmt.Errorf("Illegal value")
		}
		return nil
	}
	return fmt.Errorf("Illegal type in assignment")
}

func (p *BoolParam) Get() interface{} {
	return p.Value
}

func (p *BoolParam) IsSet() bool {
	return true
}
