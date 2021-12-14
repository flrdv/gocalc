package polishNotation

import (
	"github.com/floordiv/gocalc/lex"
)

type TokenPriority int8
const (
	LowPriority TokenPriority = iota + 1
	MediumPriority
	HighPriority
	MaxPriority
)


var OpsPriorities = map[lex.TokenType]TokenPriority {
	lex.OpAdd: LowPriority,
	lex.OpMin: LowPriority,

	lex.OpMul: MediumPriority,
	lex.OpDiv: MediumPriority,
	lex.OpUnary: MediumPriority,

	lex.OpPow: HighPriority,

	lex.InBraceExpr: MaxPriority,
}
