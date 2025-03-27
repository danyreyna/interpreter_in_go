package lexer

type Lexer struct {
	codeString            string
	currentCharacterIndex int
	currentCharacter      byte
	currentLineNumber     int
	currentColumnNumber   int
}

func newLexer(codeString string) *Lexer {
	var currentCharacter byte
	if len(codeString) == 0 {
		currentCharacter = 0
	} else {
		currentCharacter = codeString[0]
	}

	lexerInstance := &Lexer{
		codeString,
		0,
		currentCharacter,
		1,
		1,
	}

	return lexerInstance
}

func (lexerInstance *Lexer) updateCurrentCharacter() {
	if (lexerInstance.currentCharacterIndex + 1) >= len(lexerInstance.codeString) {
		lexerInstance.currentCharacter = 0
	} else {
		lexerInstance.currentCharacter = lexerInstance.codeString[lexerInstance.currentCharacterIndex+1]
	}

	lexerInstance.currentCharacterIndex += 1
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
	startIndex := lexerInstance.currentCharacterIndex

	for isAllowedCharacter(lexerInstance.currentCharacter) {
		lexerInstance.updateCurrentCharacter()
	}

	return lexerInstance.codeString[startIndex:lexerInstance.currentCharacterIndex]
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
			lineNumber:   lineNumber,
			columnNumber: columnNumber,
		}
	}

	return token{
		identifier,
		word,
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
		return token{kind: assign, lineNumber: lineNumber, columnNumber: columnNumber}
	case ',':
		return token{kind: comma, lineNumber: lineNumber, columnNumber: columnNumber}
	case '{':
		return token{kind: leftCurlyBrace, lineNumber: lineNumber, columnNumber: columnNumber}
	case '(':
		return token{kind: leftParenthesis, lineNumber: lineNumber, columnNumber: columnNumber}
	case '+':
		return token{kind: plus, lineNumber: lineNumber, columnNumber: columnNumber}
	case '}':
		return token{kind: rightCurlyBrace, lineNumber: lineNumber, columnNumber: columnNumber}
	case ')':
		return token{kind: rightParenthesis, lineNumber: lineNumber, columnNumber: columnNumber}
	case ';':
		return token{kind: semicolon, lineNumber: lineNumber, columnNumber: columnNumber}

	case 0:
		return token{kind: eof, lineNumber: lineNumber, columnNumber: columnNumber}
	default:
		return token{unknown, string(character), lineNumber, columnNumber}
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
