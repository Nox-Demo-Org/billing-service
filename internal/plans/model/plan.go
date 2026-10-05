package model

import "time"

type Frequency string

const (
	Monthly Frequency = "monthly"
	Yearly  Frequency = "yearly"
)

// Plan is how a policy's premium is paid. Table: payment_plans.
type Plan struct {
	ID                 string    `json:"id"`
	PolicyID           string    `json:"policy_id"`
	CustomerID         string    `json:"customer_id"`
	Frequency          Frequency `json:"frequency"`
	AnnualPremiumPence int64     `json:"annual_premium_pence"`
	Instalments        int       `json:"instalments"`
	FeePence           int64     `json:"fee_pence"`
	StartDate          time.Time `json:"start_date"`
	Status             string    `json:"status"` // active | in_arrears | closed
}

// Instalment is one collection. Table: instalments.
type Instalment struct {
	PlanID      string    `json:"plan_id"`
	Number      int       `json:"number"`
	DueDate     time.Time `json:"due_date"`
	AmountPence int64     `json:"amount_pence"`
	FeePence    int64     `json:"fee_pence"`
	Status      string    `json:"status"` // due | collected | failed
}
