package wp

import (
	"fmt"
)

const (
	ApiJSAPI = "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi`"
	ApiFindTransaction = "https://api.mch.weixin.qq.com/v3/pay/transactions/id/{transaction_id}"
	ApiRefund = "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds"

	ErrorCodeTrade = "TRADE_ERROR" // 交易错误
	ErrorCodeSystem = "SYSTEM_ERROR" // 系统错误
	ErrorCodeSign = "SIGN_ERROR" // 签名错误
	ErrorCodeRuleLimit = "RULE_LIMIT" // 业务规则限制
	ErrorCodeParam = "PARAM_ERROR" // 参数错误
	ErrorCodeOutTradeNoUsed = "OUT_TRADE_NO_USED" // 商户订单号重复
	ErrorCodeOrderNotExist = "ORDER_NO_TEXIST" // 订单不存在
	ErrorCodeOrderClosed = "ORDER_CLOSED" // 订单已关闭
	ErrorCodeOpenIDMismatch = "OPENID_MISMATCH" // openid和appid不匹配
	ErrorCodeNotEnough = "NOT_ENOUGH" // 余额不足
	ErrorCodeNOAuth = "NO_AUTH" // 商户无权限
	ErrorCodeMchNotExists = "MCH_NOT_EXISTS" // 商户号不存在
	ErrorCodeInvalidTransactionID = "INVALID_TRANSACTIONID" // 订单号非法
	ErrorCodeInvalidRequest = "INVALID_REQUEST" // 无效请求
	ErrorCodeFrequencyLimited = "FREQUENCY_LIMITED" // 频率超限
	ErrorCodeBankError = "BANK_ERROR" // 银行系统异常
	ErrorCodeAppidMchIDNotMatch = "APPID_MCHID_NOT_MATCH" // appid和mch_id不匹配
	ErrorCodeAccount = "ACCOUNT_ERROR" // 账号异常
)

type WxError struct {
	Code int64
	Message string
}

func NewWxError(code int64, message string) *WxError {
	return &WxError{
		Code:    code,
		Message: message,
	}
}

func (e *WxError) Error() string {
	return fmt.Sprintf("code[%d], message[%s]", e.Code, e.Message)
}

var (
	Errors = map[string] error {
		ErrorCodeTrade: NewWxError(0, ErrorCodeTrade),
		ErrorCodeSystem: NewWxError(0, ErrorCodeSystem),
		ErrorCodeSign: NewWxError(0, ErrorCodeSign),
	}
)

type Amount struct {
	Currency string `json:"currency"`
	Total int32 `json:"amount"`
	PayerCurrency string `json:"payer_currency,omitempty"`
	PayerTotal int32 `json:"payer_total,omitempty"`
}

type Payer struct {
	OpenID string `json:"openid"`
}

type SceneInfo struct {
}

type SettleInfo struct {

}