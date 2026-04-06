package bybit_models

const (
	// Categories (product types)
	CategoryLinear  = "linear"
	CategoryInverse = "inverse"
	CategorySpot    = "spot"
	CategoryOption  = "option"

	// Margin modes
	MarginModeIsolated  = "ISOLATED_MARGIN"
	MarginModeRegular   = "REGULAR_MARGIN"
	MarginModePortfolio = "PORTFOLIO_MARGIN"
	MarginModeCross     = "CROSS_MARGIN"

	// Trade sides
	SideBuy  = "Buy"
	SideSell = "Sell"

	// Order types
	OrderTypeLimit  = "Limit"
	OrderTypeMarket = "Market"

	// Time in force
	TimeInForceGTC      = "GTC"
	TimeInForceIOC      = "IOC"
	TimeInForceFOK      = "FOK"
	TimeInForcePostOnly = "PostOnly"

	// Position index (hedge mode)
	PositionIdxOneWay  = 0
	PositionIdxBuySide = 1
	PositionIdxSellSide = 2

	// Trigger price types
	TriggerByLastPrice  = "LastPrice"
	TriggerByIndexPrice = "IndexPrice"
	TriggerByMarkPrice  = "MarkPrice"

	// TP/SL modes
	TpslModeFull    = "Full"
	TpslModePartial = "Partial"

	// Account types
	AccountTypeUnified  = "UNIFIED"
	AccountTypeContract = "CONTRACT"

	// Order status
	OrderStatusNew             = "New"
	OrderStatusPartiallyFilled = "PartiallyFilled"
	OrderStatusFilled          = "Filled"
	OrderStatusCancelled       = "Cancelled"
	OrderStatusRejected        = "Rejected"
	OrderStatusDeactivated     = "Deactivated"

	// STP modes
	STPModeNone        = "none"
	STPModeCancelMaker = "CancelMaker"
	STPModeCancelTaker = "CancelTaker"
	STPModeCancelBoth  = "CancelBoth"
)
