package handlers

import (
	"fmt"
	"github.com/censync/go-dto"
	"github.com/censync/go-validator"
	"github.com/labstack/echo/v4"
	. "github.com/lidofinance/dc4bc/client/api/dto"
	cs "github.com/lidofinance/dc4bc/client/api/http_api/context_service"
	req "github.com/lidofinance/dc4bc/client/api/http_api/requests"
	"github.com/lidofinance/dc4bc/client/services"
	"net/http"
)

func GetOperations(c echo.Context) error {
	stx := c.(*cs.ContextService)

	operations, err := services.App().BaseClientService().GetOperations()

	if err == nil {
		return stx.Json(
			http.StatusOK,
			operations,
		)
	} else {
		return stx.JsonError(
			http.StatusInternalServerError,
			err,
		)
	}
}

func ProcessOperation(c echo.Context) error {
	stx := c.(*cs.ContextService)

	request := &req.OperationForm{}

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

	formDTO := &OperationDTO{}

	err = dto.RequestToDTO(formDTO, request)

	if err != nil {
		return stx.JsonError(
			http.StatusBadRequest,
			err,
		)
	}

	err = services.App().BaseClientService().ProcessOperation(formDTO)

	if err == nil {
		return stx.Json(
			http.StatusOK,
			"ok",
		)
	} else {
		return stx.JsonError(
			http.StatusInternalServerError,
			err,
		)
	}
}

func GetOperation(c echo.Context) error {
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

	operation, err := services.App().BaseClientService().GetOperation(formDTO)

	if err == nil {
		return stx.Json(
			http.StatusOK,
			operation,
		)
	} else {
		return stx.JsonError(
			http.StatusInternalServerError,
			fmt.Errorf("failed to get operations: %v", err),
		)
	}
}

func ApproveParticipation(c echo.Context) error {
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

	err = services.App().BaseClientService().ApproveParticipation(formDTO)

	if err == nil {
		return stx.Json(
			http.StatusOK,
			"ok",
		)
	} else {
		return stx.JsonError(
			http.StatusInternalServerError,
			err,
		)
	}
}
