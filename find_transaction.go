package wp

type FindTransactionPayload struct {
}

type Transaction struct {
	AppID string `json:"appid"`
	MerchantID string `json:"mchid"`
	OutTradeNo string `json:"out_trade_no"`
	TransactionID string `json:"transaction_id"`
	TradeType string `json:"trade_type"`
	TradeState string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
	BankType string `json:"bank_type"`
	Attach string `json:"attach"`
	SuccessTime string `json:"success_time"`
	Payer string `json:"payer"`
	Amount string `json:"amount"`
	SceneInfo string `json:"scene_info"`
	PromotionDetail string `json:"promotion_detail"`
}

