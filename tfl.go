package tfl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const ApiBaseURL = "https://api.tfl.gov.uk"

type Client struct {
	httpClient *http.Client
	appID      string
	appKey     string
}

func (c *Client) get(ctx context.Context, path string, responseBody any) error {
	path = addQueryParamsToPath(c.appID, c.appKey)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(responseBody)
	if err != nil {
		return err
	}

	return nil
}

func addQueryParamsToPath(appID string, appKey string) string {
	const query = "?app_id=%s&app_key=%s"
	return fmt.Sprintf(query, appID, appKey)
}
