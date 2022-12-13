package jolt

import (
	"bufio"
	"io"
	"log"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func init() {
	log.SetFlags(0)
}

type logger struct {
	input  io.Reader
	reader *bufio.Reader
	pfx    string
}

func (l *logger) Dispatch() {
	for {
		line, _, err := l.reader.ReadLine()
		if err != nil {
			return
		}
		log.Default().Println(l.pfx + string(line) + Reset)
	}
}

func LogStdout(stdout io.Reader) {
	log := &logger{pfx: Reset + "STDOUT - "}
	log.reader = bufio.NewReader(stdout)
	log.Dispatch()
}

func LogStderr(stderr io.Reader) {
	log := &logger{pfx: Reset + Red + "STDERR - "}
	log.reader = bufio.NewReader(stderr)
	log.Dispatch()
}
