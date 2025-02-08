package seeders

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/sebosun/stultux/internal/database"
)

func InitSeed() {
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	/* seedCountries(dbQueries) */
    seedGermanNames(dbQueries)
}

func seedGermanNames(db *database.Queries) {
	ctx := context.Background()
	namesMap := []database.InsertNameParams{
		{Name: "Emma", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Mia", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Hannah", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Lea", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Sophia", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Lina", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Marie", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Anna", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Johanna", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Luisa", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Paul", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Ben", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Noah", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Leon", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Finn", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Luis", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Elias", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Jonas", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Maximilian", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Felix", CountryCode: sql.NullString{String: "DE", Valid: true}}}

	for _, v := range namesMap {
		_, err := db.InsertName(ctx, v)
		if err != nil {
			fmt.Println("Errr string already in db", v.CountryCode.String)
		}
	}

}

func seedCountries(db *database.Queries) {
	//The fuck does that even do? TODO: Check it out
	ctx := context.Background()
	countriesMap := []database.InsertCountryParams{
		{CountryName: "Germany", CountryCode: "DE"},
		{CountryName: "Italy", CountryCode: "IT"},
		{CountryName: "Sweden", CountryCode: "SE"},
		{CountryName: "France", CountryCode: "FR"},
		{CountryName: "Spain", CountryCode: "ES"},
		{CountryName: "United States", CountryCode: "US"},
	}

	for _, v := range countriesMap {
		db.InsertCountry(ctx, v)
	}

}
