package parser

import (
	"strings"
)

func convertExpr(expression string) string {
	for _, replace := range [][]string{{" ", ""}, {"--", "+"}, {"+-", "-"}} {
		expression = strings.ReplaceAll(expression, replace[0], replace[1])
	}

	return expression
}

func priorityOperator(sym string) int {
	switch sym {
	case "*", "/":
		return 1
	case "(":
		return 5
	case ")":
		return -5
	default:
		return 0
	}

}
