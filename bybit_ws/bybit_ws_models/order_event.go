package bybit_ws_models

// OrderEvent is pushed when an order is created, updated, or filled.
// Topics: "order" (all-in-one) or "order.linear", "order.spot", etc.
type OrderEvent struct {
	ID           string      `json:"id"`
	Topic        string      `json:"topic"`
	CreationTime int64       `json:"creationTime"`
	Data         []OrderData `json:"data"`
}

// OrderData contains details of a single order update.
type OrderData struct {
	Category            string `json:"category"`
	OrderId             string `json:"orderId"`
	OrderLinkId         string `json:"orderLinkId"`
	IsLeverage          string `json:"isLeverage"`
	BlockTradeId        string `json:"blockTradeId"`
	Symbol              string `json:"symbol"`
	Price               string `json:"price"`
	Qty                 string `json:"qty"`
	Side                string `json:"side"`
	PositionIdx         int    `json:"positionIdx"`
	OrderStatus         string `json:"orderStatus"`
	CreateType          string `json:"createType"`
	CancelType          string `json:"cancelType"`
	RejectReason        string `json:"rejectReason"`
	AvgPrice            string `json:"avgPrice"`
	LeavesQty           string `json:"leavesQty"`
	LeavesValue         string `json:"leavesValue"`
	CumExecQty          string `json:"cumExecQty"`
	CumExecValue        string `json:"cumExecValue"`
	CumExecFee          string `json:"cumExecFee"`
	ClosedPnl           string `json:"closedPnl"`
	FeeCurrency         string `json:"feeCurrency"`
	TimeInForce         string `json:"timeInForce"`
	OrderType           string `json:"orderType"`
	StopOrderType       string `json:"stopOrderType"`
	OcoTriggerBy        string `json:"ocoTriggerBy"`
	OrderIv             string `json:"orderIv"`
	MarketUnit          string `json:"marketUnit"`
	TriggerPrice        string `json:"triggerPrice"`
	TakeProfit          string `json:"takeProfit"`
	StopLoss            string `json:"stopLoss"`
	TpslMode            string `json:"tpslMode"`
	TpLimitPrice        string `json:"tpLimitPrice"`
	SlLimitPrice        string `json:"slLimitPrice"`
	TpTriggerBy         string `json:"tpTriggerBy"`
	SlTriggerBy         string `json:"slTriggerBy"`
	TriggerDirection    int    `json:"triggerDirection"`
	TriggerBy           string `json:"triggerBy"`
	LastPriceOnCreated  string `json:"lastPriceOnCreated"`
	ReduceOnly          bool   `json:"reduceOnly"`
	CloseOnTrigger      bool   `json:"closeOnTrigger"`
	PlaceType           string `json:"placeType"`
	SmpType             string `json:"smpType"`
	SmpGroup            int    `json:"smpGroup"`
	SmpOrderId          string `json:"smpOrderId"`
	CreatedTime         string `json:"createdTime"`
	UpdatedTime         string `json:"updatedTime"`
}
