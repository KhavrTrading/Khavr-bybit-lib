package bybit_models

import (
	"strconv"
)

// ClosedOptionsPositionsRequest represents the query parameters for GET /v5/position/get-closed-positions
type ClosedOptionsPositionsRequest struct {
	Category  string // required: e.g. "option"
	Symbol    string // optional: e.g. "BTC-12JUN25-104019-C-USDT"
	StartTime int64  // optional: ms since epoch
	EndTime   int64  // optional: ms since epoch
	Limit     int    // optional: [1,100], default 50
	Cursor    string // optional: nextPageCursor token
}

// ToParams converts the request into a map of string parameters
func (r *ClosedOptionsPositionsRequest) ToParams() map[string]string {
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

// ClosedOptionsPositionsResponse wraps the full API response
type ClosedOptionsPositionsResponse struct {
	RetCode    int                          `json:"retCode"`
	RetMsg     string                       `json:"retMsg"`
	Result     ClosedOptionsPositionsResult `json:"result"`
	RetExtInfo map[string]any               `json:"retExtInfo"`
	Time       int64                        `json:"time"`
}

// ClosedOptionsPositionsResult holds paging info and the list of closed option positions
type ClosedOptionsPositionsResult struct {
	NextPageCursor string                       `json:"nextPageCursor"`
	Category       string                       `json:"category"`
	List           []ClosedOptionsPositionEntry `json:"list"`
}

// ClosedOptionsPositionEntry describes one closed options position record
type ClosedOptionsPositionEntry struct {
	Symbol        string `json:"symbol"`
	Side          string `json:"side"`
	TotalOpenFee  string `json:"totalOpenFee"`
	DeliveryFee   string `json:"deliveryFee"`
	TotalCloseFee string `json:"totalCloseFee"`
	Qty           string `json:"qty"`
	CloseTime     int64  `json:"closeTime"`
	AvgExitPrice  string `json:"avgExitPrice"`
	DeliveryPrice string `json:"deliveryPrice"`
	OpenTime      int64  `json:"openTime"`
	AvgEntryPrice string `json:"avgEntryPrice"`
	TotalPnl      string `json:"totalPnl"`
}
