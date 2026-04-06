package bybit_models

import "fmt"

// SpreadInstrumentRequest defines query parameters for GET /v5/spread/instrument.
type SpreadInstrumentRequest struct {
	Symbol   string // optional: spread symbol name
	BaseCoin string // optional: base coin, uppercase
	Limit    int    // optional: [1,500], default 200
	Cursor   string // optional: pagination token
}

func (r *SpreadInstrumentRequest) ToParams() map[string]string {
	p := make(map[string]string)
	if r.Symbol != "" {
		p["symbol"] = r.Symbol
	}
	if r.BaseCoin != "" {
		p["baseCoin"] = r.BaseCoin
	}
	if r.Limit > 0 {
		p["limit"] = fmt.Sprint(r.Limit)
	}
	if r.Cursor != "" {
		p["cursor"] = r.Cursor
	}
	return p
}

// SpreadInstrumentResponse wraps the response from GET /v5/spread/instrument.
type SpreadInstrumentResponse struct {
	RetCode    int                    `json:"retCode"`
	RetMsg     string                 `json:"retMsg"`
	Result     SpreadInstrumentResult `json:"result"`
	RetExtInfo map[string]any         `json:"retExtInfo"`
	Time       int64                  `json:"time"`
}

type SpreadInstrumentResult struct {
	List           []SpreadInstrumentItem `json:"list"`
	NextPageCursor string                 `json:"nextPageCursor"`
}

// SpreadInstrumentItem represents a single spread instrument.
type SpreadInstrumentItem struct {
	Symbol       string      `json:"symbol"`
	ContractType string      `json:"contractType"` // "FundingRateArb", "CarryTrade", "FutureSpread", "PerpBasis"
	Status       string      `json:"status"`       // "Trading", "Settling"
	BaseCoin     string      `json:"baseCoin"`
	QuoteCoin    string      `json:"quoteCoin"`
	SettleCoin   string      `json:"settleCoin"`
	TickSize     string      `json:"tickSize"`
	MinPrice     string      `json:"minPrice"`
	MaxPrice     string      `json:"maxPrice"`
	LotSize      string      `json:"lotSize"`
	MinSize      string      `json:"minSize"`
	MaxSize      string      `json:"maxSize"`
	LaunchTime   string      `json:"launchTime"`   // ms timestamp
	DeliveryTime string      `json:"deliveryTime"` // ms timestamp
	Legs         []SpreadLeg `json:"legs"`
}

// SpreadLeg represents a single leg within a spread instrument.
type SpreadLeg struct {
	Symbol       string `json:"symbol"`
	ContractType string `json:"contractType"` // "LinearPerpetual", "LinearFutures", "Spot"
}
