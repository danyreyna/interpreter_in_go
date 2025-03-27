package lexer

import (
	"interpreter_in_go/common"
	"os"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	filePath := "./test-file.code"

	tests := []struct {
		expectedKind         byte
		expectedLiteral      string
		expectedFilePath     string
		expectedLineNumber   int
		expectedColumnNumber int
	}{
		{assign, "", filePath, 1, 1},
		{plus, "", filePath, 1, 2},
		{leftParenthesis, "", filePath, 1, 3},
		{rightParenthesis, "", filePath, 1, 4},
		{leftCurlyBrace, "", filePath, 1, 5},
		{rightCurlyBrace, "", filePath, 1, 6},
		{comma, "", filePath, 1, 7},
		{semicolon, "", filePath, 1, 8},
		{let, "", filePath, 2, 1},
		{identifier, "five", filePath, 2, 5},
		{assign, "", filePath, 2, 10},
		{integer, "5", filePath, 2, 12},
		{semicolon, "", filePath, 2, 13},
		{let, "", filePath, 3, 1},
		{identifier, "ten", filePath, 3, 5},
		{assign, "", filePath, 3, 9},
		{integer, "10", filePath, 3, 11},
		{semicolon, "", filePath, 3, 13},
		{let, "", filePath, 5, 1},
		{identifier, "add", filePath, 5, 5},
		{assign, "", filePath, 5, 9},
		{fn, "", filePath, 5, 11},
		{leftParenthesis, "", filePath, 5, 13},
		{identifier, "x", filePath, 5, 14},
		{comma, "", filePath, 5, 15},
		{identifier, "y", filePath, 5, 17},
		{rightParenthesis, "", filePath, 5, 18},
		{leftCurlyBrace, "", filePath, 5, 20},
		{identifier, "x", filePath, 6, 3},
		{plus, "", filePath, 6, 5},
		{identifier, "y", filePath, 6, 7},
		{semicolon, "", filePath, 6, 8},
		{rightCurlyBrace, "", filePath, 7, 1},
		{semicolon, "", filePath, 7, 2},
		{let, "", filePath, 9, 1},
		{identifier, "result", filePath, 9, 5},
		{assign, "", filePath, 9, 12},
		{identifier, "add", filePath, 9, 14},
		{leftParenthesis, "", filePath, 9, 17},
		{identifier, "five", filePath, 9, 18},
		{comma, "", filePath, 9, 22},
		{identifier, "ten", filePath, 9, 24},
		{rightParenthesis, "", filePath, 9, 27},
		{semicolon, "", filePath, 9, 28},
		{let, "", filePath, 10, 1},
		{identifier, "犬の数", filePath, 10, 5},
		{assign, "", filePath, 10, 9},
		{integer, "5", filePath, 10, 11},
		{semicolon, "", filePath, 10, 12},
		{let, "", filePath, 11, 1},
		{identifier, "猫的数量", filePath, 11, 5},
		{assign, "", filePath, 11, 10},
		{integer, "2", filePath, 11, 12},
		{semicolon, "", filePath, 11, 13},
		{bang, "", filePath, 13, 1},
		{minus, "", filePath, 13, 2},
		{slash, "", filePath, 13, 3},
		{asterisk, "", filePath, 13, 4},
		{integer, "5", filePath, 13, 5},
		{semicolon, "", filePath, 13, 6},
		{integer, "5", filePath, 14, 1},
		{lessThan, "", filePath, 14, 3},
		{integer, "10", filePath, 14, 5},
		{greaterThan, "", filePath, 14, 8},
		{integer, "5", filePath, 14, 10},
		{semicolon, "", filePath, 14, 11},
		{ifKeyword, "", filePath, 16, 1},
		{leftParenthesis, "", filePath, 16, 4},
		{integer, "5", filePath, 16, 5},
		{lessThan, "", filePath, 16, 7},
		{integer, "10", filePath, 16, 9},
		{rightParenthesis, "", filePath, 16, 11},
		{leftCurlyBrace, "", filePath, 16, 13},
		{returnKeyword, "", filePath, 17, 5},
		{trueKeyword, "", filePath, 17, 12},
		{semicolon, "", filePath, 17, 16},
		{rightCurlyBrace, "", filePath, 18, 1},
		{elseKeyword, "", filePath, 18, 3},
		{leftCurlyBrace, "", filePath, 18, 8},
		{returnKeyword, "", filePath, 19, 5},
		{falseKeyword, "", filePath, 19, 12},
		{semicolon, "", filePath, 19, 17},
		{rightCurlyBrace, "", filePath, 20, 1},
		{integer, "10", filePath, 22, 1},
		{equality, "", filePath, 22, 4},
		{integer, "10", filePath, 22, 7},
		{semicolon, "", filePath, 22, 9},
		{integer, "10", filePath, 23, 1},
		{inequality, "", filePath, 23, 4},
		{integer, "9", filePath, 23, 7},
		{semicolon, "", filePath, 23, 8},
		{eof, "", filePath, 24, 1},
	}

	file, err := os.Open(filePath)
	common.Check(err)
	defer common.CloseFile(file)

	lexerInstance := newLexer(file, filePath)

	for i, currentTest := range tests {
		currentToken := lexerInstance.getNextToken()

		if currentToken.kind != currentTest.expectedKind {
			t.Fatalf(
				"tests[%d] — kind is wrong. expected=%d, got=%d",
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
