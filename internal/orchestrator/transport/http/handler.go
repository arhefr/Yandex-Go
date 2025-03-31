package router

import (
	"fmt"
	"net/http"
	"slices"
	"sync"

	"github.com/arhefr/Yandex-Go/internal/orchestrator/model"
	repo "github.com/arhefr/Yandex-Go/internal/repository"
	Err "github.com/arhefr/Yandex-Go/pkg/errors"
	"github.com/arhefr/Yandex-Go/pkg/tools"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func AddExpr(ctx echo.Context) error {
	req := new(model.Expression)

	if err := ctx.Bind(&req); err != nil {
		log.Warn(Err.IncorrectJSON)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	expr := model.NewExpr(tools.NewCryptoRand(), req)
	repo.Exprs.Add(expr.ID, expr)
	return ctx.JSON(http.StatusOK, struct {
		ID string `json:"id"`
	}{expr.ID})
}

func GetIDs(ctx echo.Context) error {
	exprs := repo.Exprs.GetValues()
	return ctx.JSON(http.StatusOK,
		struct {
			Exprs []model.Request `json:"expressions"`
		}{exprs})
}

func GetID(ctx echo.Context) error {
	expr, exists := repo.Exprs.Get(ctx.Param("id"))
	if !exists {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectID)
	}

	return ctx.JSON(http.StatusOK, struct {
		Expr model.Request `json:"expression"`
	}{expr})
}

func CatchTask(ctx echo.Context) error {
	var task model.Response

	if err := ctx.Bind(&task); err != nil {
		log.Warn(Err.IncorrectJSON)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	expr, _ := repo.Exprs.Get(task.ID)
	i := model.GetIndex(expr.PostNote, task.Sub_ID)
	if i == -1 {
		return nil
	}
	postNote := slices.Replace(expr.PostNote, i, i+1, model.Entity{Name: task.Result})

	if len(postNote) == 1 {
		expr.Result = postNote[0].Name
		expr.Status = model.StatusDone
	}

	repo.Exprs.Add(expr.ID, expr)
	return nil
}

func SendTask(ctx echo.Context) error {
	var mu sync.RWMutex

	mu.Lock()
	defer mu.Unlock()
	for _, expr := range repo.Exprs.GetValues() {
		if task, err := expr.GetTask(); (expr.Status == model.StatusWait || expr.Status == model.StatusCalc) && err == nil {
			i := model.GetIndex(expr.PostNote, task.Sub_ID)
			repo.Exprs.Add(expr.ID, model.Request{
				ID:       expr.ID,
				Status:   model.StatusCalc,
				Result:   expr.Result,
				PostNote: append(expr.PostNote[:i-2], expr.PostNote[i:]...),
			})

			return ctx.JSON(http.StatusOK, task)

		} else if err != nil && err != Err.NotFoundTask {
			repo.Exprs.Add(expr.ID, model.Request{
				ID:     expr.ID,
				Status: model.StatusErr,
				Result: fmt.Sprint(err),
			})

			return echo.NewHTTPError(http.StatusInternalServerError, Err.IncorrectExpr)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, Err.NotFoundTask)
}
