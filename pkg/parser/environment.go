package parser

type Environment struct {
	Name  string             `parser:"( @Ident | @String )"`
	Stmts []*EnvironmentStmt `parser:"'{' @@* '}'"`
}

type EnvironmentStmt struct {
	Assign  *Assign  `parser:"  @@"`
	Decl    *Decl    `parser:"| @@"`
	Inherit *Inherit `parser:"| @@"`
	Project *Project `parser:"| @@"`
	Rule    *Rule    `parser:"| @@"`
	Switch  *Switch  `parser:"| @@"`
	Task    *Task    `parser:"| @@"`
}
