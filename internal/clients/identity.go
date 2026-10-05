package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Customer is the profile returned by customer-identity's GET /v1/customers/{id}.
type Customer struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Postcode      string    `json:"postcode"`
	CustomerSince time.Time `json:"customer_since"`
	TenureYears   int       `json:"tenure_years"`
}

type IdentityClient struct {
	base string
	http *http.Client
}

func NewIdentityClient(base string) *IdentityClient {
	return &IdentityClient{base: base, http: &http.Client{Timeout: 2 * time.Second}}
}

// Get calls GET /v1/customers/{id}. Billing only uses the name and email today, for fee and
// arrears letters; tenure_years is returned but not used.
func (c *IdentityClient) Get(ctx context.Context, id string) (Customer, error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/customers/%s", c.base, id), nil)
	res, err := c.http.Do(req)
	if err != nil {
		return Customer{}, err
	}
	defer res.Body.Close()
	var out Customer
	return out, json.NewDecoder(res.Body).Decode(&out)
}
