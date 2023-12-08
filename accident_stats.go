package tfl

import (
	"context"
	"fmt"
	"time"
)

type AccidentDetail struct {
	ID         int        `json:"id"`
	Lat        int        `json:"lat"`
	Lon        int        `json:"lon"`
	Location   string     `json:"location"`
	Date       time.Time  `json:"date"`
	Severity   string     `json:"severity"`
	Borough    string     `json:"borough"`
	Casulaties []Casualty `json:"casualties"`
	Vehicles   []Vehicle  `json:"vehicles"`
}

type Casualty struct {
	Age      int    `json:"age"`
	Class    string `json:"class"`
	Severity string `json:"severity"`
	Mode     string `json:"mode"`
	AgeBand  string `json:"ageBand"`
}

type Vehicle struct {
	Type string `json:"type"`
}

// GetAccidentDetails gets all accident details for accidents occuring in the specified year.
// https://api.tfl.gov.uk/swagger/ui/index.html?url=/swagger/docs/v1#!/AccidentStats/AccidentStats_Get
func (c *Client) GetAccidentDetails(ctx context.Context, year int) ([]AccidentDetail, error) {
	path := fmt.Sprintf("%s/AccidentStats/%d", ApiBaseURL, year)

	accidentDetails := []AccidentDetail{}
	err := c.get(ctx, path, &accidentDetails)
	if err != nil {
		return nil, err
	}

	return accidentDetails, err
}
