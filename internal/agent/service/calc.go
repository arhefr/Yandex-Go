package service

import (
	"strconv"
	"time"

	models_agent "github.com/arhefr/Yandex-Go/internal/agent/model"
	models_orchestrator "github.com/arhefr/Yandex-Go/internal/orchestrator/model"
)

func MakeTask(task *models_orchestrator.Task, operation_time models_agent.OperationTime) string {
	var res float64

	n1, n2 := task.Arg1, task.Arg2

	switch task.Oper {
	case "*":
		time.Sleep(operation_time.Mul)
		res = n1 * n2
	case "/":
		time.Sleep(operation_time.Div)
		res = n1 / n2
	case "+":
		time.Sleep(operation_time.Add)
		res = n1 + n2
	case "-":
		time.Sleep(operation_time.Sub)
		res = n1 - n2
	}

	return strconv.FormatFloat(res, 'f', -1, 64)
}
