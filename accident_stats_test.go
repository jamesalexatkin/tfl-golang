package tfl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetAccidentDetails(t *testing.T) {
	type args struct {
		ctx  context.Context
		year int
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    []AccidentDetail
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetAccidentDetails(tt.args.ctx, tt.args.year)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetAccidentDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
