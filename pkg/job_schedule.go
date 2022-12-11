package jolt

import (
	"log"
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
	for {
		// s.Index.Dump()

		jobs := s.Index.Ready()
		if len(jobs) == 0 {
			s.done <- true
			return
		}

		for _, job := range jobs {
			cache_proxy_job := &CachedProxyJob{job: job}

			jobs, err := cache_proxy_job.Execute()
			if err != nil {
				log.Fatal(err)
			}

			for _, newjob := range jobs {
				s.Index.Add(newjob)
			}

			s.Index.Remove(job)

			// s.Index.Dump()
		}
	}
}

func (s *readyJobSchedule) Done() chan bool {
	return s.done
}
