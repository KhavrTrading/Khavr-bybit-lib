package bybit_models

// CancelOrderRequest defines the body for POST /v5/order/cancel.
// Either OrderId or OrderLinkId must be provided.
type CancelOrderRequest struct {
	Category    string `json:"category"`              // "linear", "inverse", "spot", "option"
	Symbol      string `json:"symbol"`                // e.g. "BTCUSDT"
	OrderId     string `json:"orderId,omitempty"`     // Bybit order ID
	OrderLinkId string `json:"orderLinkId,omitempty"` // Client-generated ID
	OrderFilter string `json:"orderFilter,omitempty"` // Spot only: "Order", "tpslOrder", "StopOrder"
}

// CancelOrderResponse wraps the response from POST /v5/order/cancel.
type CancelOrderResponse struct {
	RetCode int    `json:"retCode"` // 0 indicates success
	RetMsg  string `json:"retMsg"`
	Result  struct {
		OrderId     string `json:"orderId"`
		OrderLinkId string `json:"orderLinkId"`
	} `json:"result"`
	RetExtInfo map[string]any `json:"retExtInfo"`
	Time       int64          `json:"time"`
}
