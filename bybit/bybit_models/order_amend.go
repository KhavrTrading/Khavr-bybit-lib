package bybit_models

// AmendOrderRequest defines the body for POST /v5/order/amend.
type AmendOrderRequest struct {
	Category     string `json:"category"`              // "linear", "inverse", "spot", "option"
	Symbol       string `json:"symbol"`                // e.g. "BTCUSDT"
	OrderId      string `json:"orderId,omitempty"`     // Bybit order ID
	OrderLinkId  string `json:"orderLinkId,omitempty"` // Client-generated ID
	OrderIv      string `json:"orderIv,omitempty"`     // Implied volatility (option)
	TriggerPrice string `json:"triggerPrice,omitempty"`
	Qty          string `json:"qty,omitempty"`
	Price        string `json:"price,omitempty"`
	TpslMode     string `json:"tpslMode,omitempty"`
	TakeProfit   string `json:"takeProfit,omitempty"` // use "0" to cancel existing
	StopLoss     string `json:"stopLoss,omitempty"`   // use "0" to cancel existing
	TpTriggerBy  string `json:"tpTriggerBy,omitempty"`
	SlTriggerBy  string `json:"slTriggerBy,omitempty"`
	TriggerBy    string `json:"triggerBy,omitempty"`
	TpLimitPrice string `json:"tpLimitPrice,omitempty"`
	SlLimitPrice string `json:"slLimitPrice,omitempty"`
}

// AmendOrderResponse wraps the response from POST /v5/order/amend.
type AmendOrderResponse struct {
	RetCode int    `json:"retCode"` // 0 indicates success
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderId     string `json:"orderId"`
		OrderLinkId string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo map[string]any `json:"retExtInfo"`
	Time       int64          `json:"time"`
}
