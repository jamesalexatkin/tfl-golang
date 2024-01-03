package tfl_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jamesalexatkin/tfl-go"
	"github.com/jamesalexatkin/tfl-go/internal/test"
	"github.com/stretchr/testify/require"
)

func Test_GetActiveServiceTypes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		year        int
		testServer  *httptest.Server
		expected    []tfl.ActiveServiceType
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				activeServiceTypes := []tfl.ActiveServiceType{
					{Mode: "example_mode", ServiceType: "Regular"},
				}
				bytes, err := json.Marshal(activeServiceTypes)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			expected: []tfl.ActiveServiceType{
				{Mode: "example_mode", ServiceType: "Regular"},
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

			actual, err := c.GetActiveServiceTypes(context.Background())
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetActiveServiceTypes() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_GetActiveServiceTypesIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	activeServiceTypes, err := api.GetActiveServiceTypes(context.Background())
	require.NoError(t, err)

	expectedServiceTypes := []tfl.ActiveServiceType{
		{Mode: "tube", ServiceType: "Regular"},
	}

	require.Equal(t, expectedServiceTypes, activeServiceTypes)
}

func Test_GetArrivalPredictionsForMode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		year        int
		testServer  *httptest.Server
		expected    []tfl.Prediction
		expectedErr bool
	}{
		{
			name: "Success",
			testServer: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				arrivalPredictions := []tfl.Prediction{
					{ID: "123", LineName: "bakerloo"},
				}
				bytes, err := json.Marshal(arrivalPredictions)
				require.NoError(t, err)
				w.WriteHeader(http.StatusOK)
				_, err = w.Write(bytes)
				require.NoError(t, err)
			})),
			expected: []tfl.Prediction{
				{ID: "123", LineName: "bakerloo"},
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

			actual, err := c.GetArrivalPredictionsForMode(context.Background(), "tube", 1)
			if (err != nil) != tc.expectedErr {
				t.Errorf("Client.GetArrivalPredictionsForMode() error = %v, wantErr %v", err, tc.expectedErr)

				return
			}
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_GetArrivalPredictionsForModeIntegration(t *testing.T) {
	t.Parallel()

	test.SkipIfShort(t)

	api := test.NewTestClient()

	_, err := api.GetArrivalPredictionsForMode(context.Background(), "tube", 1)
	require.NoError(t, err)
}
