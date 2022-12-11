package parser

type TaskSteps struct {
	Block []*TaskStepStmt `parser:"'steps' '{' @@* '}'"`
}

type TaskStepStmt struct {
	Transform *Transform `parser:"  @@"`
}
