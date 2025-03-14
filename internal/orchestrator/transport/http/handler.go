package router

import (
	"fmt"
	"net/http"

	models_agent "github.com/arhefr/Yandex-Go/internal/agent/model"
	"github.com/arhefr/Yandex-Go/internal/orchestrator/model"
	repo "github.com/arhefr/Yandex-Go/internal/repository"
	Err "github.com/arhefr/Yandex-Go/pkg/errors"
	"github.com/arhefr/Yandex-Go/pkg/tools"

	"github.com/labstack/echo/v4"
)

func AddExpr(ctx echo.Context) error {
	request := new(model.Request)

	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	id := tools.NewCryptoRand()
	expr := model.NewExpression(id, request)
	repo.Tasks.Add(id, expr)
	return ctx.JSON(http.StatusOK, struct {
		ID string `json:"id"`
	}{id})
}

func GetIDs(ctx echo.Context) error {
	exprs := repo.Tasks.GetValues()
	return ctx.JSON(http.StatusOK,
		struct {
			Exprs []model.Expression `json:"expressions"`
		}{exprs})
}

func FetchTask(ctx echo.Context) error {
	req := new(models_agent.Response)

	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	expr, exists := repo.Tasks.Get(req.ID)
	if !exists {
		return echo.NewHTTPError(http.StatusNotFound, Err.IncorrectID)
	}

	op := expr.Parser.Ops[0]
	expr.Parser.Nums, expr.Parser.Ops = op.Replace(expr.Parser.Nums, expr.Parser.Ops, req.Result)

	if len(expr.Parser.Nums) == 1 {
		expr.Status = model.StatusDone
		expr.Result = fmt.Sprintf("%.3f", expr.Parser.Nums[0])
	} else {
		expr.Status = model.StatusWait
	}
	repo.Tasks.Add(req.ID, expr)

	return ctx.JSON(http.StatusOK, nil)
}

func GetTask(ctx echo.Context) error {

	for _, expr := range repo.Tasks.GetValues() {
		if expr.Status == model.StatusWait {
			expr.Status = model.StatusCalc
			repo.Tasks.Add(expr.ID, expr)
			return ctx.JSON(http.StatusOK, expr.GetTask())
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, Err.NotFoundTask)
}

func GetID(ctx echo.Context) error {
	id := ctx.Param("id")

	expr, exists := repo.Tasks.Get(id)
	if !exists {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectID)
	}

	return ctx.JSON(http.StatusOK, struct {
		Expr model.Expression `json:"expression"`
	}{expr})
}
