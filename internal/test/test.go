package test

import (
	"os"
	"testing"

	"github.com/jamesalexatkin/tfl-golang"
	"github.com/stretchr/testify/require"
)

const latLonEpsilon = 0.005

// SkipIfShort can be used to skip a test if the tests are run with the `-short` flag.
func SkipIfShort(t *testing.T) {
	t.Helper()

	if testing.Short() {
		t.Skip()
	}
}

// NewTestClient creates a new client for integration tests.
func NewTestClient() *tfl.Client {
	appID := os.Getenv("APP_ID")
	appKey := os.Getenv("APP_KEY")

	return tfl.New(appID, appKey)
}

// RequireBikePointsEqual ensures that two bike points are equal.
// This uses a shallow comparison, leaving out additional properties which have modification timestamps
// and would cause tests to be more brittle.
func RequireBikePointsEqual(t *testing.T, expectedBikePoint *tfl.Place, bikePoint *tfl.Place) {
	t.Helper()

	require.Equal(t, expectedBikePoint.ID, bikePoint.ID)
	require.Equal(t, expectedBikePoint.URL, bikePoint.URL)
	require.Equal(t, expectedBikePoint.CommonName, bikePoint.CommonName)
	require.Equal(t, expectedBikePoint.Distance, bikePoint.Distance)
	require.Equal(t, expectedBikePoint.PlaceType, bikePoint.PlaceType)
	require.Equal(t, expectedBikePoint.Children, bikePoint.Children)
	require.Equal(t, expectedBikePoint.ChildrenURLs, bikePoint.ChildrenURLs)
	require.InEpsilon(t, expectedBikePoint.Lat, bikePoint.Lat, latLonEpsilon)
	require.InEpsilon(t, expectedBikePoint.Lon, bikePoint.Lon, latLonEpsilon)
}
