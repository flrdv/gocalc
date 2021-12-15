package semantic

import (
	"errors"
	"fmt"
	"github.com/floordiv/gocalc/types"
)


func getClosingBraceIndex(tokens []types.Token) (int, error) {
	/*
	Returns reference on index of closing brace, or error
	 */

	openedBraces := 0

	for index, token := range tokens {
		switch token.Type {
		case types.Brace:
			switch token.Value.(types.TokenType) {
			case types.LBrace: openedBraces++
			case types.RBrace: openedBraces--; if openedBraces == 0 { return index, nil }
			}
		}
	}

	return 0, errors.New("no closing brace found")
}

func Parse(tokens []types.Token) []types.Token {
	/*
	Currently the only thing this function does is just puts in-braces expressions
	into the single token
	 */

	var skipIters int
	var outputTokens []types.Token

	for index, token := range tokens {
		if skipIters > 0 {
			// in case we parsed nested braces, we're passing already parsed tokens
			skipIters--
			continue
		}

		if token.Type == types.Brace && token.Value.(types.TokenType) == types.LBrace {
			// we found an in-brace expression, time to parse it
			closingBraceIndex, err := getClosingBraceIndex(tokens[index:])

			if err != nil {
				// no closing brace
				panic(fmt.Errorf("syntax error: %s", err))
			} else {
				skipIters = closingBraceIndex + 1	// +1 to avoid touching RBrace token
				begin := index + 1					// the same reason

				if index > 0 {
					begin++ 						// if braces aren't the first ones in the expression
				}									// we need to inc as otherwise we'll touch prev token

				inBracesExpr := Parse(tokens[begin:index + closingBraceIndex])
				outputTokens = append(outputTokens, types.Token{
					Type:  types.InBraceExpr,
					Value: inBracesExpr,
				})
			}
		} else {
			outputTokens = append(outputTokens, token)
		}
	}

	return outputTokens
}
