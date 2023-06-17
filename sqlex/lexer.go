package sqlex

import (
	"errors"
	"fmt"
)

const (
	TokenUnknown    = iota
	TokenIdentifier = iota
	TokenNumber     = iota
	TokenLParen     = iota
	TokenRParen     = iota
	TokenComma      = iota
	TokenSemicolon  = iota
)

func tokenTypeToString(t int) string {
	switch t {
	case TokenUnknown:
		return "TokenUnknown"
	case TokenIdentifier:
		return "TokenIdentifier"
	case TokenNumber:
		return "TokenNumber"
	case TokenLParen:
		return "TokenLParen"
	case TokenRParen:
		return "TokenRParen"
	case TokenComma:
		return "TokenComma"
	case TokenSemicolon:
		return "TokenSemicolon"
	default:
		return ""
	}
}

type SQLLexer struct {
	source       string
	sourceSize   int
	currentIndex int
	currentLine  int
}

type Token struct {
	Literal string
	Type    int
	Line    int
	Pos     int
}

func (t Token) ToString() string {
	return fmt.Sprintf("[%s](line:%d,pos:%d) -> %s", tokenTypeToString(t.Type), t.Line, t.Pos, t.Literal)
}

func NewSQLLexer(source string) *SQLLexer {
	return &SQLLexer{
		currentIndex: 0,
		currentLine:  1,
		source:       source,
		sourceSize:   len(source),
	}
}

func (self *SQLLexer) Reset() {
	self.currentIndex = 0
	self.currentLine = 1
}

func (self *SQLLexer) LoadAllTokens() []Token {
	tokens := make([]Token, 0)
	token, err := self.Next()
	for err == nil {
		tokens = append(tokens, token)
		token, err = self.Next()
	}
	return tokens
}

func (self *SQLLexer) Next() (Token, error) {
	if self.currentIndex >= self.sourceSize {
		return Token{}, errors.New("Index error")
	}
	for isWhitespace(self.source[self.currentIndex]) {
		if self.source[self.currentIndex] == '\n' {
			self.currentLine += 1
		}
		self.currentIndex += 1
	}

	token := Token{}
	switch self.source[self.currentIndex] {
	case '(':
		{
			token.Literal = string(self.source[self.currentIndex])
			token.Line = self.currentLine
			token.Pos = int(self.currentIndex) % self.currentLine
			token.Type = TokenLParen
			self.currentIndex += 1
		}
		break
	case ')':
		{
			token.Literal = string(self.source[self.currentIndex])
			token.Line = self.currentLine
			token.Pos = int(self.currentIndex) % self.currentLine
			token.Type = TokenRParen
			self.currentIndex += 1
		}
		break
	case ',':
		{
			token.Literal = string(self.source[self.currentIndex])
			token.Line = self.currentLine
			token.Pos = int(self.currentIndex) % self.currentLine
			token.Type = TokenComma
			self.currentIndex += 1
		}
		break
	case ';':
		{
			token.Type = TokenSemicolon
			token.Literal = string(self.source[self.currentIndex])
			token.Line = self.currentLine
			token.Pos = int(self.currentIndex) % self.currentLine
			self.currentIndex += 1
		}
		break
	default:
		{
			if isAlpha(self.source[self.currentIndex]) || self.source[self.currentIndex] == '_' {
				start := self.currentIndex
				for isAlphaNumeric(self.source[self.currentIndex]) || self.source[self.currentIndex] == '_' {
					self.currentIndex += 1
				}
				token.Literal = self.source[start:self.currentIndex]
				token.Line = self.currentLine
				token.Pos = int(self.currentIndex) % self.currentLine
				token.Type = TokenIdentifier
			} else if isNumeric(self.source[self.currentIndex]) {
				start := self.currentIndex
				for isNumeric(self.source[self.currentIndex]) {
					self.currentIndex += 1
				}
				token.Literal = self.source[start:self.currentIndex]
				token.Line = self.currentLine
				token.Pos = int(self.currentIndex) % self.currentLine
				token.Type = TokenNumber
			} else {
				token.Literal = string(self.source[self.currentIndex])
				token.Line = self.currentLine
				token.Pos = int(self.currentIndex) % self.currentLine
				token.Type = TokenUnknown
			}
		}
		break
	}
	return token, nil
}

func isAlpha(b byte) bool {
	return 'A' <= b && b <= 'Z' || 'a' <= b && b <= 'z'
}

func isNumeric(b byte) bool {
	return '0' <= b && b <= '9'
}

func isAlphaNumeric(b byte) bool {
	return isAlpha(b) || isNumeric(b)
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\n' || b == '\t'
}
