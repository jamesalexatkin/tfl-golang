package tfl_test

import (
	"context"
	"testing"

	"github.com/jamesalexatkin/tfl-go"
	"github.com/stretchr/testify/require"
)

func Test_GetAccidentDetails(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		c       *tfl.Client
		year    int
		want    []tfl.AccidentDetail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt := tt

			got, err := tt.c.GetAccidentDetails(context.Background(), tt.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
