package parser

import (
	Err "github.com/arhefr/Yandex-Go/pkg/errors"
	"github.com/arhefr/Yandex-Go/pkg/tools"
)

func GetNumsOps(expression string) ([]float64, []Operator, error) {
	var (
		numbers   []string
		operators []Operator

		priority int
		num      string
	)

	expression = convertExpr(expression)
	for i := 0; i < len(expression); i++ {
		sym := string(expression[i])

		switch sym {
		case ".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			num += sym

			if i == len(expression)-1 {
				numbers = append(numbers, num)
				num = ""
			}

		case "-", "+":
			if num == "" {
				num += sym
				if i == len(expression)-1 {
					numbers = append(numbers, num)
					num = ""
				} else if i != 0 {
					operators = append(operators, NewOperator("+", len(operators), priority))
				}
			} else {
				numbers = append(numbers, num)
				num = sym
				operators = append(operators, NewOperator("+", len(operators), priority))
			}

		default:
			if num != "" {
				if num == "+" || num == "-" {
					operators = append(operators, NewOperator(num, len(operators), priority))
					num = ""
				} else {
					numbers = append(numbers, num)
					num = ""
				}
			}

			switch sym {
			case "(", ")":
				priority += priorityOperator(sym)

			case "*", "/":
				if i != len(expression)-1 && sym == "/" && string(expression[i+1]) == "0" {
					return nil, nil, Err.IncorrectExpr
				}
				operators = append(operators, NewOperator(sym, len(operators), priorityOperator(sym)+priority))

			default:
				return nil, nil, Err.IncorrectExpr
			}
		}
	}

	if len(numbers)-1 != len(operators) || priority != 0 {
		return nil, nil, Err.IncorrectExpr
	}

	numbersFloat64, err := tools.SliceTypeToFloat64(numbers)
	if err != nil {
		return nil, nil, Err.IncorrectExpr
	}
	SortByPriority(operators)
	return numbersFloat64, operators, nil
}
