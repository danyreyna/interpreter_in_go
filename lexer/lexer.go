package lexer

type Lexer struct {
	codeString            string
	currentCharacterIndex int
	currentCharacter      byte
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

func (lexerInstance *Lexer) handleWordToken() token {
	word := lexerInstance.getWord()

	if _, isKeyword := keywords[word]; isKeyword {
		return token{kind: word}
	}

	return token{identifier, word}
}

func (lexerInstance *Lexer) handleNumberToken() token {
	number := lexerInstance.getInteger()
	return token{integer, number}
}

func (lexerInstance *Lexer) handleSingleCodePoint(character byte) token {
	lexerInstance.updateCurrentCharacter()

	switch character {
	case '=':
		return token{kind: assign}
	case ',':
		return token{kind: comma}
	case '{':
		return token{kind: leftCurlyBrace}
	case '(':
		return token{kind: leftParenthesis}
	case '+':
		return token{kind: plus}
	case '}':
		return token{kind: rightCurlyBrace}
	case ')':
		return token{kind: rightParenthesis}
	case ';':
		return token{kind: semicolon}

	case 0:
		return token{kind: eof}
	default:
		return token{unknown, string(character)}
	}
}

func (lexerInstance *Lexer) getNextToken() token {
	lexerInstance.skipWhitespace()

	character := lexerInstance.currentCharacter

	if isAlphabetical(character) {
		return lexerInstance.handleWordToken()
	}

	if isNumeric(character) {
		return lexerInstance.handleNumberToken()
	}

	return lexerInstance.handleSingleCodePoint(character)
}
