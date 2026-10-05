# billing-service

Payment plans for Tidewell Mutual policies: yearly or monthly, the instalment fee on monthly plans, collection dates and missed payments.

Owned by the **Billing & Payments** team. On-call: `#tw-billing`.

## What it exposes

| Contract | Kind | Notes |
| --- | --- | --- |
| `GET /v1/plans/{policyId}` | REST | The plan, schedule and fees for a policy. Used by customer-portal. |
| `POST /v1/plans` | REST | Creates a plan. Called internally when `policy.policy.issued` arrives. |
| `billing.instalment.due` | Event (Pub/Sub) | Published 5 days before each collection. notifications-hub and payments-gateway listen. |
| `billing.payment.missed` | Event (Pub/Sub) | Published when a collection fails twice. policy-admin and notifications-hub listen. |

## What it depends on

| Contract | Owner | Why |
| --- | --- | --- |
| `GET /v1/policies/{id}` | policy-admin | Premium, start date and customer id for a new plan |
| `POST /v1/collections` | payments-gateway | Collects each instalment |
| `GET /v1/customers/{id}` | customer-identity | Contact details for fee letters (`internal/clients/identity.go`) |
| `POST /v1/messages` | notifications-hub | Fee and arrears letters |
| `policy.policy.issued`, `policy.policy.cancelled` | policy-admin | Create and close plans |
| `payments.collection.succeeded` | payments-gateway | Marks an instalment paid |

## Fees

Monthly plans carry an instalment fee: 6% of the annual premium, capped at £38 a year, spread across the 12 instalments. Yearly plans carry no fee. The rule lives in `internal/fees/instalment.go`. See [ADR-0002](docs/adr/0002-instalment-fee.md).

## Layout

```text
cmd/billing/main.go            wiring
internal/api/plans.go          REST handlers
internal/plans/                plan creation, schedule, proration
internal/fees/instalment.go    the instalment fee rule
internal/clients/              policy-admin, payments-gateway, customer-identity clients
internal/events/               Pub/Sub publishers and subscriptions
```
