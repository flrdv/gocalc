package main

import (
	"fmt"
	"github.com/floordiv/gocalc/interpret"
	"github.com/floordiv/gocalc/lex"
	"github.com/floordiv/gocalc/polishNotation"
	"github.com/floordiv/gocalc/semantic"
	"github.com/floordiv/gocalc/types"
)


func calculate(expr string, printParsedNotation bool) interface{} {
	parsedSemantic := semantic.Parse(lex.Parse(expr))
	// Reversed Polish Notation - like ordinary, but postfix
	rpn := polishNotation.ConvertToPolishNotation(parsedSemantic)

	if printParsedNotation {
		fmt.Print("reverse polish notation> ")

		for _, token := range rpn {
			tValue := token.Value

			if token.Type == types.Operator {
				switch tValue.(types.TokenType) {
				case types.OpAdd:
					tValue = "+"
				case types.OpMin:
					tValue = "-"
				case types.OpMul:
					tValue = "*"
				case types.OpDiv:
					tValue = "/"
				case types.OpPow:
					tValue = "**"
				}
			}

			fmt.Print(tValue, " ")
		}
		fmt.Println()
	}

	return interpret.Interpret(rpn)
}

func main() {
	for {
		var expr string

		fmt.Print("expr> ")
		_, err := fmt.Scanln(&expr)

		if err != nil {
			fmt.Println("Invalid expression!", err)
			continue
		}
		if expr == "q" || expr == "quit" {
			return
		}

		fmt.Println("result>", calculate(expr, true))
	}
}
