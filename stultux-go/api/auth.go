package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegistrationDetails struct {
    FirstNames []string `json:"firstNames"`
	LastNames []string `json:"lastNames"`
	Passwords []string `json:"passwords"`
}

func (cfg *ApiConfig) GetRegistrationDetails(c echo.Context) error {
	names, err := cfg.DB.GetNames(c.Request().Context())

	if err != nil {
		fmt.Printf("Error %s \n", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	last_names, err := cfg.DB.GetLastNames(c.Request().Context())

	if err != nil {
		fmt.Printf("Error %s \n", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

	passwords, err := cfg.DB.GetPasswords(c.Request().Context())

	if err != nil {
		fmt.Printf("Error %s \n", err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong")
	}

    resp := RegistrationDetails{
        FirstNames: names,
        LastNames: last_names,
        Passwords: passwords,
    }


	return c.JSON(http.StatusOK, resp)
}
