package tfl

import (
	"context"
	"fmt"
)

// GetAllBikePoints gets all bike point locations.
// The Place object has an addtionalProperties array which contains the nbBikes, nbDocks and nbSpaces numbers which
// give the status of the BikePoint. A mismatch in these numbers i.e. nbDocks - (nbBikes + nbSpaces) != 0 indicates
// broken docks.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/BikePoint/BikePoint_GetAll
func (c *Client) GetAllBikePoints(ctx context.Context) ([]Place, error) {
	path := "/BikePoint"

	places := []Place{}
	err := c.get(ctx, path, &places)
	if err != nil {
		return nil, err
	}

	return places, err
}

// GetBikePoint gets the bike point with the given ID.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/BikePoint/BikePoint_Get
func (c *Client) GetBikePoint(ctx context.Context, bikePointID string) (*Place, error) {
	path := fmt.Sprintf("/BikePoint/%s", bikePointID)

	place := Place{}
	err := c.get(ctx, path, &place)
	if err != nil {
		return nil, err
	}

	return &place, err
}

// SearchBikePoint searches for bike stations by their name, a bike point's name often contains information about the
// name of the street or nearby landmarks, for example. Note that the search result does not contain the
// PlaceProperties i.e. the status or occupancy of the BikePoint, to get that information you should retrieve the
// BikePoint by its id on /BikePoint/id.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/BikePoint/BikePoint_Search
func (c *Client) SearchBikePoint(ctx context.Context, searchQuery string) (*Place, error) {
	path := "/BikePoint/Search"

	place := Place{}
	err := c.getWithQueryParams(ctx, path, map[string]string{
		"query": searchQuery,
	}, &place)
	if err != nil {
		return nil, err
	}

	return &place, err
}
