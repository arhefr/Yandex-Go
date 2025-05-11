package model

import (
	"strconv"

	parser "github.com/arhefr/MathParser"
	Err "github.com/arhefr/Yandex-Go/orch/internal/errors"
)

type Task struct {
	ID     string  `json:"id"`
	Sub_ID int     `json:"sub_id"`
	Arg1   float64 `json:"arg1"`
	Arg2   float64 `json:"arg2"`
	Oper   string  `json:"operation"`
}

func NewTask(id string, sub_id int, arg1, arg2 float64, oper string) Task {
	return Task{ID: id, Sub_ID: sub_id, Arg1: arg1, Arg2: arg2, Oper: oper}
}

func (r Request) GetTask() (Task, error) {
	for i := range r.PostNote {
		if parser.IsOp(r.PostNote[i].Name) && i >= 2 {

			arg1, arg2, oper := r.PostNote[i-2].Name, r.PostNote[i-1].Name, r.PostNote[i]
			if arg2 == "0" && oper.Name == "/" {
				return Task{}, Err.DivisionByZero
			}

			if parser.IsNum(arg1) && parser.IsNum(arg2) {
				n1, err1 := strconv.ParseFloat(arg1, 64)
				n2, err2 := strconv.ParseFloat(arg2, 64)
				if err1 != nil || err2 != nil {
					return Task{}, Err.IncorrectExpr
				}

				return NewTask(r.ID, oper.Index, n1, n2, oper.Name), nil
			}
		}
	}

	return Task{}, Err.NotFoundTask
}
