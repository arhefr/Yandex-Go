package model

import (
	"strconv"

	parser "github.com/arhefr/MathParser"
	Err "github.com/arhefr/Yandex-Go/orch/internal/errors"
)

const (
	StatusWait = "waiting"
	StatusCalc = "calculating"
	StatusErr  = "error"
	StatusDone = "done"
)

type Request struct {
	ID       string   `json:"id"`
	Status   string   `json:"status"`
	Result   string   `json:"result"`
	PostNote []Entity `json:"-"`
}

type Task struct {
	ID     string  `json:"id"`
	Sub_ID int     `json:"sub_id"`
	Arg1   float64 `json:"arg1"`
	Arg2   float64 `json:"arg2"`
	Oper   string  `json:"operation"`
}

type Response struct {
	ID     string `json:"id"`
	Sub_ID int    `json:"sub_id"`
	Result string `json:"result"`
}

type Expression struct {
	Expr string `json:"expression"`
}

type Entity struct {
	Name  string
	Index int
}

func NewExpr(id string, request *Expression) Request {
	postNote := parser.InfixPostfix(request.Expr)

	var postNoteEnt []Entity
	var numsCnt, opsCnt int
	for i := range postNote {
		postNoteEnt = append(postNoteEnt, Entity{Name: postNote[i], Index: i})

		if _, ok := parser.Token[postNote[i]]; ok {
			opsCnt++
		} else {
			numsCnt++
		}
	}

	if numsCnt-1 != opsCnt {
		return Request{ID: id, Status: StatusErr, Result: Err.IncorrectExpr.Error()}
	} else if numsCnt == 1 {
		return Request{ID: id, Status: StatusDone, Result: postNote[0]}
	}

	req := Request{ID: id, Status: StatusWait, PostNote: postNoteEnt}
	return req
}

func NewTask(id string, sub_id int, arg1, arg2 string, oper string) (Task, error) {
	n1, err1 := strconv.ParseFloat(arg1, 64)
	n2, err2 := strconv.ParseFloat(arg2, 64)
	if err1 != nil || err2 != nil {
		return Task{}, Err.IncorrectExpr
	}

	return Task{ID: id, Sub_ID: sub_id, Arg1: n1, Arg2: n2, Oper: oper}, nil
}

func (r Request) GetTask() (Task, error) {
	for i := range r.PostNote {
		if parser.IsOp(r.PostNote[i].Name) && i >= 2 {

			arg1, arg2, oper := r.PostNote[i-2].Name, r.PostNote[i-1].Name, r.PostNote[i]
			if arg2 == "0" && oper.Name == "/" {
				return Task{}, Err.DivisionByZero
			}

			if parser.IsNum(arg1) && parser.IsNum(arg2) {
				return NewTask(r.ID, oper.Index, arg1, arg2, oper.Name)
			}
		}
	}

	return Task{}, Err.NotFoundTask
}

func GetIndex(postNote []Entity, index int) int {
	for i, ent := range postNote {
		if ent.Index == index {
			return i
		}
	}

	return -1
}
