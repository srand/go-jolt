package jolt

type AggregatedJob struct {
	AggregatedInfluence
	jobs []Job
}

func (j *AggregatedJob) AddJob(job Job) {
	j.AddInfluence(job)
	j.jobs = append(j.jobs, job)
}

func (j *AggregatedJob) Execute() ([]Job, error) {
	var alljobs []Job
	for _, job := range j.jobs {
		jobs, err := job.Execute()
		if err != nil {
			return []Job{}, err
		}
		alljobs = append(alljobs, jobs...)
	}
	return alljobs, nil
}

func (j *AggregatedJob) Deps() []JobRef {
	deps := []JobRef{}
	for _, job := range j.jobs {
		deps = append(deps, job.Deps()...)
	}
	return deps
}

func (j *AggregatedJob) Refs() []JobRef {
	refs := []JobRef{}
	for _, job := range j.jobs {
		refs = append(refs, job.Refs()...)
	}
	return refs
}
