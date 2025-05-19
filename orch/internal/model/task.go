package model

import (
	"fmt"
	"strconv"

	parser "github.com/arhefr/MathParser"
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

func (r Request) GetTask() (Task, bool, error) {
	for i := range r.PostNote {
		if parser.IsOp(r.PostNote[i].Name) && i >= 2 {

			arg1, arg2, oper := r.PostNote[i-2].Name, r.PostNote[i-1].Name, r.PostNote[i]
			if arg2 == "0" && oper.Name == "/" {
				return Task{}, false, fmt.Errorf("error division by zero")
			}

			if parser.IsNum(arg1) && parser.IsNum(arg2) {
				n1, err1 := strconv.ParseFloat(arg1, 64)
				n2, err2 := strconv.ParseFloat(arg2, 64)
				if err1 != nil || err2 != nil {
					return Task{}, false, fmt.Errorf("error incorrect expression")
				}

				return NewTask(r.ID, oper.Index, n1, n2, oper.Name), true, nil
			}
		}
	}

	return Task{}, false, nil
}
