package tfl

import (
	"context"
	"fmt"
)

// GetAccidentDetails gets all accident details for accidents occuring in the specified year.
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/AccidentStats/AccidentStats_Get
func (c *Client) GetAccidentDetails(ctx context.Context, year int) ([]AccidentDetail, error) {
	path := fmt.Sprintf("/AccidentStats/%d", year)

	accidentDetails := []AccidentDetail{}
	err := c.get(ctx, path, &accidentDetails)
	if err != nil {
		return nil, err
	}

	return accidentDetails, err
}
