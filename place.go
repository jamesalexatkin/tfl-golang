package tfl

import (
	"context"
	"fmt"
	"strings"

	"github.com/jamesalexatkin/tfl-go/internal/conv"
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
		queryParams = map[string]string{"activeOnly": conv.BoolToString(*activeOnly)}
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

// GetPlaceByGeoParams is the set of parameters available for the GetPlaceByGeo function.
type GetPlaceByGeoParams struct {
	Radius                 *float64
	Categories             []string
	IncludeChildren        *bool
	Type                   []string
	ActiveOnly             *bool
	NumberOfPlacesToReturn *int
	PlaceGeoSWLat          *float64
	PlaceGeoSWLon          *float64
	PlaceGeoNELat          *float64
	PlaceGeoNELon          *float64
	PlaceGeoLat            *float64
	PlaceGeoLon            *float64
}

// Gets the places that lie within a geographic region. The geographic region of interest can either be
// specified by using a lat/lon geo-point and a radius in metres to return places within the locus defined
// by the lat/lon of its centre or alternatively, by the use of a bounding box defined by the lat/lon of its
// north-west and south-east corners. Optionally filters on type and can strip properties for a smaller payload.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetByGeo
func (c *Client) GetPlaceByGeo(ctx context.Context, params GetPlaceByGeoParams) ([]Place, error) {
	path := "/Place"

	places := []Place{}

	var queryParams map[string]string
	if params.Radius != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.Radius)}
	}
	if params.Categories != nil {
		queryParams = map[string]string{"radius": conv.StringSliceToString(params.Categories)}
	}
	if params.IncludeChildren != nil {
		queryParams = map[string]string{"includeChildren": conv.BoolToString(*params.IncludeChildren)}
	}
	if params.Type != nil {
		queryParams = map[string]string{"radius": conv.StringSliceToString(params.Type)}
	}
	if params.ActiveOnly != nil {
		queryParams = map[string]string{"activeOnly": conv.BoolToString(*params.ActiveOnly)}
	}
	if params.NumberOfPlacesToReturn != nil {
		queryParams = map[string]string{"radius": conv.IntToString(*params.NumberOfPlacesToReturn)}
	}
	if params.PlaceGeoSWLat != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.PlaceGeoSWLat)}
	}
	if params.PlaceGeoSWLon != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.PlaceGeoSWLon)}
	}
	if params.PlaceGeoNELat != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.PlaceGeoNELat)}
	}
	if params.PlaceGeoNELon != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.PlaceGeoNELon)}
	}
	if params.PlaceGeoLat != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.PlaceGeoLat)}
	}
	if params.PlaceGeoLon != nil {
		queryParams = map[string]string{"radius": conv.Float64ToString(*params.PlaceGeoLon)}
	}

	err := c.getWithQueryParams(ctx, path, queryParams, &places)
	if err != nil {
		return nil, err
	}

	return places, nil
}

// Gets any places of the given type whose geography intersects the given latitude and longitude. In practice this means the Place must be polygonal e.g. a BoroughBoundary.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetAt

// Gets the place overlay for a given set of co-ordinates and a given width/height.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_GetOverlay

// Gets all places that matches the given query
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Place/Place_Search
