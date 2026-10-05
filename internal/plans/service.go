package plans

import (
	"context"
	"time"

	"github.com/tidewell/billing-service/internal/clients"
	"github.com/tidewell/billing-service/internal/fees"
	"github.com/tidewell/billing-service/internal/plans/model"
)

type Service struct {
	store    *Store
	policies *clients.PolicyClient
	payments *clients.PaymentsClient
	identity *clients.IdentityClient
}

func NewService(s *Store, p *clients.PolicyClient, pay *clients.PaymentsClient, id *clients.IdentityClient) *Service {
	return &Service{store: s, policies: p, payments: pay, identity: id}
}

// CreateFromPolicy builds a plan when policy-admin publishes policy.policy.issued.
func (s *Service) CreateFromPolicy(ctx context.Context, policyID string, freq model.Frequency) (model.Plan, error) {
	pol, err := s.policies.Get(ctx, policyID) // GET /v1/policies/{id}
	if err != nil {
		return model.Plan{}, err
	}
	plan := model.Plan{
		PolicyID:           pol.ID,
		CustomerID:         pol.CustomerID,
		Frequency:          freq,
		AnnualPremiumPence: pol.AnnualPremiumPence,
		Instalments:        instalmentsFor(freq),
		StartDate:          pol.StartDate,
		Status:             "active",
	}
	plan.FeePence = fees.InstalmentFee(plan)
	return plan, s.store.Save(ctx, plan, schedule(plan))
}

// Close ends a plan on policy.policy.cancelled and works out the refund with Prorate.
func (s *Service) Close(ctx context.Context, policyID string, cancelledOn time.Time) error {
	plan, err := s.store.ByPolicy(ctx, policyID)
	if err != nil {
		return err
	}
	refund := Prorate(plan.AnnualPremiumPence, plan.StartDate, cancelledOn)
	return s.store.Close(ctx, plan.ID, refund)
}

func instalmentsFor(f model.Frequency) int {
	if f == model.Monthly {
		return 12
	}
	return 1
}

func schedule(p model.Plan) []model.Instalment {
	feeParts := fees.PerInstalment(p)
	each := p.AnnualPremiumPence / int64(p.Instalments)
	out := make([]model.Instalment, p.Instalments)
	for i := range out {
		out[i] = model.Instalment{
			PlanID: p.ID, Number: i + 1,
			DueDate:     p.StartDate.AddDate(0, i, 0),
			AmountPence: each + feeParts[i],
			FeePence:    feeParts[i],
			Status:      "due",
		}
	}
	return out
}
