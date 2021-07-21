package wp

type RefundPayload struct {
	TransactionID string `json:"transaction_id"`
	OutTradeNo string `json:"out_trade_no"`
	OutRefundNo string `json:"out_refund_no"`
	Reason string `json:"reason,omitempty"`
	NotifyURL string `json:"notify_url"`
	FundsAccount string `json:"funds_account,omitempty"`
	Amount Amount `json:"amount"`
}
