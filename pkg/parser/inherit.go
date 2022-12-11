package parser

type Inherit struct {
	Ident string `parser:"'inherit' ( @Ident | @String ) ';'"`
}
