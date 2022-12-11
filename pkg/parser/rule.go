package parser

type Rule struct {
	Ident       string    `parser:"'rule' @Ident"`
	Type        string    `parser:"[ ':' @Ident ]"`
	Assignments []*Assign `parser:"'{' @@* '}'"`
}
