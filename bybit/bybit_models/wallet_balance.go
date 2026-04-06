package bybit_models

type WalletBalanceRequest struct {
	AccountType string // e.g., "UNIFIED"
	Coin        string // e.g., "BTC,USDT"
}

type WalletBalanceResponse struct {
	RetCode    int               `json:"retCode"`
	RetMsg     string            `json:"retMsg"`
	Result     WalletBalanceData `json:"result"`
	RetExtInfo map[string]any    `json:"retExtInfo"`
	Time       int64             `json:"time"`
}

type WalletBalanceData struct {
	List []WalletAccount `json:"list"`
}

type WalletAccount struct {
	AccountType            string       `json:"accountType"`
	TotalEquity            string       `json:"totalEquity"`
	TotalWalletBalance     string       `json:"totalWalletBalance"`
	TotalMarginBalance     string       `json:"totalMarginBalance"`
	TotalAvailableBalance  string       `json:"totalAvailableBalance"`
	TotalPerpUPL           string       `json:"totalPerpUPL"`
	TotalInitialMargin     string       `json:"totalInitialMargin"`
	TotalMaintenanceMargin string       `json:"totalMaintenanceMargin"`
	AccountIMRate          string       `json:"accountIMRate"`
	AccountMMRate          string       `json:"accountMMRate"`
	AccountLTV             string       `json:"accountLTV"`
	Coin                   []WalletCoin `json:"coin"`
}

type WalletCoin struct {
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
	MarginCollateral    bool   `json:"marginCollateral"`
	CollateralSwitch    bool   `json:"collateralSwitch"`
}

func (r *WalletBalanceRequest) ToParams() map[string]string {
	params := make(map[string]string)
	if r.AccountType != "" {
		params["accountType"] = r.AccountType
	}
	if r.Coin != "" {
		params["coin"] = r.Coin
	}
	return params
}
