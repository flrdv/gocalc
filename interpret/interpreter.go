package interpret

import (
	"errors"
	"fmt"
	"github.com/floordiv/gocalc/types"
	"math"
)


func doOperation(op types.TokenType, left, right float64) (float64, error) {
	switch op {
	case types.OpAdd: return left + right, nil
	case types.OpMin: return left - right, nil

	case types.OpDiv: return left / right, nil
	case types.OpMul: return left * right, nil

	case types.OpPow: return math.Pow(left, right), nil
	}

	return float64(0), errors.New(fmt.Sprintf("unrecognized operator: %s", op))
}



func Interpret(tokens []types.Token) interface{} {
	/*
	tokens must be in polish notation format
	 */

	var valuesStack []float64

	for _, token := range tokens {
		if token.Type == types.Operator {
			if len(valuesStack) < 2 {
				panic(fmt.Errorf("invalid polish notation"))
			}

			valuesStackLen := len(valuesStack)
			right, left := valuesStack[valuesStackLen-1], valuesStack[valuesStackLen-2]
			opResult, err := doOperation(token.Value.(types.TokenType), right, left)

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
