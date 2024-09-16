package tfl

import (
	"context"
	"fmt"

	"github.com/jamesalexatkin/tfl-golang/internal/conv"
)

// GetValidModes gets a list of valid modes.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_MetaModes
func (c *Client) GetValidModes(ctx context.Context) ([]Mode, error) {
	path := "/Line/Meta/Modes"

	modes := []Mode{}
	err := c.get(ctx, path, &modes)
	if err != nil {
		return nil, err
	}

	return modes, err
}

// GetLineStatusByMode gets the line status of all lines for the given modes.
//
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/Line/Line_StatusByMode
func (c *Client) GetLineStatusByMode(ctx context.Context, modes []string) ([]Status, error) {
	path := fmt.Sprintf("/Line/Mode/%s/Status", conv.StringSliceToString(modes))

	statuses := []Status{}
	err := c.get(ctx, path, &statuses)
	if err != nil {
		return nil, err
	}

	return statuses, err
}
