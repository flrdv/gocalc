package polishNotation

import (
	"fmt"
	"github.com/floordiv/gocalc/types"
)


func countOps(tokens []types.Token) uint {
	var opsCount uint

	for _, token := range tokens {
		if token.Type == types.Operator {
			opsCount++
		}
	}

	return opsCount
}

func removeValues(arr []types.Token, indexes... int) []types.Token {
	var arrWithNoValues []types.Token

	for elementIndex, element := range arr {
		var skip = false

		for _, index := range indexes {
			if index == elementIndex {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		arrWithNoValues = append(arrWithNoValues, element)
	}

	return arrWithNoValues
}

func pickMostPriorityOp(tokens []types.Token) (opIndex int) {
	var mostPriorityOp int
	var mostPriority = LowPriority - 1

	for index, token := range tokens {
		// a bit unoptimized, as I need only each second value
		// but nobody cares, this operation isn't in interpreting
		// runtime, and go is fast enough
		if token.Type == types.Operator {
			priority, found := OpsPriorities[token.Value.(types.TokenType)]

			if !found {
				panic(fmt.Errorf("unknown operator: %s", token.Value))
			}

			if priority > mostPriority {
				mostPriorityOp = index
				mostPriority = priority
			}
		} else if token.Type == types.InBraceExpr {
			return index
		}
	}

	return mostPriorityOp
}

func ConvertToPolishNotation(tokens []types.Token) []types.Token {
	if len(tokens) == 0 || len(tokens) == 1 {
		return tokens
	}

	var tokensQueue = tokens[:]
	var polishNotatedTokens []types.Token
	var opsCount = countOps(tokens)

	for i := uint(0); i <= opsCount; i++ {
		if len(tokensQueue) == 0 {
			return polishNotatedTokens
		}

		opIndex := pickMostPriorityOp(tokensQueue)
		op := tokensQueue[opIndex]

		if op.Type == types.InBraceExpr {
			polishNotatedTokens = append(polishNotatedTokens, ConvertToPolishNotation(op.Value.([]types.Token))...)

			if opIndex > 0 {
				beforeBracesTokens, beforeBracesOp := tokensQueue[:opIndex-1], tokensQueue[opIndex-1]
				polishNotatedTokens = append(polishNotatedTokens, ConvertToPolishNotation(beforeBracesTokens)...)
				polishNotatedTokens = append(polishNotatedTokens, beforeBracesOp)
			}

			tokensQueue = tokensQueue[opIndex+1:]
		} else {
			opBegin := opIndex - 1
			opEnd := opIndex + 1

			if opEnd >= len(tokensQueue) {
				opEnd--
			}

			if opIndex == 0 {
				polishNotatedTokens = append(polishNotatedTokens, tokensQueue[opEnd], op)
			} else {
				left, right := tokensQueue[opBegin], tokensQueue[opEnd]
				polishNotatedTokens = append(polishNotatedTokens, left, right, op)
			}

			indexesOfElementsToRemove := []int{opBegin, opEnd}

			if opBegin < opIndex {
				indexesOfElementsToRemove = append(indexesOfElementsToRemove, opIndex)
			}
			tokensQueue = removeValues(tokensQueue, indexesOfElementsToRemove...)
		}
	}

	return polishNotatedTokens
}
