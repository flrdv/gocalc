package polishNotation

import (
	"github.com/floordiv/gocalc/types"
)

type TokenPriority int8
const (
	LowPriority TokenPriority = iota + 1
	MediumPriority
	HighPriority
	MaxPriority
)


var OpsPriorities = map[types.TokenType]TokenPriority {
	types.OpAdd: LowPriority,
	types.OpMin: LowPriority,

	types.OpMul:   MediumPriority,
	types.OpDiv:   MediumPriority,
	types.OpUnary: MediumPriority,

	types.OpPow: HighPriority,

	types.InBraceExpr: MaxPriority,
}
