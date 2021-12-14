package lex


type TokenType string

const (
	Float TokenType = "Float"
	Variable = "Variable"
	Operator = "Operator"
	NotOperator = "NotOperator"
	Brace = "Brace"
	InBraceExpr = "InBraceExpr"
	NoType = "NoType"

	LBrace = "LBrace"
	RBrace = "RBrace"

	OpUnary = "OpUnary"
	OpAdd = "OpAdd"
	OpMin = "OpMin"
	OpDiv = "OpDiv"
	OpMul = "OpMul"
	OpPow = "OpPow"
)
