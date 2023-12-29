package tfl

import (
	"context"
	"fmt"
	"strings"
)

// GetVehiclePredictions gets the predictions for a given list of vehicle IDs.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Vehicle/Vehicle_Get
func (c *Client) GetVehiclePredictions(ctx context.Context, vehicleIDs []string) ([]Prediction, error) {
	joinedIDs := strings.Join(vehicleIDs, ",")

	path := fmt.Sprintf("/Vehicle/%s/Arrivals", joinedIDs)

	predictions := []Prediction{}

	err := c.get(ctx, path, &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, err
}
