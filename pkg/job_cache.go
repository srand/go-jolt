package jolt

import (
	"fmt"
)

const (
	DefaultCacheLocation = ""
)

type JobCache interface {
	Add(job Job) error
	Get(job Job) (string, error)
	Has(job Job) (bool, error)
}

type LocalJobCache struct {
	Path string
}

func NewLocalJobCache() JobCache {
	return &LocalJobCache{}
}

func (c *LocalJobCache) Add(job Job) error {
	return fmt.Errorf("")
}

func (c *LocalJobCache) Get(job Job) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func (c *LocalJobCache) Has(job Job) (bool, error) {
	return false, fmt.Errorf("Not implemented")
}

type CachedProxyJob struct {
	Job
	job Job
}

func (j *CachedProxyJob) Digest() (Digest, error) {
	return j.job.Digest()
}

func (j *CachedProxyJob) Deps() []JobRef {
	return j.job.Deps()
}

func (j *CachedProxyJob) execute() ([]Job, error) {
	jobs, err := j.job.Execute()
	if err != nil {
		return []Job{}, err
	}

	newDigest, err := j.Digest()
	if err != nil {
		return []Job{}, fmt.Errorf("No digest available after running job: %w", err)
	}

	for _, ref := range j.Refs() {
		SetXattrDigest(string(ref), XattrJobDigest, newDigest)
	}

	return jobs, err
}

func (j *CachedProxyJob) Execute() ([]Job, error) {
	newDigest, err := j.job.Digest()
	if err != nil {
		return j.execute()
	}

	jobs := []Job{}
	rebuild := false

	refs := j.job.Refs()
	if len(refs) <= 0 {
		rebuild = true
	}

	for _, ref := range refs {
		oldDigest, err := GetXattrDigest(string(ref), XattrJobDigest)
		if err != nil {
			rebuild = true
			break
		}
		if oldDigest != newDigest {
			rebuild = true
			break
		}
	}

	if rebuild {
		return j.execute()
	}

	return jobs, nil
}

func (j *CachedProxyJob) Refs() []JobRef {
	return j.job.Refs()
}
