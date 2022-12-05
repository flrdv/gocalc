package tokenizer

import operator2 "github.com/floordiv/gocalc/frontend/internal/operator"

func isNumber(char letter) bool {
	return char >= '0' && char <= '9'
}

func isString(char letter) bool {
	return char == '"'
}

func isOperator(char letter) bool {
	return operator2.IsOperatorChar(string(char))
}

func isIdentifier(char letter) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

func isWhitespace(char letter) bool {
	switch char {
	// not all will be used, but here we only need to match them all
	case ' ', '\t', '\n', '\r', '\v', '\f', '\b':
		return true
	}

	return false
}
