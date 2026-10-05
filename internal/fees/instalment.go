// Package fees holds the instalment fee rule for monthly payment plans.
//
// Decision: docs/adr/0002-instalment-fee.md. Retention data on fees: Confluence TWBILL
// "Retention and fees, FY26".
package fees

import "github.com/tidewell/billing-service/internal/plans/model"

const (
	// FeeRate is the share of the annual premium charged for paying monthly.
	FeeRate = 0.06
	// AnnualCapPence caps the fee at £38 a year.
	AnnualCapPence = 3800
)

// InstalmentFee returns the yearly fee, in pence, for a plan.
//
// Every monthly plan pays the fee. There are no exemptions today: tenure, vulnerability
// and staff policies are all charged the same.
func InstalmentFee(plan model.Plan) int64 {
	if plan.Frequency != model.Monthly {
		return 0
	}
	fee := int64(float64(plan.AnnualPremiumPence) * FeeRate)
	if fee > AnnualCapPence {
		fee = AnnualCapPence
	}
	return fee
}

// PerInstalment spreads the yearly fee across the instalments, putting any remainder on the first.
func PerInstalment(plan model.Plan) []int64 {
	fee := InstalmentFee(plan)
	out := make([]int64, plan.Instalments)
	if plan.Instalments == 0 {
		return out
	}
	each := fee / int64(plan.Instalments)
	for i := range out {
		out[i] = each
	}
	out[0] += fee - each*int64(plan.Instalments)
	return out
}
