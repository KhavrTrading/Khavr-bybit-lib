package bybit_models

import (
	"strconv"
)

// ClosedPnlRequest represents the query parameters for GET /v5/position/closed-pnl
type ClosedPnlRequest struct {
	Category  string // required: "linear", "inverse", or "option"
	Symbol    string // optional: e.g. "BTCUSDT"
	StartTime int64  // optional: ms since epoch. Default behavior: 7-day window
	EndTime   int64  // optional: ms since epoch
	Limit     int    // optional: [1,100], default 50
	Cursor    string // optional: nextPageCursor token
}

// ToParams converts the request into a map of string parameters
func (r *ClosedPnlRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.Category != "" {
		params["category"] = r.Category
	}
	if r.Symbol != "" {
		params["symbol"] = r.Symbol
	}
	if r.StartTime > 0 {
		params["startTime"] = strconv.FormatInt(r.StartTime, 10)
	}
	if r.EndTime > 0 {
		params["endTime"] = strconv.FormatInt(r.EndTime, 10)
	}
	if r.Limit > 0 {
		if r.Limit < 1 {
			r.Limit = 1
		} else if r.Limit > 100 {
			r.Limit = 100
		}
		params["limit"] = strconv.Itoa(r.Limit)
	}
	if r.Cursor != "" {
		params["cursor"] = r.Cursor
	}
	return params
}

// ClosedPnlResponse wraps the full API response
type ClosedPnlResponse struct {
	RetCode    int             `json:"retCode"`
	RetMsg     string          `json:"retMsg"`
	Result     ClosedPnlResult `json:"result"`
	RetExtInfo map[string]any  `json:"retExtInfo"`
	Time       int64           `json:"time"`
}

// ClosedPnlResult holds paging info and the list of entries
type ClosedPnlResult struct {
	NextPageCursor string           `json:"nextPageCursor"`
	Category       string           `json:"category"`
	List           []ClosedPnlEntry `json:"list"`
}

// ClosedPnlEntry describes one closed PnL record
type ClosedPnlEntry struct {
	Symbol        string `json:"symbol"`
	OrderID       string `json:"orderId"`
	Side          string `json:"side"`
	Qty           string `json:"qty"`
	OrderPrice    string `json:"orderPrice"`
	OrderType     string `json:"orderType"`
	ExecType      string `json:"execType"`
	ClosedSize    string `json:"closedSize"`
	CumEntryValue string `json:"cumEntryValue"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	CumExitValue  string `json:"cumExitValue"`
	AvgExitPrice  string `json:"avgExitPrice"`
	ClosedPnl     string `json:"closedPnl"`
	FillCount     string `json:"fillCount"`
	Leverage      string `json:"leverage"`
	OpenFee       string `json:"openFee"`  // New field
	CloseFee      string `json:"closeFee"` // New field
	CreatedTime   string `json:"createdTime"`
	UpdatedTime   string `json:"updatedTime"`
}
