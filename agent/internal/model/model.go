package model

type Response struct {
	ID     string `json:"id"`
	Sub_ID int    `json:"sub_id"`
	Result string `json:"result"`
}

type Task struct {
	ID     string  `json:"id"`
	Sub_ID int     `json:"sub_id"`
	Arg1   float64 `json:"arg1"`
	Arg2   float64 `json:"arg2"`
	Oper   string  `json:"operation"`
}
