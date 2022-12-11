package parser

type Switch struct {
	Ident   string             `parser:"'switch' @Ident '{'"`
	Cases   []*SwitchCase      `parser:"@@*"`
	Default *DefaultSwitchCase `parser:"( @@ ) '}'"`
}

type SwitchCase struct {
	Value string             `parser:"'case' @String"`
	Block []*EnvironmentStmt `parser:"'{' @@* '}'"`
}

type DefaultSwitchCase struct {
	Value string             `parser:"'default'"`
	Block []*EnvironmentStmt `parser:"'{' @@* '}'"`
}
