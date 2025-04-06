package service

import (
	"strconv"
	"time"

	"github.com/arhefr/Yandex-Go/agent/internal/model"
)

// по хорошему, парсить с переменных окружения
var OPERTIME = struct {
	Add time.Duration
	Sub time.Duration
	Mul time.Duration
	Div time.Duration
}{
	time.Duration(100),
	time.Duration(100),
	time.Duration(100),
	time.Duration(100),
}

func MakeTask(task *model.Task) string {
	var res float64

	n1, n2 := task.Arg1, task.Arg2

	switch task.Oper {
	case "*":
		time.Sleep(OPERTIME.Mul)
		res = n1 * n2
	case "/":
		time.Sleep(OPERTIME.Div)
		res = n1 / n2
	case "+":
		time.Sleep(OPERTIME.Add)
		res = n1 + n2
	case "-":
		time.Sleep(OPERTIME.Sub)
		res = n1 - n2
	}

	return strconv.FormatFloat(res, 'f', -1, 64)
}
