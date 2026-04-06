package bybit_ws_models

// WalletEvent is pushed when wallet balances change.
// Topic: "wallet"
type WalletEvent struct {
	ID           string       `json:"id"`
	Topic        string       `json:"topic"`
	CreationTime int64        `json:"creationTime"`
	Data         []WalletData `json:"data"`
}

// WalletData contains account-level wallet information.
type WalletData struct {
	AccountType            string         `json:"accountType"`
	AccountLTV             string         `json:"accountLTV"`
	AccountIMRate          string         `json:"accountIMRate"`
	AccountMMRate          string         `json:"accountMMRate"`
	TotalEquity            string         `json:"totalEquity"`
	TotalWalletBalance     string         `json:"totalWalletBalance"`
	TotalMarginBalance     string         `json:"totalMarginBalance"`
	TotalAvailableBalance  string         `json:"totalAvailableBalance"`
	TotalPerpUPL           string         `json:"totalPerpUPL"`
	TotalInitialMargin     string         `json:"totalInitialMargin"`
	TotalMaintenanceMargin string         `json:"totalMaintenanceMargin"`
	Coin                   []WalletCoinWS `json:"coin"`
}

// WalletCoinWS contains per-coin balance details from the wallet stream.
type WalletCoinWS struct {
	Coin                string `json:"coin"`
	Equity              string `json:"equity"`
	UsdValue            string `json:"usdValue"`
	WalletBalance       string `json:"walletBalance"`
	Free                string `json:"free"`
	Locked              string `json:"locked"`
	SpotHedgingQty      string `json:"spotHedgingQty"`
	BorrowAmount        string `json:"borrowAmount"`
	AvailableToWithdraw string `json:"availableToWithdraw"`
	AccruedInterest     string `json:"accruedInterest"`
	TotalOrderIM        string `json:"totalOrderIM"`
	TotalPositionIM     string `json:"totalPositionIM"`
	TotalPositionMM     string `json:"totalPositionMM"`
	UnrealisedPnl       string `json:"unrealisedPnl"`
	CumRealisedPnl      string `json:"cumRealisedPnl"`
	Bonus               string `json:"bonus"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
	MarginCollateral    bool   `json:"marginCollateral"`
}
