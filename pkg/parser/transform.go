package parser

type Transform struct {
	Ident []string `parser:"'transform' @Ident (',' @Ident)* ';'"`
}
