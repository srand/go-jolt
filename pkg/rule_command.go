package jolt

type CommandRule struct {
	BasicRule
	Command string
	Output  []string
}

func (r *CommandRule) Build(env *Environment, inputs []Input) ([]Job, error) {
	jobs := []Job{}

	if r.Aggregate {
		outputs, err := r.Outputs(env, inputs)
		if err != nil {
			return jobs, err
		}

		cmd := env.Render(r.Command, map[string]interface{}{
			"Input":  inputs,
			"Output": NewOutputAttribute(outputs),
		})
		msg := env.Render(r.Message, map[string]interface{}{
			"Input":  inputs,
			"Output": NewOutputAttribute(outputs),
		})
		aggJob := &AggregatedJob{}
		cmdJob := NewCommandJob(cmd)
		cmdJob.Message = msg
		for _, input := range inputs {
			cmdJob.AddDep(JobRef(input.String()))
		}
		for _, output := range outputs {
			cmdJob.AddRef(JobRef(output.String()))
			if output.Dir().String() != "" {
				mkdir := &MkdirJob{}
				mkdir.AddDir(output.Dir())
				aggJob.AddJob(mkdir)
			}
		}
		aggJob.AddJob(cmdJob)
		jobs = append(jobs, aggJob)
	} else {
		for _, input := range inputs {
			outputs, err := r.Outputs(env, []Input{input})
			if err != nil {
				return jobs, err
			}

			cmd := env.Render(r.Command, map[string]interface{}{
				"Input":  input,
				"Output": NewOutputAttribute(outputs),
			})
			msg := env.Render(r.Message, map[string]interface{}{
				"Input":  input,
				"Output": NewOutputAttribute(outputs),
			})

			aggJob := &AggregatedJob{}
			cmdJob := NewCommandJob(cmd)
			cmdJob.Message = msg
			cmdJob.AddDep(JobRef(input.String()))

			for _, output := range outputs {
				cmdJob.AddRef(JobRef(output.String()))
				if output.Dir().String() != "" {
					mkdir := &MkdirJob{}
					mkdir.AddDir(output.Dir())
					aggJob.AddJob(mkdir)
				}
			}

			aggJob.AddJob(cmdJob)
			jobs = append(jobs, aggJob)
		}
	}

	return jobs, nil
}

func (r *CommandRule) Outputs(env *Environment, inputs []Input) ([]Input, error) {
	outputs := []Input{}

	if r.Aggregate {
		for _, output := range r.Output {
			output = env.Render(output, map[string]interface{}{
				"Input": inputs,
			})
			outputs = append(outputs, &File{output})
		}
	} else {
		for _, input := range inputs {
			for _, output := range r.Output {
				output = env.Render(output, map[string]interface{}{
					"Input": input,
				})
				outputs = append(outputs, &File{output})
			}
		}
	}

	return outputs, nil
}
