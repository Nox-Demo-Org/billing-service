package clients

import (
	"context"
	"net/http"
	"time"
)

// Policy is the subset of policy-admin's GET /v1/policies/{id} that billing reads.
type Policy struct {
	ID                 string    `json:"id"`
	CustomerID         string    `json:"customer_id"`
	AnnualPremiumPence int64     `json:"annual_premium_pence"`
	StartDate          time.Time `json:"start_date"`
	Status             string    `json:"status"`
}

type PolicyClient struct {
	base string
	http *http.Client
}

func NewPolicyClient(base string) *PolicyClient {
	return &PolicyClient{base: base, http: &http.Client{Timeout: 2 * time.Second}}
}

func (c *PolicyClient) Get(ctx context.Context, id string) (Policy, error) {
	return getJSON[Policy](ctx, c.http, c.base+"/policies/"+id)
}
