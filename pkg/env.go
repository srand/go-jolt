package jolt

import (
	"bytes"
	"log"
	"strings"
	"text/template"
)

type Environment struct {
	Name       string
	Attributes map[string]Attribute
	Parameters map[string]Parameter
	Rules      map[string]Rule
	RulesByExt map[string][]Rule
	Tasks      map[string]*Task
}

func NewEnv() *Environment {
	e := &Environment{}
	e.reset()
	return e
}

func (e *Environment) reset() {
	e.Attributes = map[string]Attribute{}
	e.Parameters = map[string]Parameter{}
	e.Rules = map[string]Rule{}
	e.RulesByExt = map[string][]Rule{}
	e.Tasks = map[string]*Task{}
}

func (e *Environment) assign(inherited *Environment) {
	for name, attr := range inherited.Attributes {
		if _, ok := e.Attributes[name]; ok {
			for _, value := range attr.(*ListAttribute).Items() {
				e.Attributes[name].Append(value)
			}
		} else {
			e.Attributes[name] = attr
		}
	}

	for name, param := range inherited.Parameters {
		e.Parameters[name] = param
	}

	for name, rule := range inherited.Rules {
		e.Rules[name] = rule
	}

	for ext, rules := range inherited.RulesByExt {
		e.RulesByExt[ext] = append(e.RulesByExt[ext], rules...)
	}

	for name, task := range inherited.Tasks {
		e.Tasks[name] = task
	}
}

func (e *Environment) copy(inherited *Environment) {
	for name, attr := range inherited.Attributes {
		e.Attributes[name] = attr.Clone()
	}

	for name, param := range inherited.Parameters {
		e.Parameters[name] = param
	}

	for ext, rule := range inherited.Rules {
		e.Rules[ext] = rule
	}

	for ext, rules := range inherited.RulesByExt {
		e.RulesByExt[ext] = append(e.RulesByExt[ext], rules...)
	}

	for name, task := range inherited.Tasks {
		e.Tasks[name] = task
	}
}

func (e *Environment) Inherit(parent *Environment) *Environment {
	result := NewEnv()
	result.copy(parent)
	result.assign(e)
	e.reset()
	e.assign(result)
	return e
}

func (e *Environment) Append(name string, values []string) *Environment {
	if attr, ok := e.Attributes[name]; ok {
		attr.(*ListAttribute).Append(values...)
	} else {
		log.Fatal("No such attribute: ", name)
	}
	return e
}

func (e *Environment) Render(templ string, extra map[string]interface{}) string {
	outp := bytes.Buffer{}
	tmpl := template.Must(template.New("tmpl").Parse(templ))
	obj := map[string]interface{}{
		"Attributes": e.Attributes,
	}
	for key, value := range extra {
		obj[key] = value
	}
	tmpl.Execute(&outp, obj)
	return strings.TrimSpace(outp.String())
}
