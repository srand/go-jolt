package parser

type Decl struct {
	Type  string `parser:"@Type"`
	Ident string `parser:"@Ident"`
	Value *Value `parser:"['=' @@ ] [';']"`
}
