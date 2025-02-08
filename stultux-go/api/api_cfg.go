package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sebosun/stultux/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}


// @Summary Returns a simple greeting
// @Description This endpoint returns a simple hello message
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func (cfg *ApiConfig) HelloHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}

