package hanlder

import (
	"errors"
	"net/http"

	"github.com/buemura/health-checker/internal/core/dto"
	"github.com/buemura/health-checker/internal/core/usecase"
	"github.com/buemura/health-checker/internal/infra/database"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, nil)
	})

	e.GET("/endpoints", getEndpoint)
	e.POST("/endpoints", createEndpoint)
}

func getEndpoint(c echo.Context) error {
	er := database.NewEndpointRepositoryImpl(database.DB)
	uc := usecase.NewGetEndpointList(er)
	res, err := uc.Execute()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func createEndpoint(c echo.Context) error {
	body := new(dto.CreateEndpointIn)
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}
	if err := validateCreateEndpoint(body); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
	}

	er := database.NewEndpointRepositoryImpl(database.DB)
	uc := usecase.NewCreateEndpoint(er)
	res, err := uc.Execute(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}

func validateCreateEndpoint(body *dto.CreateEndpointIn) error {
	if len(body.Name) < 1 {
		return errors.New("invalid Name")
	}
	if len(body.Url) < 1 {
		return errors.New("invalid Url")
	}
	if len(body.NotifyTo) < 1 {
		return errors.New("invalid NotifyTo")
	}
	if body.CheckFrequency < 1 {
		return errors.New("invalid CheckFrequency")
	}
	return nil
}
