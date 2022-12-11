package jolt

import (
	"fmt"
	"os"
	"path/filepath"
)

type TemplateRule struct {
	BasicRule
	Template string
	Output   []string
}

func (r *TemplateRule) Build(env *Environment, inputs []Input) ([]Job, error) {
	content := env.Render(r.Template, map[string]interface{}{
		"Input": inputs,
	})

	jobs := []Job{}

	outputs, err := r.Outputs(env, inputs)
	if err != nil {
		return jobs, nil
	}

	for _, output := range outputs {
		templ := &TemplateJob{content: content, file: output.String()}
		templ.AddRef(JobRef(output.String()))
		templ.AddInfluence(StringInfluence(content))
		jobs = append(jobs, templ)
	}

	return jobs, nil
}

func (r *TemplateRule) Outputs(env *Environment, inputs []Input) ([]Input, error) {
	outputs := []Input{}

	for _, output := range r.Output {
		output := env.Render(output, map[string]interface{}{
			"Input": inputs,
		})
		outputs = append(outputs, &File{output})
	}
	return outputs, nil
}

type TemplateJob struct {
	BasicJob
	content string
	file    string
}

func (j *TemplateJob) Execute() ([]Job, error) {
	dir := filepath.Dir(j.file)
	if dir != "" {
		err := os.MkdirAll(dir, 0o777)
		if err != nil {
			fmt.Println(j.file)
			return []Job{}, err
		}
	}

	file, err := os.Create(j.file)
	if err != nil {
		return []Job{}, err
	}
	defer file.Close()

	file.WriteString(j.content)

	return []Job{}, nil
}
