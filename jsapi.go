package wp

type PayPayload struct {
	AppID string `json:"appid"`
	MerchantID string `json:"mchid"`
	Description string `json:"description"`
	OutTradeNo string `json:"out_trade_no"`
	TimeExpire string `json:"time_expire"`
	Attach string `json:"attach"`
	NotifyURL string `json:"notify_url"`
	GoodsTag string `json:"goods_tag"`
	Amount string `json:"amount"`
	Payer string `json:"payer"`
	Detail string `json:"detail"`
	SceneInfo string `json:"scene_info"`
	SettleInfo string `json:"settle_info"`
}
type JSAPIReply struct {
	PrepayID string `json:"prepay_id"`
}
func CreatePrepayID () (string, error) {
	return "", nil
}