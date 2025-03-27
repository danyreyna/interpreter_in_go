package lexer

import "testing"

func TestGetNextToken(t *testing.T) {
	codeString := `=+(){},;
let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
};

let result = add(five, ten);
`

	tests := []struct {
		expectedKind         string
		expectedLiteral      string
		expectedLineNumber   int
		expectedColumnNumber int
	}{
		{assign, "", 1, 1},
		{plus, "", 1, 2},
		{leftParenthesis, "", 1, 3},
		{rightParenthesis, "", 1, 4},
		{leftCurlyBrace, "", 1, 5},
		{rightCurlyBrace, "", 1, 6},
		{comma, "", 1, 7},
		{semicolon, "", 1, 8},
		{let, "", 2, 1},
		{identifier, "five", 2, 5},
		{assign, "", 2, 10},
		{integer, "5", 2, 12},
		{semicolon, "", 2, 13},
		{let, "", 3, 1},
		{identifier, "ten", 3, 5},
		{assign, "", 3, 9},
		{integer, "10", 3, 11},
		{semicolon, "", 3, 13},
		{let, "", 5, 1},
		{identifier, "add", 5, 5},
		{assign, "", 5, 9},
		{fn, "", 5, 11},
		{leftParenthesis, "", 5, 13},
		{identifier, "x", 5, 14},
		{comma, "", 5, 15},
		{identifier, "y", 5, 17},
		{rightParenthesis, "", 5, 18},
		{leftCurlyBrace, "", 5, 20},
		{identifier, "x", 6, 3},
		{plus, "", 6, 5},
		{identifier, "y", 6, 7},
		{semicolon, "", 6, 8},
		{rightCurlyBrace, "", 7, 1},
		{semicolon, "", 7, 2},
		{let, "", 9, 1},
		{identifier, "result", 9, 5},
		{assign, "", 9, 12},
		{identifier, "add", 9, 14},
		{leftParenthesis, "", 9, 17},
		{identifier, "five", 9, 18},
		{comma, "", 9, 22},
		{identifier, "ten", 9, 24},
		{rightParenthesis, "", 9, 27},
		{semicolon, "", 9, 28},
		{eof, "", 10, 1},
	}

	lexerInstance := newLexer(codeString)

	for i, currentTest := range tests {
		currentToken := lexerInstance.getNextToken()

		if currentToken.kind != currentTest.expectedKind {
			t.Fatalf(
				"tests[%d] — kind is wrong. expected=%q, got=%q",
				i,
				currentTest.expectedKind,
				currentToken.kind,
			)
		}

		if currentToken.literal != currentTest.expectedLiteral {
			t.Fatalf(
				"tests[%d] — literal is wrong. expected=%q, got=%q",
				i,
				currentTest.expectedLiteral,
				currentToken.literal,
			)
		}

		if currentToken.lineNumber != currentTest.expectedLineNumber {
			t.Fatalf(
				"tests[%d] — line number is wrong. expected=%d, got=%d",
				i,
				currentTest.expectedLineNumber,
				currentToken.lineNumber,
			)
		}

		if currentToken.columnNumber != currentTest.expectedColumnNumber {
			t.Fatalf(
				"tests[%d] — column number is wrong. expected=%d, got=%d",
				i,
				currentTest.expectedColumnNumber,
				currentToken.columnNumber,
			)
		}
	}
}
