package parser

import "github.com/alecthomas/participle/lexer"

type File struct {
	Pos lexer.Position

	FileStmt []*FileStmt `parser:"@@*"`
}

type FileStmt struct {
	Pos lexer.Position

	Environment *Environment `parser:"  'environment' @@"`
	Workspace   *Workspace   `parser:"| 'workspace' @@"`
}
