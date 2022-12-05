package tokenizer

type Lexeme = string

type letter byte

func (l letter) Type() LexemeType {
	if isNumber(l) {
		return Integer
	} else if isString(l) {
		return String
	} else if isOperator(l) {
		return Operator
	} else if isIdentifier(l) {
		return Identifier
	} else if isWhitespace(l) {
		return Whitespace
	}

	return unknown
}

type LexemeType int

const (
	unknown LexemeType = iota
	Integer
	String
	Operator
	Identifier
	Whitespace
)

func (l LexemeType) String() string {
	switch l {
	case unknown:
		return "unknown"
	case Integer:
		return "Integer"
	case String:
		return "String"
	case Operator:
		return "Operator"
	case Identifier:
		return "Identifier"
	case Whitespace:
		return "Whitespace"
	}

	return "???"
}
