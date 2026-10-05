// Package events wires billing-service to Pub/Sub.
//
// Subscribes to:  policy.policy.issued, policy.policy.cancelled, payments.collection.succeeded
// Publishes:      billing.instalment.due, billing.payment.missed
//
// Not subscribed: customer.profile.updated. Billing reads the address once, when the plan is
// created, so fee and arrears letters go to the address the customer had then.
package events

import "github.com/tidewell/billing-service/internal/plans"

const (
	TopicPolicyIssued       = "policy.policy.issued"
	TopicPolicyCancelled    = "policy.policy.cancelled"
	TopicCollectionSucceeded = "payments.collection.succeeded"
	TopicInstalmentDue      = "billing.instalment.due"
	TopicPaymentMissed      = "billing.payment.missed"
)

func Subscribe(project string, svc *plans.Service, pub *Publisher) {
	// subscription billing-policy-issued       -> svc.CreateFromPolicy
	// subscription billing-policy-cancelled    -> svc.Close
	// subscription billing-collection-succeeded -> store.MarkCollected
}
