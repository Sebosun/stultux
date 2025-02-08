package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebosun/stultux/internal/database"
)

type CreateUserPayload struct {
	Name        string `json:"name" example:"John"`
	LastName    string `json:"lastName" example:"Doe"`
	Password    string `json:"password" example:"somepassword"`
	CountryCode string `json:"countryCode" example:"DE"`
}

type User struct {
	Name        string
	LastName    string
	Password    string
	CountryCode string
}

func getNullString(item string) sql.NullString {
	return sql.NullString{
		Valid:  true,
		String: item,
	}
}

func (cfg *ApiConfig) CreateUser(c echo.Context) error {
	u := new(CreateUserPayload)
	err := c.Bind(u)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user := database.CreateUserParams{
		Name:        getNullString(u.Name),
		LastName:    getNullString(u.LastName),
		Password:    getNullString(u.Password),
		CountryCode: getNullString(u.CountryCode),
	}

	ok, err := cfg.DB.CreateUser(c.Request().Context(), user)

	if err != nil {
        fmt.Printf("Error %s \n", err)
		return c.JSON(http.StatusBadRequest, "User either exists or some items are already taken")
	}

	return c.JSON(http.StatusOK, ok)
}
