package parser

type Assign struct {
	Ident string `parser:"@Ident '='"`
	Value *Value `parser:"@@ [';']"`
}
