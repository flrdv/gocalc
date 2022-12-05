package operator

type Operator string

const (
	Plus         Operator = "+"
	PlusPlus     Operator = "++"
	Minus        Operator = "-"
	MinusMinus   Operator = "--"
	Slash        Operator = "/"
	SlashSlash   Operator = "//"
	Star         Operator = "*"
	Caret        Operator = "^"
	Dot          Operator = "."
	Comma        Operator = ","
	Equal        Operator = "="
	EqualEqual   Operator = "=="
	NotEqual     Operator = "!="
	Greater      Operator = ">"
	GreaterEqual Operator = ">="
	Less         Operator = "<"
	LessEqual    Operator = "<="
)

var (
	allOperators = []Operator{
		Plus, PlusPlus, Minus, MinusMinus, Slash, SlashSlash, Star, Caret, Dot, Comma,
		Equal, EqualEqual, NotEqual, Greater, GreaterEqual, Less, LessEqual,
	}

	symbols = buildOperatorsCharsTable(allOperators...)
)

func IsOperator(str string) bool {
	for _, operator := range allOperators {
		if str == string(operator) {
			return true
		}
	}

	return false
}

func IsOperatorChar(char string) bool {
	_, found := symbols[char]
	return found
}

func buildOperatorsCharsTable(operators ...Operator) map[string]struct{} {
	set := make(map[string]struct{})

	for _, operator := range operators {
		for _, char := range operator {
			set[string(char)] = struct{}{}
		}
	}

	return set
}
