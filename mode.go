package tfl

import (
	"context"
	"fmt"
)

type ActiveServiceType struct {
	Mode        string `json:"mode"`
	ServiceType string `json:"serviceType"`
}

// GetActiveServiceTypes returns the service type active for a mode. Currently only supports tube.
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Mode/Mode_GetActiveServiceTypes
func (c *Client) GetActiveServiceTypes(ctx context.Context) ([]ActiveServiceType, error) {
	path := fmt.Sprintf("%s/Mode/ActiveServiceTypes", ApiBaseURL)

	activeServiceTypes := []ActiveServiceType{}
	err := c.get(ctx, path, &activeServiceTypes)
	if err != nil {
		return nil, err
	}

	return activeServiceTypes, err
}

// GetArrivalPredictionsForMode gets the next arrival predictions for all stops of a given mode.
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Mode/Mode_Arrivals
func (c *Client) GetArrivalPredictionsForMode(ctx context.Context, mode string, count int) ([]Prediction, error) {
	// TODO: figure out how count works
	path := fmt.Sprintf("%s/Mode/%s/Arrivals", ApiBaseURL, mode)

	predictions := []Prediction{}
	err := c.get(ctx, path, &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, err
}
