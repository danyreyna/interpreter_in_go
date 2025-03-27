package lexer

import (
	"bufio"
	"os"
)

type Lexer struct {
	file                *os.File
	fileReader          *bufio.Reader
	filePath            string
	currentCharacter    byte
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

	var currentCharacter byte
	if character, err := fileReader.ReadByte(); err == nil {
		currentCharacter = character
	} else {
		currentCharacter = 0
	}

	lexerInstance := &Lexer{
		file,
		fileReader,
		filePath,
		currentCharacter,
		1,
		1,
	}

	return lexerInstance
}

func (lexerInstance *Lexer) updateCurrentCharacter() {
	if character, err := lexerInstance.fileReader.ReadByte(); err == nil {
		lexerInstance.currentCharacter = character
		return
	}

	lexerInstance.currentCharacter = 0
}

func (lexerInstance *Lexer) skipWhitespace() {
	for lexerInstance.currentCharacter == ' ' ||
		lexerInstance.currentCharacter == '\t' ||
		lexerInstance.currentCharacter == '\n' ||
		lexerInstance.currentCharacter == '\r' {
		if lexerInstance.currentCharacter == ' ' {
			lexerInstance.currentColumnNumber += 1
		}
		if lexerInstance.currentCharacter == '\t' {
			lexerInstance.currentColumnNumber += 4 - (lexerInstance.currentColumnNumber % 4)
		}
		if lexerInstance.currentCharacter == '\n' {
			lexerInstance.currentLineNumber += 1
			lexerInstance.currentColumnNumber = 1
		}

		lexerInstance.updateCurrentCharacter()
	}
}

func (lexerInstance *Lexer) getMultiCharacterToken(isAllowedCharacter func(byte) bool) string {
	var tokenSlice []byte

	for isAllowedCharacter(lexerInstance.currentCharacter) {
		tokenSlice = append(tokenSlice, lexerInstance.currentCharacter)
		lexerInstance.updateCurrentCharacter()
	}

	return string(tokenSlice)
}

func isLowerCase(character byte) bool {
	return 'a' <= character && character <= 'z'
}

func isUpperCase(character byte) bool {
	return 'A' <= character && character <= 'Z'
}

func isAlphabetical(character byte) bool {
	return isLowerCase(character) || isUpperCase(character) || character == '_'
}

func (lexerInstance *Lexer) getWord() string {
	return lexerInstance.getMultiCharacterToken(isAlphabetical)
}

func isNumeric(character byte) bool {
	return '0' <= character && character <= '9'
}

func (lexerInstance *Lexer) getInteger() string {
	return lexerInstance.getMultiCharacterToken(isNumeric)
}

func (lexerInstance *Lexer) handleWordToken(lineNumber int, columnNumber int) token {
	word := lexerInstance.getWord()

	lexerInstance.currentColumnNumber += len(word)

	if _, isKeyword := keywords[word]; isKeyword {
		return token{
			kind:         word,
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

func (lexerInstance *Lexer) handleSingleCodePoint(character byte, lineNumber int, columnNumber int) token {
	if character != 0 {
		lexerInstance.currentColumnNumber += 1

		lexerInstance.updateCurrentCharacter()
	}

	switch character {
	case '=':
		return token{kind: assign, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	case ',':
		return token{kind: comma, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
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

	case 0:
		err := lexerInstance.file.Close()
		check(err)
		return token{kind: eof, filePath: lexerInstance.filePath, lineNumber: lineNumber, columnNumber: columnNumber}
	default:
		return token{unknown, string(character), lexerInstance.filePath, lineNumber, columnNumber}
	}
}

func (lexerInstance *Lexer) getNextToken() token {
	lexerInstance.skipWhitespace()

	character := lexerInstance.currentCharacter
	lineNumber := lexerInstance.currentLineNumber
	columnNumber := lexerInstance.currentColumnNumber

	if isAlphabetical(character) {
		return lexerInstance.handleWordToken(lineNumber, columnNumber)
	}

	if isNumeric(character) {
		return lexerInstance.handleNumberToken(lineNumber, columnNumber)
	}

	return lexerInstance.handleSingleCodePoint(character, lineNumber, columnNumber)
}
