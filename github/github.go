package github

import (
	"io"
	"net/http"
)

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		// Drain and close the body to allow connection reuse
		defer resp.Body.Close()
		_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 1024))
		return nil, err
	}

	// ... existing logic ...
	return nil, nil
}