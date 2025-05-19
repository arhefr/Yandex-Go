package service

import (
	"strconv"

	"github.com/arhefr/Yandex-Go/agent/internal/model"
)

func MakeTask(task *model.Task) string {
	var res float64

	n1, n2 := task.Arg1, task.Arg2

	switch task.Oper {
	case "*":
		res = n1 * n2
	case "/":
		if n2 == 0 {
			return ""
		}
		res = n1 / n2
	case "+":
		res = n1 + n2
	case "-":
		res = n1 - n2
	}

	return strconv.FormatFloat(res, 'f', -1, 64)
}
