package jolt

type EnvironmentIndex map[string]*Environment

func NewEnvIndex() *EnvironmentIndex {
	return &EnvironmentIndex{}
}
