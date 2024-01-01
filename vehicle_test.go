package tfl_test

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jamesalexatkin/tfl-go"
	"github.com/jamesalexatkin/tfl-go/internal/test"
	"github.com/stretchr/testify/require"
)

func Test_GetVehiclePredictions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		testServer  *httptest.Server
		vehicleIDs  []string
		expected    []tfl.Prediction
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				predictions := []tfl.Prediction{{ID: "1234", VehicleID: "LDN123"}, {ID: "4567", VehicleID: "LDN456"}}
				bytes, err := json.Marshal(predictions)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			vehicleIDs: []string{"LDN123", "LDN456"},
			expected:   []tfl.Prediction{{ID: "1234", VehicleID: "LDN123"}, {ID: "4567", VehicleID: "LDN456"}},
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

			actual, err := c.GetVehiclePredictions(context.Background(), tc.vehicleIDs)
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
}

// TODO
func Test_GetVehiclePredictionsIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	// TODO: figure out how to get valid IDs for vehicles, maybe by getting stop points and using a real vehicle ID?
	vehiclePredictions, err := api.GetVehiclePredictions(context.Background(), []string{"LX58CFV,LX11AZB,LX58CFE"})
	require.NoError(t, err)

	// expectedBikePoint := &tfl.Place{
	// 	ID:           "BikePoints_1",
	// 	URL:          "/Place/BikePoints_1",
	// 	CommonName:   "River Street , Clerkenwell",
	// 	Distance:     0,
	// 	PlaceType:    "BikePoint",
	// 	Children:     []tfl.Place{},
	// 	ChildrenURLs: nil,
	// 	Lat:          51.529163,
	// 	Lon:          -0.10997,
	// }

	// expectedPrediction := &tfl.Prediction{}

	slog.Info("predictions", "vp", vehiclePredictions)
	t.Fail()

	// test.RequireBikePointsEqual(t, expectedBikePoint, vehiclePredictions)
}
