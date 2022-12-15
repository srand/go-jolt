package jolt

import (
	"context"
	"log"
	"runtime"

	"golang.org/x/sync/semaphore"
)

type JobSchedule interface {
	Dispatch()
	Done() chan bool
}

type readyJobSchedule struct {
	Index *JobIndex
	done  chan bool
}

func NewJobSchedule(index *JobIndex) JobSchedule {
	schedule := &readyJobSchedule{
		Index: index,
		done:  make(chan bool),
	}
	return schedule
}

func (s *readyJobSchedule) Dispatch() {
	readyChan := make(chan bool, runtime.NumCPU())
	sem := semaphore.NewWeighted(int64(runtime.NumCPU()))

	execute := func(job Job) {
		ctx := context.TODO()

		if err := sem.Acquire(ctx, 1); err != nil {
			panic(err)
		}
		defer sem.Release(1)

		cache_proxy_job := &CachedProxyJob{job: job}

		jobs, err := cache_proxy_job.Execute()
		if err != nil {
			log.Fatal(err)
		}

		for _, newjob := range jobs {
			s.Index.Add(newjob)
		}

		s.Index.Remove(job)
		readyChan <- true
	}

	for !s.Index.IsEmpty() {
		// s.Index.Dump()

		jobs := s.Index.Ready()
		for _, job := range jobs {
			go execute(job)
		}

		<-readyChan
	}

	s.done <- true
}

func (s *readyJobSchedule) Done() chan bool {
	return s.done
}
