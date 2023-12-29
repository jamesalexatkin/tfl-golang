package tfl_test

import (
	"context"
	"testing"
	"time"

	"github.com/jamesalexatkin/tfl-go"
	"github.com/jamesalexatkin/tfl-go/internal/test"
	"github.com/stretchr/testify/require"
)

// TODO:
// GetAll
// Get
// Search

func Test_GetAllBikePointsIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	bikePoints, err := api.GetAllBikePoints(context.Background())
	require.NoError(t, err)

	expectedDetail := tfl.AccidentDetail{
		ID:       243890,
		Lat:      51.505742,
		Lon:      -0.131005,
		Location: "The Mall junction with Horse Guards Road",
		Date:     time.Date(2015, time.July, 4, 1, 55, 0, 0, time.UTC),
		Severity: "Slight",
		Borough:  "City of Westminster",
		Casualties: []tfl.Casualty{
			{
				Age:      30,
				Class:    "Driver",
				Severity: "Slight",
				Mode:     "PedalCycle",
				AgeBand:  "Adult",
			},
		},
		Vehicles: []tfl.Vehicle{
			{Type: "Taxi"},
			{Type: "PedalCycle"},
		},
	}

	require.Contains(t, bikePoints, expectedDetail)
}

func Test_GetBikePointIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	bikePoint, err := api.GetBikePoint(context.Background(), "BikePoints_1")
	require.NoError(t, err)

	// Subset of bike point because the additional properties have modified timestamps
	// Comparing these would make the test brittle
	expectedBikePoint := &tfl.Place{
		ID:           "BikePoints_1",
		URL:          "/Place/BikePoints_1",
		CommonName:   "River Street , Clerkenwell",
		Distance:     0,
		PlaceType:    "BikePoint",
		Children:     []tfl.Place{},
		ChildrenURLs: []string(nil),
		Lat:          51.529163,
		Lon:          -0.10997,
	}

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
