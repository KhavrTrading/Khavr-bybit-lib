// bybit_models/trading_stop.go
package bybit_models

// TradingStopRequest defines the body for POST /v5/position/trading-stop.
type TradingStopRequest struct {
	Category     string `json:"category"`               // "linear", "inverse"
	Symbol       string `json:"symbol"`                 // e.g. "BTCUSDT"
	TpslMode     string `json:"tpslMode"`               // "Full" or "Partial"
	PositionIdx  int    `json:"positionIdx"`            // 0: one-way, 1: hedge Buy, 2: hedge Sell
	TakeProfit   string `json:"takeProfit,omitempty"`   // price, "0" to cancel TP
	StopLoss     string `json:"stopLoss,omitempty"`     // price, "0" to cancel SL
	TrailingStop string `json:"trailingStop,omitempty"` // distance, "0" to cancel TS
	TpTriggerBy  string `json:"tpTriggerBy,omitempty"`  // "LastPrice", "IndexPrice", "MarkPrice"
	SlTriggerBy  string `json:"slTriggerBy,omitempty"`  // same as above
	ActivePrice  string `json:"activePrice,omitempty"`  // trigger price for TS
	TpSize       string `json:"tpSize,omitempty"`       // qty for partial TP
	SlSize       string `json:"slSize,omitempty"`       // qty for partial SL (equal to tpSize)
	TpLimitPrice string `json:"tpLimitPrice,omitempty"` // limit price for partial TP
	SlLimitPrice string `json:"slLimitPrice,omitempty"` // limit price for partial SL
	TpOrderType  string `json:"tpOrderType,omitempty"`  // "Market" or "Limit"
	SlOrderType  string `json:"slOrderType,omitempty"`  // "Market" or "Limit"
}

// TradingStopResponse wraps the response for POST /v5/position/trading-stop.
type TradingStopResponse struct {
	RetCode    int            `json:"retCode"` // 0 indicates success
	RetMsg     string         `json:"retMsg"`
	RetExtInfo map[string]any `json:"retExtInfo"` // any extra info
	Time       int64          `json:"time"`       // server timestamp
}
