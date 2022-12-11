package jolt

import (
	"fmt"
	"io"

	"github.com/srand/go-jolt/pkg/parser"
)

type JoltRecipe struct {
	Env       map[string]*Environment
	Workspace *Environment
}

func NewRecipe() *JoltRecipe {
	return &JoltRecipe{
		Env:       map[string]*Environment{},
		Workspace: &Environment{},
	}
}

type Builder struct {
	recipe *JoltRecipe
}

func NewBuilder() *Builder {
	return &Builder{
		recipe: NewRecipe(),
	}
}

func (r *Builder) Parse(reader io.Reader) (*JoltRecipe, error) {
	file, err := parser.Parse(reader)
	if err != nil {
		return r.recipe, err
	}

	for _, stmt := range file.FileStmt {
		err = r.handleFileStmt(stmt)
		if err != nil {
			return r.recipe, err
		}
	}

	return r.recipe, nil
}

func (r *Builder) handleFileStmt(file *parser.FileStmt) error {
	if file.Environment != nil {
		return r.handleEnvironment(file.Environment)
	}

	if file.Workspace != nil {
		return r.handleWorkspace(file.Workspace)
	}

	return nil
}

func (r *Builder) handleEnvironment(env *parser.Environment) error {
	if _, ok := r.recipe.Env[env.Name]; ok {
		return fmt.Errorf("environment redefined: %s", env.Name)
	}

	res := NewEnv()
	res.Name = env.Name
	r.recipe.Env[env.Name] = res

	for _, stmt := range env.Stmts {
		var err error

		if stmt.Assign != nil {
			err = r.handleAssign(res, stmt.Assign)
		}

		if stmt.Decl != nil {
			err = r.handleDecl(res, stmt.Decl)
		}

		if stmt.Inherit != nil {
			err = r.handleInherit(res, stmt.Inherit)
		}

		if stmt.Project != nil {
			err = r.handleProject(res, stmt.Project)
		}

		if stmt.Rule != nil {
			err = r.handleRule(res, stmt.Rule)
		}

		if stmt.Switch != nil {
			err = r.handleSwitch(res, stmt.Switch)
		}

		if stmt.Task != nil {
			err = r.handleTask(res, stmt.Task)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Builder) handleWorkspace(stmt *parser.Workspace) error {
	if _, ok := r.recipe.Env[stmt.Name]; ok {
		return fmt.Errorf("environment redefined: %s", stmt.Name)
	}

	res := NewEnv()
	res.Name = stmt.Name
	r.recipe.Workspace = res

	for _, stmt := range stmt.Stmts {
		var err error

		if stmt.Assign != nil {
			err = r.handleAssign(res, stmt.Assign)
		}

		if stmt.Decl != nil {
			err = r.handleDecl(res, stmt.Decl)
		}

		if stmt.Inherit != nil {
			err = r.handleInherit(res, stmt.Inherit)
		}

		if stmt.Project != nil {
			err = r.handleProject(res, stmt.Project)
		}

		if stmt.Rule != nil {
			err = r.handleRule(res, stmt.Rule)
		}

		if stmt.Switch != nil {
			err = r.handleSwitch(res, stmt.Switch)
		}

		if stmt.Task != nil {
			err = r.handleTask(res, stmt.Task)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Builder) handleAssign(env *Environment, stmt *parser.Assign) error {
	return r.doAssign(env, stmt.Ident, stmt.Value)
}

func (r *Builder) doAssign(env *Environment, ident string, value *parser.Value) error {
	if attr, ok := env.Attributes[ident]; ok {
		if value.List != nil {
			if list, ok := attr.(*ListAttribute); ok {
				for _, item := range value.List {
					if item.String != nil {
						list.Append(*item.String)
					} else {
						return fmt.Errorf("illegal list item type assigned to identifier: %s", ident)
					}
				}
				return nil
			}
			return fmt.Errorf("illegal type assigned to identifier: %s", ident)
		}
		if value.String != nil {
			if str, ok := attr.(*StringAttribute); ok {
				str.Assign(*value.String)
				return nil
			}
			return fmt.Errorf("illegal type assigned to identifier: %s", ident)
		}
		if value.Bool != nil {
			if b, ok := attr.(*BoolAttribute); ok {
				b.Assign(*value.Bool)
				return nil
			}
			return fmt.Errorf("illegal type assigned to identifier: %s", ident)
		}
		return nil
	}

	return fmt.Errorf("undefined symbol: %s", ident)
}

func (r *Builder) handleDecl(env *Environment, stmt *parser.Decl) error {
	if _, ok := env.Attributes[stmt.Ident]; ok {
		return fmt.Errorf("symbol redefined: %s", stmt.Ident)
	}

	switch stmt.Type {
	case "list":
		env.Attributes[stmt.Ident] = NewListAttribute()
	case "string":
		env.Attributes[stmt.Ident] = NewStringAttribute("")
	default:
		return fmt.Errorf("illegal type: %s", stmt.Type)
	}

	if stmt.Value != nil {
		return r.doAssign(env, stmt.Ident, stmt.Value)
	}

	return nil
}

func (r *Builder) handleInherit(env *Environment, stmt *parser.Inherit) error {
	if inherited, ok := r.recipe.Env[stmt.Ident]; ok {
		env.Inherit(inherited)
		return nil
	}

	return fmt.Errorf("environment undefined: %s", stmt.Ident)
}

func (r *Builder) handleProject(env *Environment, stmt *parser.Project) error {
	if _, ok := r.recipe.Env[stmt.Name]; ok {
		return fmt.Errorf("environment redefined: %s", stmt.Name)
	}

	res := NewEnv()
	res.Name = stmt.Name
	r.recipe.Env[stmt.Name] = res
	res.Inherit(env)

	for _, stmt := range stmt.Block {
		var err error

		if stmt.Assign != nil {
			err = r.handleAssign(res, stmt.Assign)
		}

		if stmt.Decl != nil {
			err = r.handleDecl(res, stmt.Decl)
		}

		if stmt.Inherit != nil {
			err = r.handleInherit(res, stmt.Inherit)
		}

		if stmt.Rule != nil {
			err = r.handleRule(res, stmt.Rule)
		}

		if stmt.Switch != nil {
			err = r.handleSwitch(res, stmt.Switch)
		}

		if stmt.Task != nil {
			err = r.handleTask(res, stmt.Task)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Builder) handleRule(env *Environment, stmt *parser.Rule) error {
	renv := NewEnv()
	renv.Attributes["aggregate"] = NewBoolAttribute(false)
	renv.Attributes["ext"] = NewListAttribute()
	renv.Attributes["outputs"] = NewListAttribute()
	renv.Attributes["depfile"] = NewStringAttribute("")
	renv.Attributes["mandatory"] = NewBoolAttribute(false)
	renv.Attributes["message"] = NewStringAttribute("")
	renv.Attributes["phony"] = NewBoolAttribute(false)

	switch stmt.Type {
	case "skip":
		renv.reset()
		renv.Attributes["ext"] = NewListAttribute()
	case "template":
		renv.Attributes["template"] = NewStringAttribute("")
	default:
		renv.Attributes["command"] = NewListAttribute()
	}

	for _, assign := range stmt.Assignments {
		err := r.handleAssign(renv, assign)
		if err != nil {
			return err
		}
	}

	var rule Rule

	switch stmt.Type {
	case "skip":
		rule = &SkipRule{}

	case "template":
		rule = &TemplateRule{
			BasicRule: BasicRule{
				Aggregate: renv.Attributes["aggregate"].(*BoolAttribute).IsTrue(),
				Mandatory: renv.Attributes["mandatory"].(*BoolAttribute).IsTrue(),
				Message:   renv.Attributes["message"].String(),
				Phony:     renv.Attributes["phony"].(*BoolAttribute).IsTrue(),
			},
			Output:   renv.Attributes["outputs"].(*ListAttribute).Items(),
			Template: renv.Attributes["template"].String(),
		}

	default:
		rule = &CommandRule{
			BasicRule: BasicRule{
				Aggregate: renv.Attributes["aggregate"].(*BoolAttribute).IsTrue(),
				Mandatory: renv.Attributes["mandatory"].(*BoolAttribute).IsTrue(),
				Message:   renv.Attributes["message"].String(),
				Phony:     renv.Attributes["phony"].(*BoolAttribute).IsTrue(),
			},
			Command: renv.Attributes["command"].(*ListAttribute).Join("&&").String(),
			Output:  renv.Attributes["outputs"].(*ListAttribute).Items(),
		}
	}

	env.Rules[stmt.Ident] = rule
	for _, ext := range renv.Attributes["ext"].(*ListAttribute).Items() {
		env.RulesByExt[ext] = append(env.RulesByExt[ext], rule)
	}

	return nil
}

func (r *Builder) handleSwitch(env *Environment, stmt *parser.Switch) error {
	return nil
}

func (r *Builder) handleTask(env *Environment, stmt *parser.Task) error {
	if _, ok := r.recipe.Env[stmt.Name]; ok {
		return fmt.Errorf("environment redefined: %s", stmt.Name)
	}

	res := NewTask(stmt.Name)
	res.Inherit(env)
	r.recipe.Env[stmt.Name] = &res.Environment
	env.Tasks[stmt.Name] = res

	for _, stmt := range stmt.Block {
		var err error

		if stmt.Assign != nil {
			err = r.handleAssign(&res.Environment, stmt.Assign)
		}

		if stmt.Decl != nil {
			err = r.handleDecl(&res.Environment, stmt.Decl)
		}

		if stmt.Inherit != nil {
			err = r.handleInherit(&res.Environment, stmt.Inherit)
		}

		if stmt.Rule != nil {
			err = r.handleRule(&res.Environment, stmt.Rule)
		}

		if stmt.Switch != nil {
			err = r.handleSwitch(&res.Environment, stmt.Switch)
		}

		if stmt.TaskSteps != nil {
			err = r.handleTaskSteps(res, stmt.TaskSteps)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Builder) handleTaskSteps(task *Task, stmt *parser.TaskSteps) error {
	for _, stmt := range stmt.Block {
		var err error

		if stmt.Transform != nil {
			err = r.handleStepTransform(task, stmt.Transform)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Builder) handleStepTransform(task *Task, stmt *parser.Transform) error {
	inputs := []string{}

	for _, ident := range stmt.Ident {
		var attr Attribute
		var list *ListAttribute
		var ok bool

		if attr, ok = task.Environment.Attributes[ident]; !ok {
			return fmt.Errorf("undefined symbol: %s", ident)
		}

		if list, ok = attr.(*ListAttribute); !ok {
			return fmt.Errorf("unexpected type: %s", ident)
		}

		inputs = append(inputs, list.Items()...)
	}

	job := NewRuleGeneratorJob(&task.Environment, inputs...)
	task.Jobs = append(task.Jobs, job)

	return nil
}
