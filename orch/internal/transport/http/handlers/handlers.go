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
	user := model.NewUser()
	if err := ctx.Bind(&user); err != nil {
		return SendJSON(ctx, ResponseWrongJSON)
	}

	if !h.services.ServiceUsers.CheckArgs(user) {
		return SendJSON(ctx, ResponseBadAuth)
	}

	if exists, err := h.services.ServiceUsers.Exists(context.TODO(), user); err != nil {
		return SendJSON(ctx, ResponseInternalError)
	} else if exists {
		return SendJSON(ctx, ResponseLoginExists)
	}

	if err := h.services.ServiceUsers.SignIn(context.TODO(), user); err != nil {
		return SendJSON(ctx, ResponseInternalError)
	}

	return SendJSON(ctx, ReponseOK)
}

func (h *Handler) LogIn(ctx echo.Context) (err error) {
	user := new(model.User)
	if err := ctx.Bind(&user); err != nil {
		return SendJSON(ctx, ResponseWrongJSON)
	}

	if exists, err := h.services.ServiceUsers.Exists(context.TODO(), user); err != nil {
		return SendJSON(ctx, ResponseInternalError)
	} else if !exists {
		return SendJSON(ctx, ResponseWrongLogin)
	}

	id, err := h.services.ServiceUsers.GetUserID(context.TODO(), user)
	if err != nil {
		return SendJSON(ctx, ResponseWrongPassword)
	}

	token, err := h.services.ServiceUsers.GetJWT(id)
	if err != nil {
		return SendJSON(ctx, ResponseInternalError)
	}

	return SendJSON(ctx, struct {
		Token string `json:"token"`
	}{token})
}

func (h *Handler) AddExpr(ctx echo.Context) (err error) {
	auth := ctx.Request().Header.Get("Authorization")
	if len(auth) < len("Bearer ") {
		return SendJSON(ctx, ResponseRequiredAuth)
	}
	user, err := h.services.ParseJWT(auth[len("Bearer "):])
	if err != nil {
		return SendJSON(ctx, ResponseWrongJWT)
	}

	expr := model.NewExpression(user.ID)
	if err := ctx.Bind(&expr); err != nil {
		return SendJSON(ctx, ResponseWrongJSON)
	}

	req := model.NewRequest(*expr)
	if req.Status == model.ExprStatusErr {
		expr.Status = model.ExprStatusErr
	} else {
		h.services.AddReq(expr.ID, req)
	}
	if err := h.services.AddExpr(context.TODO(), *expr); err != nil {
		return SendJSON(ctx, ResponseInternalError)
	} else {
		return SendJSON(ctx, struct {
			ID string `json:"id"`
		}{expr.ID})
	}
}

func (h *Handler) GetIDs(ctx echo.Context) (err error) {
	auth := ctx.Request().Header.Get("Authorization")
	if len(auth) < len("Bearer ") {
		return SendJSON(ctx, ResponseRequiredAuth)
	}
	user, err := h.services.ParseJWT(auth[len("Bearer "):])
	if err != nil {
		return SendJSON(ctx, ResponseWrongJWT)
	}

	exprs, err := h.services.GetExprs(context.TODO(), user.ID)
	if err != nil {
		return SendJSON(ctx, ResponseInternalError)
	}

	return SendJSON(ctx, struct {
		Exprs []model.Expression `json:"expressions"`
	}{exprs})
}

func (h *Handler) GetID(ctx echo.Context) (err error) {
	auth := ctx.Request().Header.Get("Authorization")
	if len(auth) < len("Bearer ") {
		return SendJSON(ctx, ResponseRequiredAuth)
	}
	user, err := h.services.ParseJWT(auth[len("Bearer "):])
	if err != nil {
		return SendJSON(ctx, ResponseWrongJWT)
	}

	expr, err := h.services.GetExprByID(context.TODO(), user.ID, ctx.Param("id"))
	if err != nil {
		return SendJSON(ctx, ResponseWrongID)
	}

	return SendJSON(ctx, struct {
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
		req.Status = model.ExprStatusDone
		h.services.ReplaceExpr(context.TODO(), req.ID, model.ExprStatusDone, postNote[0].Name)
	}

	h.services.ReplaceReq(req.ID, *req)
	return nil
}

func (h *Handler) SendTask(ctx echo.Context) error {

	h.mu.Lock()
	defer h.mu.Unlock()
	for _, req := range h.services.GetReq() {
		if task, err := req.GetTask(); err != nil && err != Err.NotFoundTask {

			h.services.ReplaceExpr(context.TODO(), req.ID, model.ExprStatusDone, "")
			h.services.DeleteReq(req.ID)
			return echo.NewHTTPError(http.StatusInternalServerError, Err.IncorrectExpr)
		} else if req.Status == model.ExprStatusDone {

			h.services.DeleteReq(req.ID)
		} else if req.Status == model.ExprStatusWait {

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
