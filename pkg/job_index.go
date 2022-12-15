package jolt

import (
	"fmt"
	"strings"
	"sync"
)

type jobrec struct {
	job      Job
	deps     *StringSet
	refs     *StringSet
	doneChan chan bool
}

type JobIndex struct {
	RefToJob map[JobRef]*jobrec
	DepToJob map[JobRef][]*jobrec
	ready    map[*jobrec]bool
	mutex    sync.Mutex
}

func NewJobIndex() *JobIndex {
	return &JobIndex{
		RefToJob: map[JobRef]*jobrec{},
		DepToJob: map[JobRef][]*jobrec{},
		ready:    map[*jobrec]bool{},
	}
}

func (i *JobIndex) Add(job Job) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	jobrec := &jobrec{
		job:  job,
		deps: NewStringSet(job.Deps()),
		refs: NewStringSet(job.Refs()),
	}

	for _, ref := range job.Refs() {
		i.RefToJob[ref] = jobrec
	}

	deps := job.Deps()
	for _, dep := range deps {
		i.DepToJob[dep] = append(i.DepToJob[dep], jobrec)
	}

	if len(deps) == 0 {
		i.ready[jobrec] = true
	}

	return nil
}

func (i *JobIndex) Remove(job Job) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	for _, ref := range job.Refs() {
		delete(i.RefToJob, ref)

		parents, ok := i.DepToJob[ref]
		if !ok {
			continue
		}
		delete(i.DepToJob, ref)

		for _, parent := range parents {
			parent.deps.Delete(string(ref))
			if len(*parent.deps) == 0 {
				i.ready[parent] = true
			}
		}
	}
}

func (i *JobIndex) IsEmpty() bool {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	return len(i.RefToJob) == 0 && len(i.ready) == 0
}

func (i *JobIndex) Ready() []Job {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	jobs := []Job{}
	for rec := range i.ready {
		jobs = append(jobs, rec.job)
	}
	i.ready = map[*jobrec]bool{}
	return jobs
}

func (i *JobIndex) Dump() {
	for _, job := range i.RefToJob {
		refs := job.job.Refs()
		deps := job.job.Deps()
		refstr := []string{}
		depstr := []string{}

		for _, ref := range refs {
			refstr = append(refstr, string(ref))
		}
		for _, dep := range deps {
			depstr = append(depstr, string(dep))
		}
		fmt.Println(strings.Join(refstr, " "), ": ", strings.Join(depstr, " "))
	}
}
