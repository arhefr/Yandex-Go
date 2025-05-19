package service

import (
	"testing"

	"github.com/arhefr/Yandex-Go/agent/internal/model"
)

type test struct {
	in  model.Task
	out string
}

var tests = []test{
	{
		in:  model.Task{Arg1: 2, Arg2: 2, Oper: "+"},
		out: "4",
	},
	{
		in:  model.Task{Arg1: 2, Arg2: 0, Oper: "/"},
		out: "",
	},
	{
		in:  model.Task{Arg1: 2, Arg2: 2, Oper: "/"},
		out: "1",
	},
	{
		in:  model.Task{Arg1: 2, Arg2: 2, Oper: "*"},
		out: "4",
	},
	{
		in:  model.Task{Arg1: 2, Arg2: 2, Oper: "-"},
		out: "0",
	},
}

func TestMakeTask(t *testing.T) {
	for _, test := range tests {
		res := MakeTask(&test.in)
		if res != test.out {
			t.Fatalf("TestMakeTask: get %s, expected: %s", res, test.out)
		}
	}
}
