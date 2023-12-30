package test

import (
	"os"
	"testing"

	"github.com/jamesalexatkin/tfl-go"
	"github.com/stretchr/testify/require"
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

func RequireBikePointsEqual(t *testing.T, expectedBikePoint *tfl.Place, bikePoint *tfl.Place) {
	require.Equal(t, expectedBikePoint.ID, bikePoint.ID)
	require.Equal(t, expectedBikePoint.URL, bikePoint.URL)
	require.Equal(t, expectedBikePoint.CommonName, bikePoint.CommonName)
	require.Equal(t, expectedBikePoint.Distance, bikePoint.Distance)
	require.Equal(t, expectedBikePoint.PlaceType, bikePoint.PlaceType)
	require.Equal(t, expectedBikePoint.Children, bikePoint.Children)
	require.Equal(t, expectedBikePoint.ChildrenURLs, bikePoint.ChildrenURLs)
	require.Equal(t, expectedBikePoint.Lat, bikePoint.Lat)
	require.Equal(t, expectedBikePoint.Lon, bikePoint.Lon)
}
