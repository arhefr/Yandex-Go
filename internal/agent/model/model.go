package model

import (
	"time"
)

type Response struct {
	ID     string `json:"id"`
	Sub_ID int    `json:"sub_id"`
	Result string `json:"result"`
}

type OperationTime struct {
	Add time.Duration
	Sub time.Duration
	Mul time.Duration
	Div time.Duration
}
