package parser

type Value struct {
	String   *string  `parser:"  @String"`
	Bool     *string  `parser:"| ( @'true' | 'false' )"`
	GlobList []*Value `parser:"| 'glob' '{' @@ ( ',' @@ )* [','] '}' "`
	List     []*Value `parser:"| '{' @@ ( ',' @@ )* [','] '}' "`
}
