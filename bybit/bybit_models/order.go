package bybit_models

// CreateOrderRequest defines parameters for POST /v5/order/create.
type CreateOrderRequest struct {
	Category              string `json:"category"`                        // Product type: "linear", "inverse", "spot", or "option"
	Symbol                string `json:"symbol"`                          // Trading pair, e.g. "BTCUSDT"
	IsLeverage            *int   `json:"isLeverage,omitempty"`            // 0: spot, 1: margin (unified spot)
	Side                  string `json:"side"`                            // "Buy" or "Sell"
	OrderType             string `json:"orderType"`                       // "Market" or "Limit"
	Qty                   string `json:"qty"`                             // Order size (always by qty for perps/futures)
	MarketUnit            string `json:"marketUnit,omitempty"`            // "baseCoin" or "quoteCoin" for spot market orders
	SlippageToleranceType string `json:"slippageToleranceType,omitempty"` // "TickSize" or "Percent"
	SlippageTolerance     string `json:"slippageTolerance,omitempty"`     // Tolerance value (range depends on type)
	Price                 string `json:"price,omitempty"`                 // Limit price (ignored for market orders)
	TriggerDirection      *int   `json:"triggerDirection,omitempty"`      // 1=rise to trigger, 2=fall to trigger (conditional orders)
	OrderFilter           string `json:"orderFilter,omitempty"`           // "tpslOrder" or "StopOrder" (spot only)
	TriggerPrice          string `json:"triggerPrice,omitempty"`          // Price at which a conditional order triggers
	TriggerBy             string `json:"triggerBy,omitempty"`             // "LastPrice", "IndexPrice", or "MarkPrice"
	OrderIv               string `json:"orderIv,omitempty"`               // Implied volatility (option only)
	TimeInForce           string `json:"timeInForce,omitempty"`           // "GTC", "IOC"
	PositionIdx           *int   `json:"positionIdx,omitempty"`           // 0: one-way, 1: hedge Buy, 2: hedge Sell (hedge mode)
	OrderLinkId           string `json:"orderLinkId,omitempty"`           // Client-generated unique ID (max 36 chars)
	TakeProfit            string `json:"takeProfit,omitempty"`            // Take profit price
	StopLoss              string `json:"stopLoss,omitempty"`              // Stop loss price
	TpTriggerBy           string `json:"tpTriggerBy,omitempty"`           // Trigger type for TP
	SlTriggerBy           string `json:"slTriggerBy,omitempty"`           // Trigger type for SL
	ReduceOnly            bool   `json:"reduceOnly,omitempty"`            // If true, order only reduces position
	CloseOnTrigger        bool   `json:"closeOnTrigger,omitempty"`        // Force-close on trigger
	SmpType               string `json:"smpType,omitempty"`               // SMP execution type
	Mmp                   bool   `json:"mmp,omitempty"`                   // Market maker protection (option only)
	TpslMode              string `json:"tpslMode,omitempty"`              // "Full" or "Partial" TP/SL mode
	TpLimitPrice          string `json:"tpLimitPrice,omitempty"`          // Limit price for partial TP
	SlLimitPrice          string `json:"slLimitPrice,omitempty"`          // Limit price for partial SL
	TpOrderType           string `json:"tpOrderType,omitempty"`           // Order type for TP when Partial
	SlOrderType           string `json:"slOrderType,omitempty"`           // Order type for SL when Partial
}

type CreateOrderResponse struct {
	RetCode int               `json:"retCode"` // 0 indicates success
	RetMsg  string            `json:"retMsg"`
	Result  CreateOrderResult `json:"result"` // Contains order IDs
	Time    int64             `json:"time"`   // Server timestamp
}

// CreateOrderResult holds the identifiers returned by Bybit.
type CreateOrderResult struct {
	OrderId     string `json:"orderId"`     // Bybit-generated order ID
	OrderLinkId string `json:"orderLinkId"` // Your client-generated link ID
}
