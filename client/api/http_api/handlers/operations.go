package handlers

import (
	"fmt"
	"net/http"

	"github.com/censync/go-dto"
	"github.com/censync/go-validator"
	"github.com/labstack/echo/v4"
	. "github.com/lidofinance/dc4bc/client/api/dto"
	cs "github.com/lidofinance/dc4bc/client/api/http_api/context_service"
	req "github.com/lidofinance/dc4bc/client/api/http_api/requests"
)

func (a *HTTPApp) GetOperations(c echo.Context) error {
	stx := c.(*cs.ContextService)
	operations, err := a.operation.GetOperations()
	if err != nil {
		return stx.JsonError(http.StatusInternalServerError, err)
	}
	return stx.Json(http.StatusOK, operations)
}

func (a *HTTPApp) ProcessOperation(c echo.Context) error {
	stx := c.(*cs.ContextService)
	formDTO := &OperationDTO{}
	if err := stx.BindToDTO(&req.OperationForm{}, formDTO); err != nil {
		return err
	}

	if err := a.node.ProcessOperation(formDTO); err != nil {
		return stx.JsonError(http.StatusInternalServerError, err)
	}
	return stx.Json(http.StatusOK, "ok")
}

func (a *HTTPApp) GetOperation(c echo.Context) error {
	stx := c.(*cs.ContextService)

	request := &req.OperationIdForm{}

	err := stx.Bind(request)
	if err != nil {
		return stx.JsonError(
			http.StatusBadRequest,
			fmt.Errorf("failed to read request body: %v", err),
		)
	}

	if err := validator.Validate(request); !err.IsEmpty() {
		return stx.JsonError(
			http.StatusBadRequest,
			err.Error(),
		)
	}

	formDTO := &OperationIdDTO{}

	err = dto.RequestToDTO(formDTO, request)
	if err != nil {
		return stx.JsonError(
			http.StatusBadRequest,
			err,
		)
	}

	operation, err := a.operation.GetOperationByID(formDTO.OperationID)
	if err != nil {
		return stx.JsonError(
			http.StatusInternalServerError,
			fmt.Errorf("failed to get operations: %v", err),
		)
	}

	return stx.Json(
		http.StatusOK,
		operation,
	)
}

func (a *HTTPApp) ApproveParticipation(c echo.Context) error {
	stx := c.(*cs.ContextService)
	formDTO := &OperationIdDTO{}
	if err := stx.BindToDTO(&req.OperationIdForm{}, formDTO); err != nil {
		return err
	}

	if err := a.node.ApproveParticipation(formDTO); err != nil {
		return stx.JsonError(http.StatusInternalServerError, err)
	}
	return stx.Json(http.StatusOK, "ok")
}
