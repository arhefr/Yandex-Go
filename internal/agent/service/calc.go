package service

import (
	"strconv"
	"time"

	"github.com/arhefr/Yandex-Go/internal/orchestrator/model"
)

func MakeTask(task *model.Task, operTime OperTime) string {
	var res float64

	n1, n2 := task.Arg1, task.Arg2

	switch task.Oper {
	case "*":
		time.Sleep(operTime.Mul)
		res = n1 * n2
	case "/":
		time.Sleep(operTime.Div)
		res = n1 / n2
	case "+":
		time.Sleep(operTime.Add)
		res = n1 + n2
	case "-":
		time.Sleep(operTime.Sub)
		res = n1 - n2
	}

	return strconv.FormatFloat(res, 'f', -1, 64)
}
