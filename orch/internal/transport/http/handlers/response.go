package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	StatusCode int    `json:"status"`
	Message    string `json:"message"`
}

func NewResponse(statusCode int, message string) Response {
	return Response{StatusCode: statusCode, Message: message}
}

func SendJSON(ctx echo.Context, json interface{}) error {
	return ctx.JSON(200, json)
}

var (
	ReponseOK = NewResponse(http.StatusOK, "ok")

	ResponseWrongJSON     = NewResponse(http.StatusUnprocessableEntity, "error wrong JSON")
	ResponseInternalError = NewResponse(http.StatusInternalServerError, "error something went wrong")

	ResponseRequiredAuth = NewResponse(http.StatusUnauthorized, "error authentication")
	ResponseWrongJWT     = NewResponse(http.StatusUnprocessableEntity, "error wrong JWT token")
	ResponseBadAuth      = NewResponse(http.StatusUnprocessableEntity, "error password must contain 8 characters or more and login must contain 3 characters or more")
	ResponseLoginExists  = NewResponse(http.StatusUnprocessableEntity, "error login already exists")

	ResponseWrongLogin    = NewResponse(http.StatusUnprocessableEntity, "error login not exists")
	ResponseWrongPassword = NewResponse(http.StatusUnprocessableEntity, "error wrong password")

	ResponseWrongID = NewResponse(http.StatusUnprocessableEntity, "error wrong uuid")
)
