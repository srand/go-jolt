package jolt

import (
	"os"
)

type MkdirJob struct {
	BasicJob
	dirs []Input
}

func (j *MkdirJob) AddDir(dir Input) {
	j.dirs = append(j.dirs, dir)
}

func (j *MkdirJob) Execute() ([]Job, error) {
	for _, dir := range j.dirs {
		if err := os.MkdirAll(dir.String(), 0o755); err != nil {
			return []Job{}, err
		}
	}
	return []Job{}, nil
}
