package tfl_test

import (
	"context"
	"testing"

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

	expectedBikePoint := &tfl.Place{
		ID:           "BikePoints_1",
		URL:          "/Place/BikePoints_1",
		CommonName:   "River Street , Clerkenwell",
		Distance:     0,
		PlaceType:    "BikePoint",
		Children:     []tfl.Place{},
		ChildrenURLs: []string{},
		Lat:          51.529163,
		Lon:          -0.10997,
	}

	test.RequireBikePointsEqual(t, expectedBikePoint, &bikePoints[0])
}

func Test_GetBikePointIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	bikePoint, err := api.GetBikePoint(context.Background(), "BikePoints_1")
	require.NoError(t, err)

	expectedBikePoint := &tfl.Place{
		ID:           "BikePoints_1",
		URL:          "/Place/BikePoints_1",
		CommonName:   "River Street , Clerkenwell",
		Distance:     0,
		PlaceType:    "BikePoint",
		Children:     []tfl.Place{},
		ChildrenURLs: nil,
		Lat:          51.529163,
		Lon:          -0.10997,
	}

	test.RequireBikePointsEqual(t, expectedBikePoint, bikePoint)
}

func Test_SearchBikePointIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	bikePoints, err := api.SearchBikePoint(context.Background(), "River Street")
	require.NoError(t, err)

	expectedBikePoint := &tfl.Place{
		ID:           "BikePoints_1",
		URL:          "/Place/BikePoints_1",
		CommonName:   "River Street , Clerkenwell",
		Distance:     0,
		PlaceType:    "BikePoint",
		Children:     []tfl.Place{},
		ChildrenURLs: []string{},
		Lat:          51.529163,
		Lon:          -0.10997,
	}

	test.RequireBikePointsEqual(t, expectedBikePoint, &bikePoints[0])
}
