package lexer

import "testing"

func TestGetNextToken(t *testing.T) {
	tests := []struct {
		expectedKind         byte
		expectedLiteral      string
		expectedFilePath     string
		expectedLineNumber   int
		expectedColumnNumber int
	}{
		{assign, "", "./test-file.code", 1, 1},
		{plus, "", "./test-file.code", 1, 2},
		{leftParenthesis, "", "./test-file.code", 1, 3},
		{rightParenthesis, "", "./test-file.code", 1, 4},
		{leftCurlyBrace, "", "./test-file.code", 1, 5},
		{rightCurlyBrace, "", "./test-file.code", 1, 6},
		{comma, "", "./test-file.code", 1, 7},
		{semicolon, "", "./test-file.code", 1, 8},
		{let, "", "./test-file.code", 2, 1},
		{identifier, "five", "./test-file.code", 2, 5},
		{assign, "", "./test-file.code", 2, 10},
		{integer, "5", "./test-file.code", 2, 12},
		{semicolon, "", "./test-file.code", 2, 13},
		{let, "", "./test-file.code", 3, 1},
		{identifier, "ten", "./test-file.code", 3, 5},
		{assign, "", "./test-file.code", 3, 9},
		{integer, "10", "./test-file.code", 3, 11},
		{semicolon, "", "./test-file.code", 3, 13},
		{let, "", "./test-file.code", 5, 1},
		{identifier, "add", "./test-file.code", 5, 5},
		{assign, "", "./test-file.code", 5, 9},
		{fn, "", "./test-file.code", 5, 11},
		{leftParenthesis, "", "./test-file.code", 5, 13},
		{identifier, "x", "./test-file.code", 5, 14},
		{comma, "", "./test-file.code", 5, 15},
		{identifier, "y", "./test-file.code", 5, 17},
		{rightParenthesis, "", "./test-file.code", 5, 18},
		{leftCurlyBrace, "", "./test-file.code", 5, 20},
		{identifier, "x", "./test-file.code", 6, 3},
		{plus, "", "./test-file.code", 6, 5},
		{identifier, "y", "./test-file.code", 6, 7},
		{semicolon, "", "./test-file.code", 6, 8},
		{rightCurlyBrace, "", "./test-file.code", 7, 1},
		{semicolon, "", "./test-file.code", 7, 2},
		{let, "", "./test-file.code", 9, 1},
		{identifier, "result", "./test-file.code", 9, 5},
		{assign, "", "./test-file.code", 9, 12},
		{identifier, "add", "./test-file.code", 9, 14},
		{leftParenthesis, "", "./test-file.code", 9, 17},
		{identifier, "five", "./test-file.code", 9, 18},
		{comma, "", "./test-file.code", 9, 22},
		{identifier, "ten", "./test-file.code", 9, 24},
		{rightParenthesis, "", "./test-file.code", 9, 27},
		{semicolon, "", "./test-file.code", 9, 28},
		{let, "", "./test-file.code", 10, 1},
		{identifier, "犬の数", "./test-file.code", 10, 5},
		{assign, "", "./test-file.code", 10, 9},
		{integer, "5", "./test-file.code", 10, 11},
		{semicolon, "", "./test-file.code", 10, 12},
		{let, "", "./test-file.code", 11, 1},
		{identifier, "猫的数量", "./test-file.code", 11, 5},
		{assign, "", "./test-file.code", 11, 10},
		{integer, "2", "./test-file.code", 11, 12},
		{semicolon, "", "./test-file.code", 11, 13},
		{eof, "", "./test-file.code", 12, 1},
	}

	lexerInstance := newLexer("./test-file.code")

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

		if currentToken.filePath != currentTest.expectedFilePath {
			t.Fatalf(
				"tests[%d] — file path is wrong. expected=%q, got=%q",
				i,
				currentTest.expectedFilePath,
				currentToken.filePath,
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
