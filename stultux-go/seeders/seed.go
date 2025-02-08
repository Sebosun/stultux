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

	seedCountries(dbQueries)
	seedGermanNames(dbQueries)
	seedGermanLast(dbQueries)
    seedPasswords(dbQueries)
}

func seedGermanNames(db *database.Queries) {
	fmt.Println("=== Seeding German First Names ===")
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
			fmt.Println("Errr string already in db", v.Name)
		}
	}

}

func seedGermanLast(db *database.Queries) {
	fmt.Println("=== Seeding German Last Names ===")
	ctx := context.Background()
	names := []struct {
		Name        string
		CountryCode sql.NullString
	}{
		{Name: "Müller", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Schmidt", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Schneider", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Fischer", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Weber", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Meyer", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Wagner", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Becker", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Schulz", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Hoffmann", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Koch", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Bauer", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Richter", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Klein", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Wolf", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Schröder", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Neumann", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Braun", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Zimmermann", CountryCode: sql.NullString{String: "DE", Valid: true}},
		{Name: "Hartmann", CountryCode: sql.NullString{String: "DE", Valid: true}},
	}

	for _, v := range names {
		_, err := db.InsertLastName(ctx, v)
		if err != nil {
			fmt.Println("Errr string already in db", v.Name)
		}
	}
}

func seedCountries(db *database.Queries) {
	fmt.Println("=== Seeding Countries ===")
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
		_, err := db.InsertCountry(ctx, v)
		if err != nil {
			fmt.Println("Errr string already in db", v.CountryCode)
		}
	}

}

func seedPasswords(db *database.Queries) {
	fmt.Println("=== Seeding Passwords ===")
	ctx := context.Background()
	passwords := []string{
		"TaxFraudSpeedrun", 
        "BankSaysImJoking", 
        "OopsIDidWarCrime2", 
        "EpsteinLoggedIn", 
        "DarkWebDad96",
		"Bankai4BankFraudJP",
		"NarutoTaxEvasion420",
		"GokuScammedMeUS",
		"OnePiecePonziRU",
		"OverdraftNoJutsu69",
		"FBI_OpenUp_Tokyo",
		"CryptoUzumaki1337",
		"TaxFraudNoKaisen",
		"EvangelionBankruptFR",
		"SasukeGotMyPESEL",
		"FullmetalWelfareCheck",
		"PayPalChargebackIT",
		"YakuzaBitcoinRug",
		"DeathNoteForDebtsUK",
		"BleachMyBankBalanceDE",
		"HellsingOffshoreHK",
		"GundamFakeIDBR",
		"NeonGenesisPonzi",
		"ZelleMeUrRentSP",
        "ArmySurplusKatanaBR",
		"OverlordOfUnpaidLoans",
        "FavelaWiFiAteMyBitcoin",
        "ZUSOkrutnyJakDoflamingo",
    }


	for _, v := range passwords {
		_, err := db.InsertPassword(ctx, v)
		if err != nil {
			fmt.Println("Errr string already in db", v)
		}
	}
}
