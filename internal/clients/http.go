package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func getJSON[T any](ctx context.Context, c *http.Client, url string) (T, error) {
	var out T
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	res, err := c.Do(req)
	if err != nil {
		return out, err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return out, fmt.Errorf("GET %s: %s", url, res.Status)
	}
	return out, json.NewDecoder(res.Body).Decode(&out)
}

func postJSON(ctx context.Context, c *http.Client, url string, body any) error {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode >= 300 {
		return fmt.Errorf("POST %s: %s", url, res.Status)
	}
	return nil
}
