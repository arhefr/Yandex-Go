package handler

import (
	"calculator/internal/agent/model"
	"calculator/internal/orchestrator/transport/http/models"
	repo "calculator/internal/repository"
	"calculator/pkg/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var (
	tasks = repo.Tasks
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error":"Not valid JSON"}`, http.StatusUnprocessableEntity)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, `{"error":"Not valid JSON"}`, http.StatusUnprocessableEntity)
		return
	}

	id := tools.NewCryptoRand()
	expr := models.NewExpression(id, request)
	tasks.Add(id, expr)

	fmt.Fprintf(w, `{"id": %d}`, id)
}

func GetIDs(w http.ResponseWriter, r *http.Request) {
	tasks := tasks.GetValues()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	if id, err := strconv.Atoi(r.RequestURI[len("/api/v1/expressions/:"):]); err == nil {
		task, err := tasks.Get(id)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error":"%s"}`, err), 404)
			return
		}
		json.NewEncoder(w).Encode(task)
		return
	}

	http.Error(w, `{"error":"id must be int"}`, 500)
}

func GetOperation(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		for _, expr := range tasks.GetValues() {
			if expr.Status == "in process" {
				task := expr.GetTask()
				expr.Status = "calculating..."
				tasks.Delete(expr.ID)
				tasks.Add(expr.ID, expr)
				json.NewEncoder(w).Encode(task)
				return
			}
		}

		http.Error(w, ``, 404)
	} else if r.Method == http.MethodPost {
		var resp model.Response
		body, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(body, &resp); err != nil {
			http.Error(w, ``, 500)
		}

		expr, _ := tasks.Get(resp.ID)
		op := expr.Parser.Ops[0]
		expr.Parser.Nums, expr.Parser.Ops = op.Replace(expr.Parser.Nums, expr.Parser.Ops, resp.Result)
		if len(expr.Parser.Nums) == 1 {
			expr.Status = "complete"
			expr.Result = fmt.Sprintf("%.3f", expr.Parser.Nums[0])
		} else {
			expr.Status = "in process"
		}
		tasks.Delete(resp.ID)
		tasks.Add(resp.ID, expr)
	}
}
