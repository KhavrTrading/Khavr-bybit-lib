package bybit_models

import "fmt"

// OrderRealtimeRequest defines query parameters for GET /v5/order/realtime
type OrderRealtimeRequest struct {
	Category    string `url:"category"`             // "linear", "inverse", "spot", "option"
	Symbol      string `url:"symbol,omitempty"`     // trading pair, e.g. "BTCUSDT"
	BaseCoin    string `url:"baseCoin,omitempty"`   // for linear: symbol, baseCoin, or settleCoin required
	SettleCoin  string `url:"settleCoin,omitempty"` // USDT or USDC for option
	OrderId     string `url:"orderId,omitempty"`
	OrderLinkId string `url:"orderLinkId,omitempty"`
	OpenOnly    *int   `url:"openOnly,omitempty"` // 0 = open only, others per API doc
	OrderFilter string `url:"orderFilter,omitempty"`
	Limit       *int   `url:"limit,omitempty"` // page size [1,50]
	Cursor      string `url:"cursor,omitempty"`
}

// OrderRealtimeResponse wraps GET /v5/order/realtime response
type OrderRealtimeResponse struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category       string              `json:"category"`
		NextPageCursor string              `json:"nextPageCursor"`
		List           []OrderRealtimeItem `json:"list"`
	} `json:"result"`
	Time int64 `json:"time"`
}

// OrderRealtimeItem represents a single order in the realtime list
type OrderRealtimeItem struct {
	OrderId        string `json:"orderId"`
	OrderLinkId    string `json:"orderLinkId"`
	Symbol         string `json:"symbol"`
	Side           string `json:"side"` // Buy or Sell
	OrderType      string `json:"orderType"`
	OrderStatus    string `json:"orderStatus"` // eg. "live", "partially_filled"
	Price          string `json:"price"`
	Qty            string `json:"qty"`
	CumExecQty     string `json:"cumExecQty"`
	AvgPrice       string `json:"avgPrice"`
	CreatedTime    string `json:"createdTime"` // ms timestamp as string
	UpdatedTime    string `json:"updatedTime"`
	ReduceOnly     bool   `json:"reduceOnly"`
	CloseOnTrigger bool   `json:"closeOnTrigger"`
}

func (r *OrderRealtimeRequest) ToParams() map[string]string {
	p := make(map[string]string)
	p["category"] = r.Category
	if r.Symbol != "" {
		p["symbol"] = r.Symbol
	}
	if r.BaseCoin != "" {
		p["baseCoin"] = r.BaseCoin
	}
	if r.SettleCoin != "" {
		p["settleCoin"] = r.SettleCoin
	}
	if r.OrderId != "" {
		p["orderId"] = r.OrderId
	}
	if r.OrderLinkId != "" {
		p["orderLinkId"] = r.OrderLinkId
	}
	if r.OpenOnly != nil {
		p["openOnly"] = fmt.Sprint(*r.OpenOnly)
	}
	if r.OrderFilter != "" {
		p["orderFilter"] = r.OrderFilter
	}
	if r.Limit != nil {
		p["limit"] = fmt.Sprint(*r.Limit)
	}
	if r.Cursor != "" {
		p["cursor"] = r.Cursor
	}
	return p
}
