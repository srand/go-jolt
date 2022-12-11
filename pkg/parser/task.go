package parser

type Task struct {
	Name  string      `parser:"'task' @String"`
	Block []*TaskStmt `parser:"'{' @@* '}'"`
}

type TaskStmt struct {
	Assign    *Assign    `parser:"  @@"`
	Decl      *Decl      `parser:"| @@"`
	Inherit   *Inherit   `parser:"| @@"`
	Rule      *Rule      `parser:"| @@"`
	Switch    *Switch    `parser:"| @@"`
	TaskSteps *TaskSteps `parser:"| @@"`
}
