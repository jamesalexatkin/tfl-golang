package tfl

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// APIBaseURL is the base URL domain for the TfL API.
const APIBaseURL = "https://api.tfl.gov.uk"

// Client provides a mechanism to interact with the TfL API.
type Client struct {
	AppID      string
	AppKey     string
	APIBaseURL string
	HTTPClient *http.Client
}

// New returns a new client.
func New(appID string, appKey string) *Client {
	return &Client{
		AppID:      appID,
		AppKey:     appKey,
		APIBaseURL: APIBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) getWithQueryParams(
	ctx context.Context,
	path string,
	params map[string]string,
	responseBody any,
) error {
	path = c.APIBaseURL + path

	if c.AppID == "" {
		return errors.New("app id not set")
	}
	if c.AppKey == "" {
		return errors.New("app id not set")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "go")
	q := req.URL.Query()
	q.Add("app_id", c.AppID)
	q.Add("app_key", c.AppKey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyStr, _ := io.ReadAll(resp.Body)

		return HTTPError{Status: resp.Status, Body: string(bodyStr)}
	}

	err = json.NewDecoder(resp.Body).Decode(responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) get(ctx context.Context, path string, responseBody any) error {
	return c.getWithQueryParams(ctx, path, map[string]string{}, responseBody)
}
