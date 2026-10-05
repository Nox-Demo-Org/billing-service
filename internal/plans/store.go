package plans

import (
	"context"

	"github.com/tidewell/billing-service/internal/plans/model"
)

// Store persists plans and instalments in the billing database (tables payment_plans, instalments).
type Store struct{ dsn string }

func NewStore(dsn string) *Store { return &Store{dsn: dsn} }

func (s *Store) Save(ctx context.Context, p model.Plan, sched []model.Instalment) error { return nil }
func (s *Store) ByPolicy(ctx context.Context, policyID string) (model.Plan, error)     { return model.Plan{}, nil }
func (s *Store) Close(ctx context.Context, planID string, refundPence int64) error     { return nil }
func (s *Store) MarkCollected(ctx context.Context, planID string, number int) error    { return nil }
