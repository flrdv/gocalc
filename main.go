package main

import (
	"example.com/interpret"
	"example.com/lex"
	"example.com/polishNotation"
	"example.com/semantic"
	"fmt"
)


func calculate(expr string, printParsedNotation bool) interface{} {
	parsedSemantic := semantic.Parse(lex.Parse(expr))
	ast := polishNotation.ConvertToPolishNotation(parsedSemantic)

	if printParsedNotation {
		fmt.Print("reverse polish notation> ")

		for _, token := range ast {
			tValue := token.Value

			if token.Type == lex.Operator {
				switch tValue.(lex.TokenType) {
				case lex.OpAdd:
					tValue = "+"
				case lex.OpMin:
					tValue = "-"
				case lex.OpMul:
					tValue = "*"
				case lex.OpDiv:
					tValue = "/"
				case lex.OpPow:
					tValue = "**"
				}
			}

			fmt.Print(tValue, " ")
		}
		fmt.Println()
	}

	return interpret.Interpret(ast)
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
