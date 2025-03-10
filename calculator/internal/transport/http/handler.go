package router

import (
	models_agent "calculator/internal/agent/models"
	"calculator/internal/orchestrator/models"
	repo "calculator/internal/repository"
	"calculator/pkg/tools"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	var request models.Request

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, `{"error":"cannot reading body"}`, http.StatusUnprocessableEntity)
		return
	}

	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, `{"error":"incorrect JSON"}`, http.StatusUnprocessableEntity)
		return
	}

	id := tools.NewCryptoRand()
	expr := models.NewExpression(id, request)
	repo.Tasks.Add(id, expr)
	fmt.Fprintf(w, `{"id": %d}`, id)
}

func GetIDs(w http.ResponseWriter, r *http.Request) {
	exprs := repo.Tasks.GetValues()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		struct {
			Exprs []models.Expression `json:"expressions"`
		}{exprs})
}

func GetID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	expr, exists := repo.Tasks.Get(id)
	if !exists {
		http.Error(w, `{"error":"not found task"}`, http.StatusUnprocessableEntity)
		return
	}

	json.NewEncoder(w).Encode(struct {
		Expr models.Expression `json:"expression"`
	}{expr})
}

func GetOperation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		for _, expr := range repo.Tasks.GetValues() {
			if expr.Status == "processing" {
				expr.Status = "calculating"
				repo.Tasks.Add(expr.ID, expr)
				json.NewEncoder(w).Encode(expr.GetTask())
				return
			}
		}
		http.Error(w, "", http.StatusNotFound)

	case http.MethodPost:
		var resp models_agent.Response

		body, _ := io.ReadAll(r.Body)
		json.Unmarshal(body, &resp)

		expr, _ := repo.Tasks.Get(resp.ID)
		op := expr.Parser.Ops[0]
		expr.Parser.Nums, expr.Parser.Ops = op.Replace(expr.Parser.Nums, expr.Parser.Ops, resp.Result)

		if len(expr.Parser.Nums) == 1 {
			expr.Status = "done"
			expr.Result = fmt.Sprintf("%.3f", expr.Parser.Nums[0])
		} else {
			expr.Status = "processing"
		}
		repo.Tasks.Add(resp.ID, expr)
	}
}
