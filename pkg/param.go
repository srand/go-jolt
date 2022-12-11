package jolt

type Parameter interface {
	IsSet() bool
	Set(value interface{}) error
	Get() interface{}
}
