package lex

import "strconv"


func isInt(str string) bool {
	_, err := strconv.ParseInt(str, 10, 64)

	return err == nil
}

func isFloat(str string) bool {
	_, err := strconv.ParseFloat(str, 64)

	return err == nil
}

