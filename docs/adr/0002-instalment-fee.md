# ADR-0002: Instalment fee on monthly plans

Status: Accepted (2023-03)

## Context
Monthly plans cost us more to run than yearly ones: twelve collections, failed debits and arrears letters. Finance asked for the cost to be recovered from customers who choose to pay monthly.

## Decision
Charge 6% of the annual premium on monthly plans, capped at £38 a year, spread across the instalments. The fee is worked out once, when the plan is created, in `internal/fees/instalment.go`.

## Consequences
- One rule for every customer. No exemptions were agreed at the time.
- Changing the rule only affects new plans and renewals; existing schedules keep their fee.
- The fee appears as its own line on `GET /v1/plans/{policyId}` so customer-portal can show it.
