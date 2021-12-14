package lex

import (
	"fmt"
	"strconv"
	"strings"
)


func parseTokens(rawString string) []PrimaryToken {
	var lexemes []PrimaryToken
	currentToken := PrimaryToken{
		Type: NoType,
		Value: "",
	}
	operators := []string { "+", "-", "/", "*", "**" }
	operatorsChars := strings.Join(operators, "")
	specialCharacters := map[string]TokenType {
		"(": Brace,
		")": Brace,
	}

	for _, char := range rawString {
		if string(char) == " " {
			continue
		}

		if res, err := specialCharacters[string(char)]; err {
			if len(currentToken.Value) > 0 {
				if currentToken.Type == NoType {
					currentToken.Type = getPrimaryTypeOfToken(operators, currentToken.Value)
				}

				lexemes = append(lexemes, currentToken)
			}

			lexemes = append(lexemes, PrimaryToken{
				Type:  res,
				Value: string(char),
			})

			continue
		}

		if currentToken.Type == Operator && !strings.Contains(operatorsChars, string(char)) {
			lexemes = append(lexemes, currentToken)
			currentToken = PrimaryToken{
				Type: NoType,
				Value: string(char),
			}
		} else if strings.Contains(operatorsChars, string(char)) {
			if currentToken.Type == Operator {
				currentToken.Value += string(char)
			} else if currentToken.Type == NoType && len(currentToken.Value) == 0 {
				currentToken.Type = Operator
				currentToken.Value += string(char)
			} else {
				currentToken.Type = NotOperator
				lexemes = append(lexemes, currentToken)
				currentToken = PrimaryToken{
					Type: Operator,
					Value: string(char),
				}
			}
		} else if string(char) == " " {
			if len(currentToken.Value) != 0 {
				lexemes = append(lexemes, currentToken)
				currentToken = PrimaryToken{
					Type: NoType,
					Value: "",
				}
			}
		} else {
			if currentToken.Type == NoType && len(currentToken.Value) > 0 {
				currentToken.Type = NotOperator
			}

			currentToken.Value += string(char)
		}
	}

	for _, operator := range operators {
		if currentToken.Value == operator {
			currentToken.Type = Operator
			break
		}
	}

	if currentToken.Type != Operator {
		currentToken.Type = NotOperator
	}

	lexemes = append(lexemes, currentToken)

	return lexemes
}

func getPrimaryTypeOfToken(operators []string, value string) TokenType {
	for _, elem := range operators {
		if value == elem {
			return Operator
		}
	}

	return NotOperator
}

func getTypeOfToken(tokenType TokenType, tokenValue string) TokenType {
	switch tokenType {
	case Operator:
		switch tokenValue {
		case "+": return OpAdd
		case "-": return OpMin
		case "/": return OpDiv
		case "*": return OpMul
		case "**": return OpPow
		}
	case NotOperator:
		switch {
		default:
			return Float
		}
	case Brace:
		switch tokenValue {
		case "(": return LBrace
		case ")": return RBrace
		}
	}

	return NoType
}

func Parse(rawString string) []Token {
	primaryTokens := parseTokens(rawString)
	var outputTokens []Token

	for _, token := range primaryTokens {
		tokenType := getTypeOfToken(token.Type, token.Value)

		if tokenType == NoType {
			panic(fmt.Sprintf("unrecognized token: %s\n", token.Value))
		}

		var parsedValue interface{}

		// ignoring errors from parsing integers, floats, etc. only because
		// I am already sure it's valid as validation is happening before this moment
		switch tokenType {
		case Float:
			parsedValue, _ = strconv.ParseFloat(token.Value, 64)
		default:
			parsedValue = tokenType
		}

		switch token.Type {
		case Operator: tokenType = Operator
		case Brace: tokenType = Brace
		}

		newToken := Token{
			Type: tokenType,
			Value: parsedValue,
		}

		outputTokens = append(outputTokens, newToken)
	}

	return outputTokens
}
