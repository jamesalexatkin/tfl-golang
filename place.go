package tfl

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// GetAllPlaceCategories gets a list of all of the available place property categories and keys.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_MetaCategories
func (c *Client) GetAllPlaceCategories(ctx context.Context, carParkID string) ([]PlaceCategory, error) {
	path := "/Place/Meta/Categories"

	placeCategories := []PlaceCategory{}
	err := c.get(ctx, path, &placeCategories)
	if err != nil {
		return nil, err
	}

	return placeCategories, err
}

// GetAllPlaceTypes Gets a list of the available types of Place.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_MetaPlaceTypes
func (c *Client) GetAllPlaceTypes(ctx context.Context, carParkID string) ([]PlaceCategory, error) {
	path := "/Place/Meta/PlaceTypes"

	placeCategories := []PlaceCategory{}
	err := c.get(ctx, path, &placeCategories)
	if err != nil {
		return nil, err
	}

	return placeCategories, err
}

// Gets the set of streets associated with a post code.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetStreetsByPostCode
// TODO

// GetPlacesByType gets all places of a given type.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetByType
func (c *Client) GetPlacesByType(ctx context.Context, types []string, activeOnly *bool) ([]Place, error) {
	path := fmt.Sprintf("/Place/Type/%s", strings.Join(types, ","))

	places := []Place{}

	var queryParams map[string]string
	if activeOnly != nil {
		queryParams = map[string]string{"activeOnly": strconv.FormatBool(*activeOnly)}
	}

	err := c.getWithQueryParams(ctx, path, queryParams, &places)
	if err != nil {
		return nil, err
	}

	return places, nil
}

// GetPlace gets the place with the given id.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_Get
func (c *Client) GetPlace(ctx context.Context, bikePointID string) (*Place, error) {
	path := fmt.Sprintf("/Place/%s", bikePointID)

	place := Place{}
	err := c.get(ctx, path, &place)
	if err != nil {
		return nil, err
	}

	return &place, err
}

// Gets the places that lie within a geographic region. The geographic region of interest can either be specified by using a lat/lon geo-point and a radius in metres to return places within the locus defined by the lat/lon of its centre or alternatively, by the use of a bounding box defined by the lat/lon of its north-west and south-east corners. Optionally filters on type and can strip properties for a smaller payload.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetByGeo

// Gets any places of the given type whose geography intersects the given latitude and longitude. In practice this means the Place must be polygonal e.g. a BoroughBoundary.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetAt

// Gets the place overlay for a given set of co-ordinates and a given width/height.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetOverlay

// Gets all places that matches the given query
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_Search
