package tokenizer

import (
	"errors"
	"github.com/floordiv/gocalc/frontend/internal/operator"
	"io"
	"strconv"
	"strings"
)

type Tokenizer interface {
	Next() (LexemeType, Lexeme, error)
}

type tokenizer struct {
	input string
}

func NewTokenizer(input string) Tokenizer {
	return &tokenizer{
		input: input,
	}
}

func (t *tokenizer) Next() (LexemeType, Lexeme, error) {
	if len(t.input) == 0 {
		return unknown, "", io.EOF
	}

	t.trimSpaces()
	lexemeType := letter(t.input[0]).Type()
	switch lexemeType {
	case String:
		lexeme, err := t.parseString()
		return String, lexeme, err
	case Operator:
		lexeme, err := t.parseOperator()
		return Operator, lexeme, err
	case Integer:
		lexeme, err := t.parseInteger()
		return Integer, lexeme, err
	case Identifier:
		lexeme, err := t.parseIdentifier()
		return Identifier, lexeme, err
	default:
		panic(
			"BUG: frontend/internal/tokenizer.go:Next(): unknown lexeme type: " + string(rune(lexemeType)),
		)
	}
}

// trim the leading spaces
func (t *tokenizer) trimSpaces() {
	t.input = strings.TrimLeft(t.input, " ")
}

func (t *tokenizer) parseString() (string, error) {
	// when this function is called, we suppose that the first character is an opening
	// quote. Let's check it to make sure
	if t.input[0] != '"' {
		panic(
			"BUG: frontend/internal/tokenizer.go:parseString(): unexpected string quote: " +
				strconv.Quote(string(t.input[0])),
		)
	}

	// strings are not that simple they pretended to be. We have to handle plain-text escape
	// characters as they have been literals
	buffer := []byte{'"'}

	for i := 1; i < len(t.input); i++ {
		char := t.input[i]

		switch char {
		case '\\':
			i++
			escapeChar, err := literalAsEscapeChar(t.input[i])
			if err != nil {
				return "", err
			}

			char = escapeChar
		case '"':
			t.input = t.input[i+1:]
			buffer = append(buffer, '"')

			return string(buffer), nil
		}

		buffer = append(buffer, char)
	}

	// in case we did not exit from function directly from for-loop, our string has no closer
	// TODO: add better errors that can tell us more information about source of the error
	return "", errors.New("string is not closed")
}

func (t *tokenizer) parseOperator() (Lexeme, error) {
	if letter(t.input[0]).Type() != Operator {
		panic(
			"BUG: frontend/internal/tokenizer.go:parseOperator(): first character of input is not an operator",
		)
	}

	for i := 1; i < len(t.input); i++ {
		char := t.input[i]

		if letter(char).Type() != Operator {
			lexeme := t.input[:i]
			t.input = t.input[i:]

			return lexeme, nil
		} else if !operator.IsOperator(t.input[:i+1]) {
			lexeme := t.input[:i]
			t.input = t.input[i:]

			return lexeme, nil
		}
	}

	return "", errors.New("syntax error: operators in the end of the expression")
}

func (t *tokenizer) parseInteger() (Lexeme, error) {
	if !isNumber(letter(t.input[0])) {
		panic(
			"BUG: frontend/internal/tokenizer.go:parseInteger(): first character of input is not a number",
		)
	}

	for i := range t.input {
		if letter(t.input[i]).Type() == Integer {
			continue
		}

		num := t.input[:i]
		t.input = t.input[i:]

		return num, nil
	}

	num := t.input
	t.input = ""

	return num, nil
}

func (t *tokenizer) parseIdentifier() (Lexeme, error) {
	if !isIdentifier(letter(t.input[0])) {
		panic(
			"BUG: frontend/internal/tokenizer.go:parseIdentifier(): first character of input is not an identifier",
		)
	}

	for i := range t.input {
		switch letter(t.input[i]).Type() {
		case Identifier, Integer:
		default:
			ident := t.input[:i]
			t.input = t.input[i:]

			return ident, nil
		}
	}

	ident := t.input
	t.input = ""

	return ident, nil
}

func literalAsEscapeChar(literal byte) (byte, error) {
	switch literal {
	case 'n':
		return '\n', nil
	case 'r':
		return '\r', nil
	case 't':
		return '\t', nil
	case 'v':
		return '\v', nil
	case 'b':
		return '\b', nil
	case 'f':
		return '\f', nil
	case '"':
		return '"', nil
	}

	return '\x00', errors.New("unknown escape character: \\" + string(literal))
}
