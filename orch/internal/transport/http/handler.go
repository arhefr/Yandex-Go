package router

import (
	"fmt"
	"net/http"
	"slices"
	"sync"

	Err "github.com/arhefr/Yandex-Go/orch/internal/errors"
	"github.com/arhefr/Yandex-Go/orch/internal/model"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type Service interface {
	Get() []model.Request
	GetByID(id string) (model.Request, bool)
	Add(expr model.Expression) (string, error)
	Replace(id string, req model.Request)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h *Handler) AddExpr(ctx echo.Context) error {

	req := new(model.Expression)
	if err := ctx.Bind(&req); err != nil {
		log.Warn(Err.IncorrectJSON)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	if id, err := h.service.Add(*req); err != nil {
		log.Warn(Err.Common)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.Common)
	} else {
		return ctx.JSON(http.StatusOK, struct {
			ID string `json:"id"`
		}{id})
	}
}

func (h *Handler) GetIDs(ctx echo.Context) error {
	exprs := h.service.Get()
	return ctx.JSON(http.StatusOK,
		struct {
			Exprs []model.Request `json:"expressions"`
		}{exprs})
}

func (h *Handler) GetID(ctx echo.Context) error {
	expr, exists := h.service.GetByID(ctx.Param("id"))
	if !exists {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectID)
	}

	return ctx.JSON(http.StatusOK, struct {
		Expr model.Request `json:"expression"`
	}{expr})
}

func (h *Handler) CatchTask(ctx echo.Context) error {
	var task model.Response

	if err := ctx.Bind(&task); err != nil {
		log.Warn(Err.IncorrectJSON)
		return echo.NewHTTPError(http.StatusUnprocessableEntity, Err.IncorrectJSON)
	}

	req, _ := h.service.GetByID(task.ID)
	i := model.GetIndex(req.PostNote, task.Sub_ID)
	if i == -1 {
		return nil
	}
	postNote := slices.Replace(req.PostNote, i, i+1, model.Entity{Name: task.Result})

	if len(postNote) == 1 {
		req.Result = postNote[0].Name
		req.Status = model.StatusDone
	}

	h.service.Replace(req.ID, req)
	return nil
}

func (h *Handler) SendTask(ctx echo.Context) error {
	var mu sync.RWMutex

	mu.Lock()
	defer mu.Unlock()
	for _, req := range h.service.Get() {
		if task, err := req.GetTask(); (req.Status == model.StatusWait || req.Status == model.StatusCalc) && err == nil {
			i := model.GetIndex(req.PostNote, task.Sub_ID)
			h.service.Replace(req.ID, model.Request{
				ID:       req.ID,
				Status:   model.StatusCalc,
				Result:   req.Result,
				PostNote: append(req.PostNote[:i-2], req.PostNote[i:]...),
			})

			return ctx.JSON(http.StatusOK, task)

		} else if err != nil && err != Err.NotFoundTask {
			h.service.Replace(req.ID, model.Request{
				ID:     req.ID,
				Status: model.StatusErr,
				Result: fmt.Sprint(err),
			})

			return echo.NewHTTPError(http.StatusInternalServerError, Err.IncorrectExpr)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, Err.NotFoundTask)
}
