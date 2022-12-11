package jolt

type JobRef string

type Job interface {
	Influence
	Deps() []JobRef
	Execute() ([]Job, error)
	Refs() []JobRef
}

type GeneratorJob interface {
	Job
}

type BasicJob struct {
	AggregatedInfluence
	refs []JobRef
	deps []JobRef
}

func (j *BasicJob) AddDep(ref JobRef) {
	j.deps = append(j.deps, ref)
}

func (j *BasicJob) AddRef(ref JobRef) {
	j.refs = append(j.refs, ref)
}

func (j *BasicJob) Deps() []JobRef {
	return j.deps
}

func (j *BasicJob) Refs() []JobRef {
	return j.refs
}
