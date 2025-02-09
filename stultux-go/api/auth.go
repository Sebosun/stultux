package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (cfg *ApiConfig) GetRegistrationDetails(c echo.Context) error {
	result, err := cfg.DB.GetRegistrationDetails(c.Request().Context())

	if err != nil {
		fmt.Printf("Error %s \n", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	return c.JSON(http.StatusOK, result)
}
