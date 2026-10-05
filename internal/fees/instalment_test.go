package fees

import (
	"testing"

	"github.com/tidewell/billing-service/internal/plans/model"
)

func TestYearlyPlansPayNoFee(t *testing.T) {
	p := model.Plan{Frequency: model.Yearly, AnnualPremiumPence: 48000, Instalments: 1}
	if got := InstalmentFee(p); got != 0 {
		t.Fatalf("want 0, got %d", got)
	}
}

func TestMonthlyFeeIsCapped(t *testing.T) {
	p := model.Plan{Frequency: model.Monthly, AnnualPremiumPence: 120000, Instalments: 12}
	if got := InstalmentFee(p); got != AnnualCapPence {
		t.Fatalf("want cap %d, got %d", AnnualCapPence, got)
	}
}

func TestFeeRemainderGoesOnFirstInstalment(t *testing.T) {
	p := model.Plan{Frequency: model.Monthly, AnnualPremiumPence: 50000, Instalments: 12}
	parts := PerInstalment(p)
	var sum int64
	for _, x := range parts {
		sum += x
	}
	if sum != InstalmentFee(p) {
		t.Fatalf("instalments %d do not add up to fee %d", sum, InstalmentFee(p))
	}
}
