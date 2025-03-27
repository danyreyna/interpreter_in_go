package lexer

const (
	fn  = byte(0)
	let = byte(1)

	identifier = byte(2)

	integer = byte(3)

	assign           = byte(4)
	comma            = byte(5)
	leftCurlyBrace   = byte(6)
	leftParenthesis  = byte(7)
	plus             = byte(8)
	rightCurlyBrace  = byte(9)
	rightParenthesis = byte(10)
	semicolon        = byte(11)

	eof     = byte(12)
	unknown = byte(13)
)

type token struct {
	kind         byte
	literal      string
	filePath     string
	lineNumber   int
	columnNumber int
}

var keywords = map[string]byte{
	"fn":  fn,
	"let": let,
}
