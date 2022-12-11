package parser

import (
	"io"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	joltlexer = lexer.MustSimple([]lexer.SimpleRule{
		{"comment", `//.*|/\*.*?\*/`},
		{"whitespace", `\s+`},

		{"Type", `\b(list|string)\b`},
		{`String`, `"(?:\\.|[^"])*"`},
		{"Ident", `\b([a-zA-Z_][a-zA-Z0-9_]*)\b`},
		{"Punct", `[-,()*/+%{};&!=:<>]|\[|\]`},
	})

	parser = participle.MustBuild[File](
		participle.Lexer(joltlexer),
		participle.Unquote("String"),
	)
)

func Parse(reader io.Reader) (*File, error) {
	return parser.Parse("", reader)
}
