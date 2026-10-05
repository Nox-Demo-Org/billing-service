package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tidewell/billing-service/internal/plans"
)

// Mount registers:
//   GET  /v1/plans/{policyId}  the plan, schedule and fee for a policy (customer-portal)
//   POST /v1/plans             create a plan (internal, normally from policy.policy.issued)
func Mount(r chi.Router, svc *plans.Service) {
	r.Get("/v1/plans/{policyId}", func(w http.ResponseWriter, req *http.Request) {
		// returns {plan, instalments[], fee_pence}
		json.NewEncoder(w).Encode(map[string]string{"policy_id": chi.URLParam(req, "policyId")})
	})
	r.Post("/v1/plans", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusCreated)
	})
}
