package jolt

type depsJob struct {
	Job
	job  Job
	deps []JobRef
}

func WithDeps(job Job, deps []Input) Job {
	j := &depsJob{job: job}
	for _, d := range deps {
		j.deps = append(j.deps, JobRef(d.String()))
	}
	return j
}

func (j *depsJob) Deps() []JobRef {
	return append(j.job.Deps(), j.deps...)
}

func (j *depsJob) Execute() ([]Job, error) {
	return j.job.Execute()
}

func (j *depsJob) Refs() []JobRef {
	return j.job.Refs()
}

func (j *depsJob) Digest() (Digest, error) {
	return j.job.Digest()
}
