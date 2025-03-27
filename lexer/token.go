package lexer

const (
	fn  = "fn"
	let = "let"

	identifier = "identifier"

	integer = "int"

	assign           = "="
	comma            = ","
	leftCurlyBrace   = "{"
	leftParenthesis  = "("
	plus             = "+"
	rightCurlyBrace  = "}"
	rightParenthesis = ")"
	semicolon        = ";"

	eof     = "eof"
	unknown = "unknown"
)

type token struct {
	kind         string
	literal      string
	filePath     string
	lineNumber   int
	columnNumber int
}

var keywords = map[string]struct{}{
	fn:  {},
	let: {},
}
