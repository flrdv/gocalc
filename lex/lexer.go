package lex

import (
	"fmt"
	"github.com/floordiv/gocalc/types"
	"strconv"
	"strings"
)


func parseTokens(rawString string) []types.PrimaryToken {
	var lexemes []types.PrimaryToken
	currentToken := types.PrimaryToken{
		Type:  types.NoType,
		Value: "",
	}
	operators := []string { "+", "-", "/", "*", "**" }
	operatorsChars := strings.Join(operators, "")
	specialCharacters := map[string]types.TokenType{
		"(": types.Brace,
		")": types.Brace,
	}

	for _, char := range rawString {
		char := string(char)

		if char == " " {
			continue
		}

		if res, err := specialCharacters[char]; err {
			if len(currentToken.Value) > 0 {
				if currentToken.Type == types.NoType {
					currentToken.Type = getPrimaryTypeOfToken(operators, currentToken.Value)
				}

				lexemes = append(lexemes, currentToken)
			}

			lexemes = append(lexemes, types.PrimaryToken{
				Type:  res,
				Value: char,
			})

			continue
		}

		if currentToken.Type == types.Operator && !strings.Contains(operatorsChars, char) {
			lexemes = append(lexemes, currentToken)
			currentToken = types.PrimaryToken{
				Type:  types.NoType,
				Value: char,
			}
		} else if strings.Contains(operatorsChars, char) {
			if currentToken.Type == types.Operator {
				currentToken.Value += char
			} else if currentToken.Type == types.NoType && len(currentToken.Value) == 0 {
				currentToken.Type = types.Operator
				currentToken.Value += char
			} else {
				currentToken.Type = types.NotOperator
				lexemes = append(lexemes, currentToken)
				currentToken = types.PrimaryToken{
					Type:  types.Operator,
					Value: char,
				}
			}
		} else if char == " " {
			if len(currentToken.Value) != 0 {
				lexemes = append(lexemes, currentToken)
				currentToken = types.PrimaryToken{
					Type:  types.NoType,
					Value: "",
				}
			}
		} else {
			if currentToken.Type == types.NoType && len(currentToken.Value) > 0 {
				currentToken.Type = types.NotOperator
			}

			currentToken.Value += char
		}
	}

	for _, operator := range operators {
		if currentToken.Value == operator {
			currentToken.Type = types.Operator
			break
		}
	}

	if currentToken.Type != types.Operator {
		currentToken.Type = types.NotOperator
	}

	lexemes = append(lexemes, currentToken)

	return lexemes
}

func getPrimaryTypeOfToken(operators []string, value string) types.TokenType {
	for _, elem := range operators {
		if value == elem {
			return types.Operator
		}
	}

	return types.NotOperator
}

func getTypeOfToken(tokenType types.TokenType, tokenValue string) types.TokenType {
	switch tokenType {
	case types.Operator:
		switch tokenValue {
		case "+": return types.OpAdd
		case "-": return types.OpMin
		case "/": return types.OpDiv
		case "*": return types.OpMul
		case "**": return types.OpPow
		}
	case types.NotOperator: return types.Float
	case types.Brace:
		switch tokenValue {
		case "(": return types.LBrace
		case ")": return types.RBrace
		}
	}

	return types.NoType
}

func Parse(rawString string) []types.Token {
	primaryTokens := parseTokens(rawString)
	var outputTokens []types.Token

	for _, token := range primaryTokens {
		tokenType := getTypeOfToken(token.Type, token.Value)

		if tokenType == types.NoType {
			panic(fmt.Sprintf("unrecognized token: %s\n", token.Value))
		}

		var parsedValue interface{}

		// ignoring errors from parsing integers, floats, etc. only because
		// I am already sure it's valid as validation is happening before this moment
		switch tokenType {
		case types.Float:
			parsedValue, _ = strconv.ParseFloat(token.Value, 64)
		default:
			parsedValue = tokenType
		}

		switch token.Type {
		case types.Operator: tokenType = types.Operator
		case types.Brace: tokenType = types.Brace
		}

		newToken := types.Token{
			Type: tokenType,
			Value: parsedValue,
		}

		outputTokens = append(outputTokens, newToken)
	}

	return outputTokens
}
