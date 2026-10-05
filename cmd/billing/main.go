package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/tidewell/billing-service/internal/api"
	"github.com/tidewell/billing-service/internal/clients"
	"github.com/tidewell/billing-service/internal/events"
	"github.com/tidewell/billing-service/internal/plans"
)

func main() {
	policies := clients.NewPolicyClient(os.Getenv("POLICY_ADMIN_URL"))     // http://policy-admin/v1
	payments := clients.NewPaymentsClient(os.Getenv("PAYMENTS_URL"))       // http://payments-gateway/v1
	identity := clients.NewIdentityClient(os.Getenv("CUSTOMER_IDENTITY_URL")) // http://customer-identity/v1

	store := plans.NewStore(os.Getenv("DATABASE_URL"))
	svc := plans.NewService(store, policies, payments, identity)

	pub := events.NewPublisher(os.Getenv("GCP_PROJECT"))
	events.Subscribe(os.Getenv("GCP_PROJECT"), svc, pub)

	r := chi.NewRouter()
	api.Mount(r, svc)
	log.Println("billing-service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
