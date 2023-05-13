package handler

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type ResultHandler struct {
}

type SaveGameResultReq struct {
	gameUuid string `param:"uuid" validate:"required"`
}

func (r *SaveGameResultReq) Validate() error {
	return validator.New().Struct(r)
}

func (h *ResultHandler) SaveGameResult(c echo.Context) error {
	// Get the request
	var resReq SaveGameResultReq
	if err := c.Bind(&resReq); err != nil {
		log.Errorln(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request: %v", err))
	}

	// Validate the request
	if err := resReq.Validate(); err != nil {
		log.Errorln(err)
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid request: %v", err))
	}

	return c.JSON(http.StatusOK, "GetResult")
}

func NewResultHandler() *ResultHandler {
	return &ResultHandler{}
}
