package parser

import (
	"slices"
	"sort"
)

type Operator struct {
	Index    int
	Name     string
	Priority int
}

func NewOperator(name string, index, priority int) Operator {
	return Operator{index, name, priority}
}

func SortByPriority(ops []Operator) {
	sort.Slice(ops, func(i, j int) bool {
		return ops[i].Priority > ops[j].Priority
	})
}

func (op Operator) GetOperands(nums []float64) (float64, float64) {
	return nums[op.Index], nums[op.Index+1]
}

func (op Operator) Replace(nums []float64, ops []Operator, res float64) ([]float64, []Operator) {
	nums = slices.Insert(append(nums[:op.Index], nums[op.Index+2:]...), op.Index, res)

	for I, opNext := range ops[1:] {
		if opNext.Index > op.Index {
			ops[I+1].Index--
		}
	}

	ops = ops[1:]
	return nums, ops
}
