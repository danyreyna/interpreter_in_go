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
		expectedKind    string
		expectedLiteral string
	}{
		{assign, ""},
		{plus, ""},
		{leftParenthesis, ""},
		{rightParenthesis, ""},
		{leftCurlyBrace, ""},
		{rightCurlyBrace, ""},
		{comma, ""},
		{semicolon, ""},
		{let, ""},
		{identifier, "five"},
		{assign, ""},
		{integer, "5"},
		{semicolon, ""},
		{let, ""},
		{identifier, "ten"},
		{assign, ""},
		{integer, "10"},
		{semicolon, ""},
		{let, ""},
		{identifier, "add"},
		{assign, ""},
		{fn, ""},
		{leftParenthesis, ""},
		{identifier, "x"},
		{comma, ""},
		{identifier, "y"},
		{rightParenthesis, ""},
		{leftCurlyBrace, ""},
		{identifier, "x"},
		{plus, ""},
		{identifier, "y"},
		{semicolon, ""},
		{rightCurlyBrace, ""},
		{semicolon, ""},
		{let, ""},
		{identifier, "result"},
		{assign, ""},
		{identifier, "add"},
		{leftParenthesis, ""},
		{identifier, "five"},
		{comma, ""},
		{identifier, "ten"},
		{rightParenthesis, ""},
		{semicolon, ""},
		{eof, ""},
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
	}
}
