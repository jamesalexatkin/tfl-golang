package test

import (
	"os"
	"testing"

	"github.com/jamesalexatkin/tfl-go"
)

func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
}

func NewTestClient() *tfl.Client {
	appID := os.Getenv("APP_ID")
	appKey := os.Getenv("APP_KEY")
	return tfl.New(appID, appKey)
}
