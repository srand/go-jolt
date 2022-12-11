package jolt

import (
	"bytes"
	"os/exec"
	"text/template"
)

func RenderTemplate(format string, obj interface{}) string {
	outp := bytes.Buffer{}
	tmpl := template.Must(template.New("tmpl").Parse(format))
	tmpl.Execute(&outp, obj)
	return outp.String()
}

func RunCommand(executable string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(executable, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
