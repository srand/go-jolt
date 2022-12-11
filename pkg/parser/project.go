package parser

type Project struct {
	Name  string             `parser:"'project' @String"`
	Block []*EnvironmentStmt `parser:"'{' @@* '}'"`
}
