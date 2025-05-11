package router

import (
	"context"
	"net/http"
	"slices"
	"sync"

	Err "github.com/arhefr/Yandex-Go/orch/internal/errors"
	"github.com/arhefr/Yandex-Go/orch/internal/model"
	repo "github.com/arhefr/Yandex-Go/orch/internal/repository"
	"github.com/labstack/echo/v4"
)

type Service interface {
	Postgresql
}

type Postgresql interface {
	Get(ctx context.Context) ([]model.Expression, error)
	GetByID(ctx context.Context, id string) (model.Expression, error)
	Add(ctx context.Context, expr model.Expression) error
	Replace(ctx context.Context, id, status, result string) error
}

type Handler struct {
	ts      *repo.SafeMap
	service Service
	mu      sync.RWMutex
}

func NewHandler(service Service, ts *repo.SafeMap) Handler {
	return Handler{service: service, mu: sync.RWMutex{}, ts: ts}
}

func (h *Handler) AddExpr(ctx echo.Context) (err error) {

	expr := model.NewExpression()
	if err := ctx.Bind(&expr); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	req := model.NewRequest(*expr)
	if req.Status == model.StatusErr {
		expr.Status = model.StatusErr
	} else {
		h.ts.Add(req, expr.ID)
	}
	if err := h.service.Add(context.TODO(), *expr); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.Common)
	} else {
		return ctx.JSON(http.StatusOK, struct {
			ID string `json:"id"`
		}{expr.ID})
	}
}

func (h *Handler) GetIDs(ctx echo.Context) (err error) {
	exprs, err := h.service.Get(context.TODO())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.Common)
	}

	return ctx.JSON(http.StatusOK,
		struct {
			Exprs []model.Expression `json:"expressions"`
		}{exprs})
}

func (h *Handler) GetID(ctx echo.Context) (err error) {
	expr, err := h.service.GetByID(context.TODO(), ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectID)
	}

	return ctx.JSON(http.StatusOK, struct {
		Expr model.Expression `json:"expression"`
	}{expr})
}

func (h *Handler) CatchTask(ctx echo.Context) error {
	var task model.Response

	if err := ctx.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	req, _ := h.ts.GetById(task.ID)
	i := model.GetIndex(req.PostNote, task.Sub_ID)
	if i == -1 {
		return nil
	}
	postNote := slices.Replace(req.PostNote, i, i+1, model.Entity{Name: task.Result})

	if len(postNote) == 1 {
		req.Status = model.StatusDone
		h.service.Replace(context.TODO(), req.ID, model.StatusDone, postNote[0].Name)
		h.ts.Delete(req.ID)
	}

	h.ts.Refresh(req.ID, *req)
	return nil
}

func (h *Handler) SendTask(ctx echo.Context) error {
	var mu sync.RWMutex

	mu.Lock()
	defer mu.Unlock()
	for _, req := range h.ts.Get() {
		if task, err := req.GetTask(); (req.Status == model.StatusWait) && err == nil {
			i := model.GetIndex(req.PostNote, task.Sub_ID)
			h.ts.Refresh(req.ID, model.Request{
				ID:       req.ID,
				Status:   model.StatusWait,
				PostNote: append(req.PostNote[:i-2], req.PostNote[i:]...),
			})

			return ctx.JSON(http.StatusOK, task)

		} else if err != nil && err != Err.NotFoundTask {
			h.service.Replace(context.TODO(), req.ID, model.StatusErr, "")
			h.ts.Delete(req.ID)

			return echo.NewHTTPError(http.StatusInternalServerError, Err.IncorrectExpr)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, Err.NotFoundTask)
}
