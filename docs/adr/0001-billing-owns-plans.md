# ADR-0001: Billing owns payment plans

Status: Accepted (2022-11)

Payment plans and instalments moved out of policy-admin into billing-service. policy-admin keeps the premium; billing keeps how and when it is paid. The two talk through `policy.policy.issued`, `policy.policy.cancelled` and `billing.payment.missed`, never through each other's tables.
