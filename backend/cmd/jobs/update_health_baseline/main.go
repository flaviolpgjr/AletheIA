package main

import (
	"context"
	"log"
	"os"

	"github.com/flaviolpgjr/aletheia/backend/internal/database"
	"github.com/flaviolpgjr/aletheia/backend/internal/publicdata/health"
	"github.com/flaviolpgjr/aletheia/backend/internal/repositories"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()

	dbPool, err := database.NewPostgresPool(
		ctx,
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	healthClient := health.NewClient()

	baseline, err := healthClient.FetchHospitalFacilitiesBaseline(ctx)
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewPublicDataBaselineRepository(dbPool)

	if err := repository.Save(ctx, baseline); err != nil {
		log.Fatal(err)
	}

	log.Printf(
		"health baseline updated: indicator=%s scope=%s value=%.0f unit=%s",
		baseline.Indicator,
		baseline.Scope,
		baseline.Value,
		baseline.Unit,
	)
}