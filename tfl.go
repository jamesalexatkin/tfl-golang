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

func (c *Client) getWithQueryParams(ctx context.Context, path string, params map[string]string, responseBody any) error {
	path = ApiBaseURL + path

	params["app_id"] = c.appID
	params["app_key"] = c.appKey
	path = path + "?"
	for key, value := range params {
		path = path + fmt.Sprintf("%s=%s&", key, value)
	}

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

func (c *Client) get(ctx context.Context, path string, responseBody any) error {
	return c.getWithQueryParams(ctx, path, map[string]string{}, responseBody)
}
