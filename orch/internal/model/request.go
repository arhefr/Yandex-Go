package model

import (
	"fmt"

	parser "github.com/arhefr/MathParser"
)

type Request struct {
	ID       string   `json:"id"`
	Status   string   `json:"status"`
	PostNote []Entity `json:"postnote"`
}

type Entity struct {
	Name  string
	Index int
}

func NewRequest(expr Expression) (req Request) {
	postNote := parser.InfixPostfix(expr.Expr)

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

	if numsCnt-1 != opsCnt || numsCnt == 1 {
		return Request{ID: expr.ID, Status: ExprStatusErr}
	}

	req = Request{ID: expr.ID, Status: ExprStatusWait, PostNote: postNoteEnt}
	return req
}

func GetIndex(postNote []Entity, index int) (int, error) {
	for i, ent := range postNote {
		if ent.Index == index {
			return i, nil
		}
	}

	return -1, fmt.Errorf("model: GetIndex: error missing index")
}
