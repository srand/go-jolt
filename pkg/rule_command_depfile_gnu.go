package jolt

type DepfileCommandRule struct {
	CommandRule
	Depfile string
}

func (r *DepfileCommandRule) Build(env *Environment, inputs []Input) ([]Job, error) {
	alljobs := []Job{}

	if r.IsAggregate() {
		jobs, err := r.CommandRule.Build(env, inputs)
		if err != nil {
			return []Job{}, err
		}

		outputs, err := r.Outputs(env, inputs)
		if err != nil {
			return []Job{}, err
		}

		outputStrings := []string{}
		for _, output := range outputs {
			outputStrings = append(outputStrings, output.String())
		}

		depfile := env.Render(r.Depfile, map[string]interface{}{
			"Input":  inputs,
			"Output": NewOutputAttribute(outputs),
		})

		influence := NewGnuDepfileInfluence(depfile, outputStrings)

		for _, job := range jobs {
			job.(*AggregatedJob).AddInfluence(influence)
		}

		alljobs = append(alljobs, jobs...)
	} else {
		for _, input := range inputs {
			inputs := []Input{input}

			jobs, err := r.CommandRule.Build(env, inputs)
			if err != nil {
				return []Job{}, err
			}

			outputs, err := r.Outputs(env, inputs)
			if err != nil {
				return []Job{}, err
			}

			outputStrings := []string{}
			for _, output := range outputs {
				outputStrings = append(outputStrings, output.String())
			}

			depfile := env.Render(r.Depfile, map[string]interface{}{
				"Input":  input,
				"Output": NewOutputAttribute(outputs),
			})

			influence := NewGnuDepfileInfluence(depfile, outputStrings)

			for _, job := range jobs {
				job.(*AggregatedJob).AddInfluence(influence)
			}

			alljobs = append(alljobs, jobs...)
		}
	}

	return alljobs, nil
}
