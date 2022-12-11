package parser

type Workspace struct {
	Name  string           `parser:"( @Ident | @String )"`
	Stmts []*WorkspaceStmt `parser:"'{' @@* '}'"`
}

type WorkspaceStmt struct {
	Assign  *Assign  `parser:"  @@"`
	Decl    *Decl    `parser:"| @@"`
	Inherit *Inherit `parser:"| @@"`
	Project *Project `parser:"| @@"`
	Rule    *Rule    `parser:"| @@"`
	Switch  *Switch  `parser:"| @@"`
	Task    *Task    `parser:"| @@"`
}
