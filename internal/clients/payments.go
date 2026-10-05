package clients

import (
	"context"
	"net/http"
	"time"
)

// CollectionRequest is the body of payments-gateway's POST /v1/collections.
type CollectionRequest struct {
	PlanID        string `json:"plan_id"`
	Instalment    int    `json:"instalment"`
	AmountPence   int64  `json:"amount_pence"`
	Method        string `json:"method"` // card | direct_debit
	IdempotencyID string `json:"idempotency_id"`
}

type PaymentsClient struct {
	base string
	http *http.Client
}

func NewPaymentsClient(base string) *PaymentsClient {
	return &PaymentsClient{base: base, http: &http.Client{Timeout: 5 * time.Second}}
}

func (c *PaymentsClient) Collect(ctx context.Context, r CollectionRequest) error {
	return postJSON(ctx, c.http, c.base+"/collections", r)
}
