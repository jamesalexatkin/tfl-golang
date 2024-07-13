package tfl_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jamesalexatkin/tfl-golang"
	"github.com/jamesalexatkin/tfl-golang/internal/test"
	"github.com/stretchr/testify/require"
)

func Test_GetAllBikePoints(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testServer  *httptest.Server
		expected    []tfl.Place
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				bikePoints := []tfl.Place{
					{ID: "BikePoint_1", CommonName: "Cycle Street"},
				}
				bytes, err := json.Marshal(bikePoints)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			expected: []tfl.Place{
				{ID: "BikePoint_1", CommonName: "Cycle Street"},
			},
		},
		{
			name: "Fail on non-200 response code",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
			})),
			expected:    nil,
			expectedErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			defer tc.testServer.Close()

			c := tfl.Client{
				AppID:      "app_id",
				AppKey:     "app_key",
				APIBaseURL: tc.testServer.URL,
				HTTPClient: &http.Client{},
			}

			actual, err := c.GetAllBikePoints(context.Background())
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
}

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

func Test_GetBikePoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testServer  *httptest.Server
		bikePointID string
		expected    *tfl.Place
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				bikePoint := &tfl.Place{ID: "BikePoint_1", CommonName: "Cycle Street"}
				bytes, err := json.Marshal(bikePoint)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			bikePointID: "BikePoint_1",
			expected:    &tfl.Place{ID: "BikePoint_1", CommonName: "Cycle Street"},
		},
		{
			name: "Fail on non-200 response code",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
			})),
			expected:    nil,
			expectedErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			defer tc.testServer.Close()

			c := tfl.Client{
				AppID:      "app_id",
				AppKey:     "app_key",
				APIBaseURL: tc.testServer.URL,
				HTTPClient: &http.Client{},
			}

			actual, err := c.GetBikePoint(context.Background(), tc.bikePointID)
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
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

func Test_SearchBikePoint(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testServer  *httptest.Server
		searchQuery string
		expected    []tfl.Place
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				bikePoint := []tfl.Place{
					{ID: "BikePoint_1", CommonName: "Cycle Street"},
				}
				bytes, err := json.Marshal(bikePoint)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			searchQuery: "Cycle Street",
			expected: []tfl.Place{
				{ID: "BikePoint_1", CommonName: "Cycle Street"},
			},
		},
		{
			name: "Fail on non-200 response code",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusForbidden)
			})),
			expected:    nil,
			expectedErr: true,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			defer tc.testServer.Close()

			c := tfl.Client{
				AppID:      "app_id",
				AppKey:     "app_key",
				APIBaseURL: tc.testServer.URL,
				HTTPClient: &http.Client{},
			}

			actual, err := c.SearchBikePoint(context.Background(), tc.searchQuery)
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
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
