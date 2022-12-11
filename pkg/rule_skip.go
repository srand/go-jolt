package jolt

type SkipRule struct {
	BasicRule
}

func (r *SkipRule) Build(env *Environment, input []Input) ([]Job, error) {
	return []Job{}, nil
}

func (r *SkipRule) Outputs(env *Environment, input []Input) ([]Input, error) {
	return []Input{}, nil
}
