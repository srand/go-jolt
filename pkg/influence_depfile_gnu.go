package jolt

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Depfile map[string][]string

type GnuDepfileInfluence struct {
	path    string
	outputs []string
}

func NewGnuDepfileInfluence(path string, outputs []string) Influence {
	return &GnuDepfileInfluence{path: path, outputs: outputs}
}

func ParseDepfile(reader io.Reader) (Depfile, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	content := string(data)
	content = strings.TrimSpace(content)
	content = strings.Replace(content, "\\\n", "", -1)
	deplines := strings.Split(content, "\n")

	deps := make(Depfile)

	for _, dep := range deplines {
		dep = strings.TrimSpace(dep)
		if dep == "" {
			continue
		}

		outputs, inputs, ok := strings.Cut(dep, ": ")
		if !ok {
			return nil, fmt.Errorf("Parse error")
		}
		outputs = strings.TrimSpace(outputs)
		outputs = strings.Replace(outputs, "\\ ", "\x00", -1)

		inputs = strings.TrimSpace(inputs)
		inputs = strings.Replace(inputs, "\\ ", "\x00", -1)

		for _, output := range strings.Fields(outputs) {
			output = strings.Replace(output, "\x00", " ", -1)
			for _, input := range strings.Fields(inputs) {
				deps[output] = append(deps[output], strings.Replace(input, "\x00", " ", -1))
			}
		}
	}

	return deps, nil
}

func (fi *GnuDepfileInfluence) Digest() (Digest, error) {
	fd, err := os.Open(fi.path)
	if err != nil {
		return "", PendingError
	}
	defer fd.Close()

	deps, err := ParseDepfile(fd)
	if err != nil {
		return "", err
	}

	influence := NewAggregatedInfluence()
	for _, output := range fi.outputs {
		inputs, ok := deps[output]
		if !ok {
			return "", PendingError
		}

		for _, input := range inputs {
			if _, err := os.Stat(input); err != nil {
				return "", PendingError
			}
			influence.AddInfluence(FileInfluence(input))
		}
	}

	return influence.Digest()
}
