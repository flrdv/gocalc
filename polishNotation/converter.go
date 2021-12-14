package polishNotation

import (
	"example.com/lex"
	"fmt"
)


func countOps(tokens []lex.Token) uint {
	var opsCount uint

	for _, token := range tokens {
		if token.Type == lex.Operator {
			opsCount++
		}
	}

	return opsCount
}

func removeValues(arr []lex.Token, indexes... int) []lex.Token {
	var arrWithNoValues []lex.Token

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

func pickMostPriorityOp(tokens []lex.Token) (opIndex int) {
	var mostPriorityOp int
	var mostPriority = LowPriority - 1

	for index, token := range tokens {
		// a bit unoptimized, as I need only each second value
		// but nobody cares, this operation isn't in interpreting
		// runtime, and go is fast enough
		if token.Type == lex.Operator {
			priority, found := OpsPriorities[token.Value.(lex.TokenType)]

			if !found {
				panic(fmt.Errorf("unknown operator: %s", token.Value))
			}

			if priority > mostPriority {
				mostPriorityOp = index
				mostPriority = priority
			}
		} else if token.Type == lex.InBraceExpr {
			return index
		}
	}

	return mostPriorityOp
}

func ConvertToPolishNotation(tokens []lex.Token) []lex.Token {
	if len(tokens) == 0 || len(tokens) == 1 {
		return tokens
	}

	var tokensQueue = tokens[:]
	var polishNotatedTokens []lex.Token
	var opsCount = countOps(tokens)

	for i := uint(0); i <= opsCount; i++ {
		if len(tokensQueue) == 0 {
			return polishNotatedTokens
		}

		opIndex := pickMostPriorityOp(tokensQueue)
		op := tokensQueue[opIndex]

		if op.Type == lex.InBraceExpr {
			polishNotatedTokens = append(polishNotatedTokens, ConvertToPolishNotation(op.Value.([]lex.Token))...)

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
