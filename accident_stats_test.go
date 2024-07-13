package tfl_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jamesalexatkin/tfl-golang"
	"github.com/jamesalexatkin/tfl-golang/internal/test"
	"github.com/stretchr/testify/require"
)

func Test_GetAccidentDetails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		year        int
		testServer  *httptest.Server
		expected    []tfl.AccidentDetail
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				accidentDetails := []tfl.AccidentDetail{
					{ID: 123},
				}
				bytes, err := json.Marshal(accidentDetails)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			expected: []tfl.AccidentDetail{
				{ID: 123},
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

			actual, err := c.GetAccidentDetails(context.Background(), tc.year)
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_GetAccidentDetailsIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	accidentDetails, err := api.GetAccidentDetails(context.Background(), 2015)
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

	require.Contains(t, accidentDetails, expectedDetail)
}
