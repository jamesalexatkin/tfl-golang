package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jamesalexatkin/tfl-golang"
)

func main() {
	ctx := context.Background()
	appID := os.Getenv("APP_ID")
	appKey := os.Getenv("APP_KEY")

	api := tfl.New(appID, appKey)

	bikePoints, err := api.GetAllBikePoints(ctx)
	if err != nil {
		slog.Error(err.Error())

		return
	}

	slog.Info("fetched bike points", "bikePoints", bikePoints)

	bikePoint, err := api.GetBikePoint(ctx, "BikePoints_1")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	slog.Info("fetched bike point", "bikePoint", bikePoint)
}
