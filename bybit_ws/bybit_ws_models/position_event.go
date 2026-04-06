package bybit_ws_models

// PositionEvent is pushed when a position changes.
// Topics: "position" (all-in-one) or "position.linear", "position.inverse", etc.
type PositionEvent struct {
	ID           string         `json:"id"`
	Topic        string         `json:"topic"`
	CreationTime int64          `json:"creationTime"`
	Data         []PositionData `json:"data"`
}

// PositionData contains details of a single position update.
type PositionData struct {
	Category               string `json:"category"`
	Symbol                 string `json:"symbol"`
	Side                   string `json:"side"`
	Size                   string `json:"size"`
	PositionIdx            int    `json:"positionIdx"`
	TradeMode              int    `json:"tradeMode"`
	PositionValue          string `json:"positionValue"`
	RiskId                 int    `json:"riskId"`
	RiskLimitValue         string `json:"riskLimitValue"`
	EntryPrice             string `json:"entryPrice"`
	MarkPrice              string `json:"markPrice"`
	Leverage               string `json:"leverage"`
	PositionBalance        string `json:"positionBalance"`
	AutoAddMargin          int    `json:"autoAddMargin"`
	PositionIM             string `json:"positionIM"`
	PositionMM             string `json:"positionMM"`
	LiqPrice               string `json:"liqPrice"`
	BustPrice              string `json:"bustPrice"`
	TpslMode               string `json:"tpslMode"`
	TakeProfit             string `json:"takeProfit"`
	StopLoss               string `json:"stopLoss"`
	TrailingStop           string `json:"trailingStop"`
	UnrealisedPnl          string `json:"unrealisedPnl"`
	CurRealisedPnl         string `json:"curRealisedPnl"`
	SessionAvgPrice        string `json:"sessionAvgPrice"`
	Delta                  string `json:"delta"`
	Gamma                  string `json:"gamma"`
	Vega                   string `json:"vega"`
	Theta                  string `json:"theta"`
	CumRealisedPnl         string `json:"cumRealisedPnl"`
	PositionStatus         string `json:"positionStatus"`
	AdlRankIndicator       int    `json:"adlRankIndicator"`
	IsReduceOnly           bool   `json:"isReduceOnly"`
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"`
	CreatedTime            string `json:"createdTime"`
	UpdatedTime            string `json:"updatedTime"`
	Seq                    int64  `json:"seq"`
}
