package handlers

import (
	"context"
	"net/http"
	"slices"

	Err "github.com/arhefr/Yandex-Go/orch/internal/errors"
	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignIn(ctx echo.Context) (err error) {
	return nil
}

func (h *Handler) LogIn(ctx echo.Context) (err error) {
	return nil
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
		h.services.AddReq(expr.ID, req)
	}
	if err := h.services.AddExpr(context.TODO(), *expr); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.Common)
	} else {
		return ctx.JSON(http.StatusOK, struct {
			ID string `json:"id"`
		}{expr.ID})
	}
}

func (h *Handler) GetIDs(ctx echo.Context) (err error) {
	exprs, err := h.services.GetExprs(context.TODO())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.Common)
	}

	return ctx.JSON(http.StatusOK,
		struct {
			Exprs []model.Expression `json:"expressions"`
		}{exprs})
}

func (h *Handler) GetID(ctx echo.Context) (err error) {
	expr, err := h.services.GetExprByID(context.TODO(), ctx.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectID)
	}

	return ctx.JSON(http.StatusOK, struct {
		Expr model.Expression `json:"expression"`
	}{expr})
}

func (h *Handler) CatchTask(ctx echo.Context) error {
	var task model.Response

	h.mu.RLock()
	defer h.mu.RUnlock()

	if err := ctx.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	req, _ := h.services.GetReqByID(task.ID)
	i, err := model.GetIndex(req.PostNote, task.Sub_ID)
	if err != nil {
		return nil
	}
	postNote := slices.Replace(req.PostNote, i, i+1, model.Entity{Name: task.Result})

	if len(postNote) == 1 {
		req.Status = model.StatusDone
		h.services.ReplaceExpr(context.TODO(), req.ID, model.StatusDone, postNote[0].Name)
	}

	h.services.ReplaceReq(req.ID, *req)
	return nil
}

func (h *Handler) SendTask(ctx echo.Context) error {

	h.mu.Lock()
	defer h.mu.Unlock()
	for _, req := range h.services.GetReq() {
		if task, err := req.GetTask(); err != nil && err != Err.NotFoundTask {

			h.services.ReplaceExpr(context.TODO(), req.ID, model.StatusErr, "")
			h.services.DeleteReq(req.ID)
			return echo.NewHTTPError(http.StatusInternalServerError, Err.IncorrectExpr)
		} else if req.Status == model.StatusDone {

			h.services.DeleteReq(req.ID)
		} else if req.Status == model.StatusWait {

			i, err := model.GetIndex(req.PostNote, task.Sub_ID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, Err.IncorrectExpr)
			}

			req.PostNote = append(req.PostNote[:i-2], req.PostNote[i:]...)
			h.services.ReplaceReq(req.ID, req)

			return ctx.JSON(http.StatusOK, task)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, Err.NotFoundTask)
}
