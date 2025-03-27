package lexer

const (
	elseKeyword   = byte(0)
	falseKeyword  = byte(1)
	fn            = byte(2)
	ifKeyword     = byte(3)
	let           = byte(4)
	returnKeyword = byte(5)
	trueKeyword   = byte(6)

	identifier = byte(7)

	integer = byte(8)

	assign           = byte(9)
	asterisk         = byte(10)
	bang             = byte(11)
	comma            = byte(12)
	equality         = byte(13)
	greaterThan      = byte(14)
	inequality       = byte(15)
	leftCurlyBrace   = byte(16)
	leftParenthesis  = byte(17)
	lessThan         = byte(18)
	minus            = byte(19)
	plus             = byte(20)
	rightCurlyBrace  = byte(21)
	rightParenthesis = byte(22)
	semicolon        = byte(23)
	slash            = byte(24)

	eof     = byte(25)
	unknown = byte(26)
)

type token struct {
	kind         byte
	literal      string
	filePath     string
	lineNumber   int
	columnNumber int
}

var keywords = map[string]byte{
	"else":   elseKeyword,
	"false":  falseKeyword,
	"fn":     fn,
	"if":     ifKeyword,
	"let":    let,
	"return": returnKeyword,
	"true":   trueKeyword,
}
