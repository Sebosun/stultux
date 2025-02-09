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
		{Name: "Emma", CountryCode: "DE"},
		{Name: "Mia", CountryCode: "DE"},
		{Name: "Hannah", CountryCode: "DE"},
		{Name: "Lea", CountryCode: "DE"},
		{Name: "Sophia", CountryCode: "DE"},
		{Name: "Lina", CountryCode: "DE"},
		{Name: "Marie", CountryCode: "DE"},
		{Name: "Anna", CountryCode: "DE"},
		{Name: "Johanna", CountryCode: "DE"},
		{Name: "Luisa", CountryCode: "DE"},
		{Name: "Paul", CountryCode: "DE"},
		{Name: "Ben", CountryCode: "DE"},
		{Name: "Noah", CountryCode: "DE"},
		{Name: "Leon", CountryCode: "DE"},
		{Name: "Finn", CountryCode: "DE"},
		{Name: "Luis", CountryCode: "DE"},
		{Name: "Elias", CountryCode: "DE"},
		{Name: "Jonas", CountryCode: "DE"},
		{Name: "Maximilian", CountryCode: "DE"},
		{Name: "Felix", CountryCode: "DE"}}

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
		CountryCode string
	}{
		{Name: "Müller", CountryCode: "DE"},
		{Name: "Schmidt", CountryCode: "DE"},
		{Name: "Schneider", CountryCode: "DE"},
		{Name: "Fischer", CountryCode: "DE"},
		{Name: "Weber", CountryCode: "DE"},
		{Name: "Meyer", CountryCode: "DE"},
		{Name: "Wagner", CountryCode: "DE"},
		{Name: "Becker", CountryCode: "DE"},
		{Name: "Schulz", CountryCode: "DE"},
		{Name: "Hoffmann", CountryCode: "DE"},
		{Name: "Koch", CountryCode: "DE"},
		{Name: "Bauer", CountryCode: "DE"},
		{Name: "Richter", CountryCode: "DE"},
		{Name: "Klein", CountryCode: "DE"},
		{Name: "Wolf", CountryCode: "DE"},
		{Name: "Schröder", CountryCode: "DE"},
		{Name: "Neumann", CountryCode: "DE"},
		{Name: "Braun", CountryCode: "DE"},
		{Name: "Zimmermann", CountryCode: "DE"},
		{Name: "Hartmann", CountryCode: "DE"},
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
