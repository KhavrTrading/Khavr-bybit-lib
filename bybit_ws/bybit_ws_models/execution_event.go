package bybit_ws_models

// ExecutionEvent is pushed when a trade is executed.
// Topics: "execution" (all-in-one) or "execution.linear", "execution.spot", etc.
type ExecutionEvent struct {
	ID           string          `json:"id"`
	Topic        string          `json:"topic"`
	CreationTime int64           `json:"creationTime"`
	Data         []ExecutionData `json:"data"`
}

// ExecutionData contains details of a single execution (fill).
type ExecutionData struct {
	Category        string `json:"category"`
	Symbol          string `json:"symbol"`
	IsLeverage      string `json:"isLeverage"`
	OrderId         string `json:"orderId"`
	OrderLinkId     string `json:"orderLinkId"`
	Side            string `json:"side"`
	OrderPrice      string `json:"orderPrice"`
	OrderQty        string `json:"orderQty"`
	LeavesQty       string `json:"leavesQty"`
	CreateType      string `json:"createType"`
	OrderType       string `json:"orderType"`
	StopOrderType   string `json:"stopOrderType"`
	ExecFee         string `json:"execFee"`
	ExecId          string `json:"execId"`
	ExecPrice       string `json:"execPrice"`
	ExecQty         string `json:"execQty"`
	ExecPnl         string `json:"execPnl"`
	ExecType        string `json:"execType"`
	ExecValue       string `json:"execValue"`
	ExecTime        string `json:"execTime"`
	IsMaker         bool   `json:"isMaker"`
	FeeRate         string `json:"feeRate"`
	TradeIv         string `json:"tradeIv"`
	MarkIv          string `json:"markIv"`
	MarkPrice       string `json:"markPrice"`
	IndexPrice      string `json:"indexPrice"`
	UnderlyingPrice string `json:"underlyingPrice"`
	BlockTradeId    string `json:"blockTradeId"`
	ClosedSize      string `json:"closedSize"`
	Seq             int64  `json:"seq"`
	MarketUnit      string `json:"marketUnit"`
}

// FastExecutionEvent is the low-latency execution stream.
// Topics: "execution.fast" (all-in-one) or "execution.fast.linear", etc.
type FastExecutionEvent struct {
	ID           string              `json:"id"`
	Topic        string              `json:"topic"`
	CreationTime int64               `json:"creationTime"`
	Data         []FastExecutionData `json:"data"`
}

// FastExecutionData is a minimal execution record for low-latency use.
type FastExecutionData struct {
	Category      string `json:"category"`
	Symbol        string `json:"symbol"`
	OrderId       string `json:"orderId"`
	OrderLinkId   string `json:"orderLinkId"`
	Side          string `json:"side"`
	OrderPrice    string `json:"orderPrice"`
	OrderQty      string `json:"orderQty"`
	LeavesQty     string `json:"leavesQty"`
	OrderType     string `json:"orderType"`
	StopOrderType string `json:"stopOrderType"`
	ExecFee       string `json:"execFee"`
	ExecId        string `json:"execId"`
	ExecPrice     string `json:"execPrice"`
	ExecQty       string `json:"execQty"`
	ExecType      string `json:"execType"`
	ExecValue     string `json:"execValue"`
	ExecTime      string `json:"execTime"`
	IsMaker       bool   `json:"isMaker"`
	BlockTradeId  string `json:"blockTradeId"`
	Seq           int64  `json:"seq"`
}
