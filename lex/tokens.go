package lex

/*
Primary token is a token that is used on first lexical parsing
step, it's value is a raw string

Also it's type is only Operator, NotOperator, and NoType after init
*/
type PrimaryToken struct {
	Type TokenType
	Value string
}

/*
And this token is a final token, the difference between this and
primary token that it's value may have any type (bool, int, str,
float, etc.)
*/
type Token struct {
	Type TokenType
	Value interface{}
}


type BracesExpr struct {
	Type TokenType
	Value []Token
}
