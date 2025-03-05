package service

import (
	agent "calculator/internal/agent/config"
	"calculator/internal/orchestrator/transport/http/models"
	"fmt"
	"time"
)

func MakeTask(task models.Task, operation_time agent.OperationTime) (float64, error) {
	var res float64
	n1, n2 := task.Operand1, task.Operand2

	switch task.Operation {
	case "*":
		time.Sleep(operation_time.Mul)
		res = n1 * n2
	case "/":
		time.Sleep(operation_time.Div)
		if n2 == 0.0 {
			return 0.0, fmt.Errorf("error division by zero")
		}
		res = n1 / n2
	case "+":
		time.Sleep(operation_time.Add)
		res = n1 + n2
	case "-":
		time.Sleep(operation_time.Sub)
		res = n1 - n2
	}

	return res, nil
}
