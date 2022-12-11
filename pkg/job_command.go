package jolt

import (
	"log"
	"os/exec"
)

type CommandJob struct {
	BasicJob
	Message string
	Command string
}

func NewCommandJob(command string) *CommandJob {
	cmd := &CommandJob{Command: command}
	cmd.AddInfluence(StringInfluence(command))
	cmd.AddInfluence(DepsInfluence(cmd))
	return cmd
}

func (job *CommandJob) Execute() ([]Job, error) {
	cmd := exec.Command("sh", "-c", job.Command)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return []Job{}, err
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return []Job{}, err
	}

	go func() {
		go LogStdout(stdout)
	}()

	go func() {
		go LogStderr(stderr)
	}()

	if job.Message != "" {
		log.Println(job.Message)
	} else {
		log.Println(job.Command)
	}

	err = cmd.Run()
	if err != nil {
		log.Println(job.Command)
	}

	return []Job{}, err
}
