package bybit_models

import "fmt"

// PositionListRequest defines query parameters for GET /v5/position/list.
type PositionListRequest struct {
	Category   string `url:"category"`             // "linear", "inverse", "option"
	Symbol     string `url:"symbol,omitempty"`     // e.g. "BTCUSDT"
	BaseCoin   string `url:"baseCoin,omitempty"`   // option only
	SettleCoin string `url:"settleCoin,omitempty"` // e.g. "USDT", linear requires either symbol or settleCoin
	Limit      *int   `url:"limit,omitempty"`      // page size [1,200]
	Cursor     string `url:"cursor,omitempty"`     // pagination token
}

// PositionListResponse wraps GET /v5/position/list response
type PositionListResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category       string             `json:"category"`
		NextPageCursor string             `json:"nextPageCursor"`
		List           []PositionListItem `json:"list"`
	} `json:"result"`
	Time int64 `json:"time"`
}

// PositionListItem represents a single open position
type PositionListItem struct {
	PositionIdx            int    `json:"positionIdx"`     // Position idx: 0=One-Way, 1=Buy side hedge, 2=Sell side hedge
	RiskId                 int    `json:"riskId"`          // Risk tier ID
	RiskLimitValue         string `json:"riskLimitValue"`  // Risk limit value
	Symbol                 string `json:"symbol"`          // Trading pair
	Side                   string `json:"side"`            // "Buy" (long) or "Sell" (short)
	Size                   string `json:"size"`            // Position size (always positive)
	AvgPrice               string `json:"avgPrice"`        // Average entry price
	PositionValue          string `json:"positionValue"`   // Position value
	TradeMode              int    `json:"tradeMode"`       // 0=cross-margin, 1=isolated margin
	AutoAddMargin          int    `json:"autoAddMargin"`   // 0=false, 1=true for isolated margin auto-add
	PositionStatus         string `json:"positionStatus"`  // "Normal", "Liq", "Adl"
	Leverage               string `json:"leverage"`        // Leverage multiplier
	MarkPrice              string `json:"markPrice"`       // Current mark price
	LiqPrice               string `json:"liqPrice"`        // Liquidation price
	BustPrice              string `json:"bustPrice"`       // Bankruptcy price
	PositionIM             string `json:"positionIM"`      // Initial margin
	PositionMM             string `json:"positionMM"`      // Maintenance margin
	PositionBalance        string `json:"positionBalance"` // Position margin balance
	TakeProfit             string `json:"takeProfit"`      // Take profit price
	StopLoss               string `json:"stopLoss"`        // Stop loss price
	TrailingStop           string `json:"trailingStop"`    // Trailing stop distance
	SessionAvgPrice        string `json:"sessionAvgPrice"` // USDC session avg price
	Delta                  string `json:"delta"`           // Option greek
	Gamma                  string `json:"gamma"`
	Vega                   string `json:"vega"`
	Theta                  string `json:"theta"`
	UnrealisedPnl          string `json:"unrealisedPnl"`          // Unrealised PnL
	CurRealisedPnl         string `json:"curRealisedPnl"`         // Realised PnL for current position
	CumRealisedPnl         string `json:"cumRealisedPnl"`         // Cumulative realised PnL
	AdlRankIndicator       int    `json:"adlRankIndicator"`       // Auto-deleverage rank
	CreatedTime            string `json:"createdTime"`            // First creation timestamp (ms)
	UpdatedTime            string `json:"updatedTime"`            // Last update timestamp (ms)
	Seq                    int64  `json:"seq"`                    // Sequence for updates
	IsReduceOnly           bool   `json:"isReduceOnly"`           // true if only reduce allowed
	MmrSysUpdatedTime      string `json:"mmrSysUpdatedTime"`      // System MMR update timestamp
	LeverageSysUpdatedTime string `json:"leverageSysUpdatedTime"` // System leverage update timestamp
	TpslMode               string `json:"tpslMode"`               // Deprecated, always "Full"
}

func (r *PositionListRequest) ToParams() map[string]string {
	p := map[string]string{
		"category": r.Category,
	}
	if r.Symbol != "" {
		p["symbol"] = r.Symbol
	}
	if r.BaseCoin != "" {
		p["baseCoin"] = r.BaseCoin
	}
	if r.SettleCoin != "" {
		p["settleCoin"] = r.SettleCoin
	}
	if r.Limit != nil {
		p["limit"] = fmt.Sprint(*r.Limit)
	}
	if r.Cursor != "" {
		p["cursor"] = r.Cursor
	}
	return p
}
