package interpret

import (
	"errors"
	"fmt"
	"github.com/floordiv/gocalc/lex"
	"math"
)


func doOperation(op lex.TokenType, left, right float64) (float64, error) {
	switch op {
	case lex.OpAdd: return left + right, nil
	case lex.OpMin: return left - right, nil

	case lex.OpDiv: return left / right, nil
	case lex.OpMul: return left * right, nil

	case lex.OpPow: return math.Pow(left, right), nil
	}

	return float64(0), errors.New(fmt.Sprintf("unrecognized operator: %s", op))
}



func Interpret(tokens []lex.Token) interface{} {
	/*
	tokens must be in polish notation format
	 */

	var valuesStack []float64

	for _, token := range tokens {
		if token.Type == lex.Operator {
			if len(valuesStack) < 2 {
				panic(fmt.Errorf("invalid syntax of polish notation"))
			}

			valuesStackLen := len(valuesStack)
			right, left := valuesStack[valuesStackLen-1], valuesStack[valuesStackLen-2]
			opResult, err := doOperation(token.Value.(lex.TokenType), right, left)

			if err != nil {
				panic(err)
			}

			valuesStack = append(valuesStack[:valuesStackLen-2], opResult)
		} else {
			valuesStack = append(valuesStack, token.Value.(float64))
		}
	}

	return valuesStack[0]
}
