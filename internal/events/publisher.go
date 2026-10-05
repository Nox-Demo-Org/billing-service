package events

// InstalmentDue is the payload of billing.instalment.due.
type InstalmentDue struct {
	PlanID      string `json:"plan_id"`
	PolicyID    string `json:"policy_id"`
	CustomerID  string `json:"customer_id"`
	Instalment  int    `json:"instalment"`
	AmountPence int64  `json:"amount_pence"`
	FeePence    int64  `json:"fee_pence"`
	DueDate     string `json:"due_date"`
}

// PaymentMissed is the payload of billing.payment.missed.
type PaymentMissed struct {
	PlanID     string `json:"plan_id"`
	PolicyID   string `json:"policy_id"`
	Instalment int    `json:"instalment"`
	Attempts   int    `json:"attempts"`
}

type Publisher struct{ project string }

func NewPublisher(project string) *Publisher { return &Publisher{project: project} }
