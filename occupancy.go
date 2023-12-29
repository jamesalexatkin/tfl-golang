package tfl

import (
	"context"
	"fmt"
	"strings"
)

// GetCarParkOccupancy returns the occupancy for a car park with a given ID.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Occupancy/Occupancy_Get
func (c *Client) GetCarParkOccupancy(ctx context.Context, carParkID string) (*CarParkOccupancy, error) {
	path := fmt.Sprintf("/Occupancy/CarPark/%s", carParkID)

	carParkOccupancy := CarParkOccupancy{}

	err := c.get(ctx, path, &carParkOccupancy)

	if err != nil {
		return nil, err
	}

	return &carParkOccupancy, err
}

// GetAllCarParkOccupancies returns the occupancy for all car parks that have occupancy data.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Occupancy/Occupancy_Get_0
func (c *Client) GetAllCarParkOccupancies(ctx context.Context) ([]CarParkOccupancy, error) {
	path := "/Occupancy/CarPark"

	carParkOccupancies := []CarParkOccupancy{}

	err := c.get(ctx, path, &carParkOccupancies)

	if err != nil {
		return nil, err
	}

	return carParkOccupancies, err
}

// GetChargeConnectorOccupancy gets the occupancy for a charge connector with a given ID (sourceSystemPlaceId).
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Occupancy/Occupancy_GetChargeConnectorStatus
func (c *Client) GetChargeConnectorOccupancy(
	ctx context.Context,
	chargeConnectorIDs []string,
) (
	[]ChargeConnectorOccupancy,
	error,
) {
	path := fmt.Sprintf("/Occupancy/ChargeConnector/%s", strings.Join(chargeConnectorIDs, ","))

	chargeConnectorOccupancies := []ChargeConnectorOccupancy{}

	err := c.get(ctx, path, &chargeConnectorOccupancies)

	if err != nil {
		return []ChargeConnectorOccupancy{}, err
	}

	return chargeConnectorOccupancies, err
}

// GetAllChargeConnectorOccupancies gets the occupancy for all charge connectors.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Occupancy/Occupancy_GetAllChargeConnectorStatus
func (c *Client) GetAllChargeConnectorOccupancies(ctx context.Context) ([]ChargeConnectorOccupancy, error) {
	path := "/Occupancy/ChargeConnector"

	chargeConnectorOccupancies := []ChargeConnectorOccupancy{}

	err := c.get(ctx, path, &chargeConnectorOccupancies)

	if err != nil {
		return nil, err
	}

	return chargeConnectorOccupancies, err

}

// GetBikePointOccupancies gets the occupancy for bike points.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Occupancy/Occupancy_GetBikePointsOccupancies
func (c *Client) GetBikePointOccupancies(ctx context.Context, bikePointIDs []string) ([]BikePointOccupancy, error) {
	path := fmt.Sprintf("/Occupancy/ChargeConnector/%s", strings.Join(bikePointIDs, ","))

	bikePointOccupancies := []BikePointOccupancy{}

	err := c.get(ctx, path, &bikePointOccupancies)
	if err != nil {
		return nil, err
	}

	return bikePointOccupancies, err
}
