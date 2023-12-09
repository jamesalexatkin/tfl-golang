package tfl

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

const ApiBaseURL = "https://api.tfl.gov.uk"

type Client struct {
	appID      string
	appKey     string
	httpClient *http.Client
}

func New(appID string, appKey string) *Client {
	return &Client{
		appID:      appID,
		appKey:     appKey,
		httpClient: &http.Client{},
	}
}

func (c *Client) getWithQueryParams(ctx context.Context, path string, params map[string]string, responseBody any) error {
	path = ApiBaseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("User-Agent", "go")
	q := req.URL.Query()
	q.Add("app_id", c.appID)
	q.Add("app_key", c.appKey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	slog.Info(req.Method + " " + req.URL.String())

	resp, err := c.httpClient.Do(req)
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
