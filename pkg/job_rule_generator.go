package jolt

import "fmt"

type RuleArgs map[string]string

type RuleGeneratorJob struct {
	Influence AggregatedInfluence
	Files     []Input
	Rules     []Rule
	Env       *Environment
}

func NewRuleGeneratorJob(env *Environment, inputs ...string) *RuleGeneratorJob {
	gen := &RuleGeneratorJob{Env: env}

	for _, file := range inputs {
		gen.Files = append(gen.Files, &File{file})
		gen.Influence.AddInfluence(FileInfluence(file))
	}

	return gen
}

func (j *RuleGeneratorJob) Digest() (Digest, error) {
	return j.Influence.Digest()
}

func (j *RuleGeneratorJob) Deps() []JobRef {
	return []JobRef{}
}

func (j *RuleGeneratorJob) Execute() ([]Job, error) {
	jobs := []Job{}

	type SourceRule struct {
		Input  Input
		Origin Rule
	}

	inputs := []SourceRule{}
	for _, file := range j.Files {
		inputs = append(inputs, SourceRule{file, nil})
		jobs = append(jobs, NewInputJob(file))
	}

	var inputs_to_rule = map[Rule][]SourceRule{}

	for _, rule := range j.Env.Rules {
		if rule.IsMandatory() {
			inputs_to_rule[rule] = []SourceRule{}
		}
	}

	for {
		if len(inputs) <= 0 {
			break
		}

		file := inputs[0]
		inputs = inputs[1:]

		rules, ok := j.Env.RulesByExt[file.Input.Ext()]
		if !ok {
			return []Job{}, fmt.Errorf("no rule for file with extension: %s", file.Input.Ext())
		}

		for _, rule := range rules {
			if rule == file.Origin {
				continue
			}

			inputs_to_rule[rule] = append(inputs_to_rule[rule], file)

			outputs, err := rule.Outputs(j.Env, []Input{file.Input})
			if err != nil {
				return []Job{}, err
			}

			for _, output := range outputs {
				inputs = append(inputs, SourceRule{output, rule})
			}
		}
	}

	for rule, inputs := range inputs_to_rule {
		inputInputs := []Input{}
		phonyInputs := []Input{}

		for _, input := range inputs {
			inputInputs = append(inputInputs, input.Input)
			if input.Origin != nil && input.Origin.IsPhony() {
				phonyInputs = append(phonyInputs, input.Input)
			}
		}

		rulejobs, err := rule.Build(j.Env, inputInputs)
		if err != nil {
			return []Job{}, err
		}

		if len(phonyInputs) > 0 {
			for _, job := range rulejobs {
				jobs = append(jobs, WithDeps(job, phonyInputs))
			}
		} else {
			jobs = append(jobs, rulejobs...)
		}
	}

	return jobs, nil
}

func (j *RuleGeneratorJob) Outputs() []string {
	return []string{}
}

func (j *RuleGeneratorJob) Refs() []JobRef {
	return []JobRef{}
}
