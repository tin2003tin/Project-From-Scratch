package compiler

import (
	"strings"
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	Terminals []string
}

func NewLexer(terminals []string) *Lexer {
	l := Lexer{
		Terminals: terminals,
	}
	return &l
}

func (l *Lexer) Convert(input string) []Token {
	tokens := make([]Token, 0)
	var currentToken strings.Builder
	var ones []rune

	for _, terminal := range l.Terminals {
		if len(terminal) == 1 {
			ones = append(ones, rune(terminal[0]))
		}
	}
	ones = append(ones, rune('$'))
	for _, char := range input {
		if unicode.IsSpace(char) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, l.getTokenType(currentToken.String()))
				currentToken.Reset()
			}
		} else if containsRune(ones, char) {
			if currentToken.Len() > 0 {
				tokens = append(tokens, l.getTokenType(currentToken.String()))
				currentToken.Reset()
			}
			tokens = append(tokens, Token{Type: string(char), Value: string(char)})
		} else {
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, l.getTokenType(currentToken.String()))
	}

	return tokens
}

func (l *Lexer) getTokenType(tokenValue string) Token {
	for _, terminal := range l.Terminals {
		if terminal == tokenValue {
			return Token{Type: terminal, Value: tokenValue}
		}
	}
	return Token{Type: "ID", Value: tokenValue}
}

func containsRune(runes []rune, r rune) bool {
	for _, run := range runes {
		if run == r {
			return true
		}
	}
	return false
}
