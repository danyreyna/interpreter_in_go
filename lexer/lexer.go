package lexer

import (
	"bufio"
	"os"
)

type Lexer struct {
	file                *os.File
	fileReader          *bufio.Reader
	filePath            string
	currentCodePoint    rune
	currentLineNumber   int
	currentColumnNumber int
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func newLexer(filePath string) *Lexer {
	file, err := os.Open(filePath)
	check(err)

	fileReader := bufio.NewReader(file)

	var currentCodePoint rune
	if codePoint, _, err := fileReader.ReadRune(); err == nil {
		currentCodePoint = codePoint
	} else {
		currentCodePoint = 0
	}

	lexerInstance := &Lexer{
		file,
		fileReader,
		filePath,
		currentCodePoint,
		1,
		1,
	}

	return lexerInstance
}

func (lexerInstance *Lexer) updateCurrentCodePoint() {
	if codePoint, _, err := lexerInstance.fileReader.ReadRune(); err == nil {
		lexerInstance.currentCodePoint = codePoint
		return
	}

	lexerInstance.currentCodePoint = 0
}

func (lexerInstance *Lexer) peekAtNextCodePoint() rune {
	codePoint, _, readErr := lexerInstance.fileReader.ReadRune()
	check(readErr)

	unreadErr := lexerInstance.fileReader.UnreadRune()
	check(unreadErr)

	return codePoint
}

func (lexerInstance *Lexer) skipWhitespace() {
	for lexerInstance.currentCodePoint == ' ' ||
		lexerInstance.currentCodePoint == '\t' ||
		lexerInstance.currentCodePoint == '\n' ||
		lexerInstance.currentCodePoint == '\r' {
		if lexerInstance.currentCodePoint == ' ' {
			lexerInstance.currentColumnNumber += 1
		}
		if lexerInstance.currentCodePoint == '\t' {
			lexerInstance.currentColumnNumber += 4 - (lexerInstance.currentColumnNumber % 4)
		}
		if lexerInstance.currentCodePoint == '\n' {
			lexerInstance.currentLineNumber += 1
			lexerInstance.currentColumnNumber = 1
		}

		lexerInstance.updateCurrentCodePoint()
	}
}

func (lexerInstance *Lexer) getMultiCodePointToken(isAllowedCodePoint func(rune) bool) string {
	var tokenSlice []rune

	for isAllowedCodePoint(lexerInstance.currentCodePoint) {
		tokenSlice = append(tokenSlice, lexerInstance.currentCodePoint)
		lexerInstance.updateCurrentCodePoint()
	}

	return string(tokenSlice)
}

func isNumeric(codePoint rune) bool {
	return '0' <= codePoint && codePoint <= '9'
}

func (lexerInstance *Lexer) getInteger() string {
	return lexerInstance.getMultiCodePointToken(isNumeric)
}

func isPartOfWord(codePoint rune) bool {
	excludedCodePoints := map[rune]struct{}{
		' ':  {},
		'\t': {},
		'\n': {},
		'\r': {},
		'=':  {},
		'*':  {},
		'!':  {},
		',':  {},
		'>':  {},
		'{':  {},
		'(':  {},
		'<':  {},
		'-':  {},
		'+':  {},
		'}':  {},
		')':  {},
		';':  {},
		'/':  {},
		0:    {},
	}

	_, isExcludedCodePoint := excludedCodePoints[codePoint]
	return !isExcludedCodePoint
}

func (lexerInstance *Lexer) getWord() string {
	return lexerInstance.getMultiCodePointToken(isPartOfWord)
}

func isStartOfWord(codePoint rune) bool {
	return isPartOfWord(codePoint) && !isNumeric(codePoint)
}

func (lexerInstance *Lexer) handleWordToken(lineNumber int, columnNumber int) token {
	word := lexerInstance.getWord()

	lexerInstance.currentColumnNumber += len([]rune(word))

	if keywordValue, isKeyword := keywords[word]; isKeyword {
		return token{
			kind:         keywordValue,
			filePath:     lexerInstance.filePath,
			lineNumber:   lineNumber,
			columnNumber: columnNumber,
		}
	}

	return token{
		identifier,
		word,
		lexerInstance.filePath,
		lineNumber,
		columnNumber,
	}
}

func (lexerInstance *Lexer) handleNumberToken(lineNumber int, columnNumber int) token {
	number := lexerInstance.getInteger()

	lexerInstance.currentColumnNumber += len(number)

	return token{
		integer,
		number,
		lexerInstance.filePath,
		lineNumber,
		columnNumber,
	}
}

func (lexerInstance *Lexer) handleSingleCodePoint(codePoint rune, lineNumber int, columnNumber int) token {
	if codePoint != 0 {
		lexerInstance.currentColumnNumber += 1

		lexerInstance.updateCurrentCodePoint()
	}

	switch codePoint {
	case '*':
		return token{kind: asterisk, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case ',':
		return token{kind: comma, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case '>':
		return token{kind: greaterThan, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case '{':
		return token{
			kind: leftCurlyBrace, filePath: lexerInstance.filePath,
			lineNumber: lineNumber, columnNumber: columnNumber,
		}
	case '(':
		return token{
			kind: leftParenthesis, filePath: lexerInstance.filePath,
			lineNumber: lineNumber, columnNumber: columnNumber,
		}
	case '<':
		return token{kind: lessThan, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case '-':
		return token{kind: minus, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case '+':
		return token{kind: plus, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case '}':
		return token{
			kind: rightCurlyBrace, filePath: lexerInstance.filePath,
			lineNumber: lineNumber, columnNumber: columnNumber,
		}
	case ')':
		return token{
			kind: rightParenthesis, filePath: lexerInstance.filePath,
			lineNumber: lineNumber, columnNumber: columnNumber,
		}
	case ';':
		return token{kind: semicolon, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case '/':
		return token{kind: slash, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}

	case 0:
		err := lexerInstance.file.Close()
		check(err)
		return token{kind: eof, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	default:
		return token{unknown, string(codePoint), lexerInstance.filePath, lineNumber, columnNumber}
	}
}

func (lexerInstance *Lexer) handlePotentialDoubleCodePointToken(
	expectedNextCodePoint rune,
	singleTokenKind byte,
	doubleTokenKind byte,
	lineNumber int,
	columnNumber int,
) token {
	if lexerInstance.peekAtNextCodePoint() == expectedNextCodePoint {
		lexerInstance.currentColumnNumber += 2

		lexerInstance.updateCurrentCodePoint()
		lexerInstance.updateCurrentCodePoint()

		return token{
			kind: doubleTokenKind, filePath: lexerInstance.filePath,
			lineNumber: lineNumber, columnNumber: columnNumber,
		}
	}

	lexerInstance.currentColumnNumber += 1

	lexerInstance.updateCurrentCodePoint()

	return token{
		kind: singleTokenKind, filePath: lexerInstance.filePath,
		lineNumber: lineNumber, columnNumber: columnNumber,
	}
}

func (lexerInstance *Lexer) handleExclamationMarkCodePoint(lineNumber int, columnNumber int) token {
	return lexerInstance.handlePotentialDoubleCodePointToken('=', bang, inequality, lineNumber, columnNumber)
}

func (lexerInstance *Lexer) handleEqualsSignCodePoint(lineNumber int, columnNumber int) token {
	return lexerInstance.handlePotentialDoubleCodePointToken('=', assign, equality, lineNumber, columnNumber)
}

func (lexerInstance *Lexer) getNextToken() token {
	lexerInstance.skipWhitespace()

	codePoint := lexerInstance.currentCodePoint
	lineNumber := lexerInstance.currentLineNumber
	columnNumber := lexerInstance.currentColumnNumber

	if isStartOfWord(codePoint) {
		return lexerInstance.handleWordToken(lineNumber, columnNumber)
	}

	if isNumeric(codePoint) {
		return lexerInstance.handleNumberToken(lineNumber, columnNumber)
	}

	if codePoint == '!' {
		return lexerInstance.handleExclamationMarkCodePoint(lineNumber, columnNumber)
	}

	if codePoint == '=' {
		return lexerInstance.handleEqualsSignCodePoint(lineNumber, columnNumber)
	}

	return lexerInstance.handleSingleCodePoint(codePoint, lineNumber, columnNumber)
}
