package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	_ "github.com/sebosun/stultux/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/sebosun/stultux/seeders"

	"database/sql"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware

	"github.com/sebosun/stultux/api"
	_ "github.com/sebosun/stultux/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/sebosun/stultux/internal/database"
)

// @title Stultex
// @Summary Api for Stultex
// @version 1.0
// @description API for Stultex
// @host localhost:1323
// @BasePath /api/v1
func main() {
	godotenv.Load(".env")

	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		if args[0] == "seed" {
			fmt.Println("init seed")
			seeders.InitSeed()
		}
		return
	}

	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(" error opening connection to the database: ", err)
	}

	dbQueries := database.New(db)

	apiConfig := api.ApiConfig{
		DB: dbQueries,
	}

	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	g := e.Group("/api/v1")

	g.GET("/hello", apiConfig.HelloHandler)

	g.POST("/user", apiConfig.CreateUser)

	g.GET("/registration", apiConfig.GetRegistrationDetails)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	portFormatted := fmt.Sprintf(":%s", port)
	fmt.Println("Swagger server running on http://localhost:1323/swagger/index.html")
	e.Logger.Fatal(e.Start(portFormatted))
}
