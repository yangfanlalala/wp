package wp

type Refund struct {
	RefundID string `json:"refund_id"`
	OutRefundNo string `json:"out_refund_no"`
	TransactionID string `json:"transaction_id"`
	OutTradeNo string `json:"out_trade_no"`
	Channel string `json:"channel"`
	UserReceivedAccount string `json:"user_received_account"`
	SuccessTime string `json:"success_time"`
	CreateTime string `json:"create_time"`
	Status string `json:"status"`
	FundsAccount string `json:"funds_account"`
	Amount Amount `json:"amount"`
	PromotionData string `json:"promotion_data"`
}
