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

	statuses, err := api.GetLineStatusByMode(ctx, []string{"tube"})
	if err != nil {
		slog.Error(err.Error())

		return
	}

	slog.Info("fetched tube line statuses", "statuses", statuses)
}
