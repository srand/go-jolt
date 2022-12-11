package jolt

type InputJob struct {
	BasicJob
	input Input
}

func NewInputJob(input Input) Job {
	j := &InputJob{input: input}
	j.AddRef(JobRef(input.String()))
	j.AddInfluence(FileInfluence(input.String()))
	return j
}

func (j *InputJob) Execute() ([]Job, error) {
	return []Job{}, nil
}
