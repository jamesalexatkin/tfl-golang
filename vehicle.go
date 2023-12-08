package tfl

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type Prediction struct {
	ID                  string            `json:"id"`
	OperationType       int               `json:"operationType"`
	VehicleID           string            `json:"vehicleId"`
	NaptanID            string            `json:"naptanId"`
	StationName         string            `json:"stationName"`
	LineID              string            `json:"lineId"`
	LineName            string            `json:"lineName"`
	PlatformName        string            `json:"platformName"`
	Direction           string            `json:"direction"`
	Bearing             string            `json:"bearing"`
	DestinationNaptanID string            `json:"destinationNaptanId"`
	DestinationName     string            `json:"destinationName"`
	Timestamp           time.Time         `json:"timestamp"`
	TimeToStation       int               `json:"timeToStation"`
	CurrentLocation     string            `json:"currentLocation"`
	Towards             string            `json:"towards"`
	ExpectedArrival     time.Time         `json:"expectedArrival"`
	TimeToLive          time.Time         `json:"timeToLive"`
	ModeName            string            `json:"modeName"`
	Timing              PredicitionTiming `json:"timing"`
}

type PredicitionTiming struct {
	CountdownServerAdjustment string `json:"countdownServerAdjustment"`
	Source                    string `json:"source"`
	Insert                    string `json:"insert"`
	Read                      string `json:"read"`
	Sent                      string `json:"sent"`
	Received                  string `json:"received"`
}

// GetVehiclePredictions gets the predictions for a given list of vehicle IDs.
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Vehicle/Vehicle_Get
func (c *Client) GetVehiclePredictions(ctx context.Context, vehicleIDs []string) ([]Prediction, error) {
	joinedIDs := strings.Join(vehicleIDs, ",")
	path := fmt.Sprintf("%s/Vehicle/%s/Arrivals", ApiBaseURL, joinedIDs)

	predictions := []Prediction{}
	err := c.get(ctx, path, &predictions)
	if err != nil {
		return nil, err
	}

	return predictions, err
}
