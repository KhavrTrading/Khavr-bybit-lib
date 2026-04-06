package bybit_models

// SetMarginModeRequest defines the body for POST /v5/account/set-margin-mode.
type SetMarginModeRequest struct {
	SetMarginMode string `json:"setMarginMode"` // "ISOLATED_MARGIN", "REGULAR_MARGIN", or "PORTFOLIO_MARGIN"
}

// SetMarginModeResponse wraps Bybit's JSON response to the margin-mode request.
type SetMarginModeResponse struct {
	RetCode int    `json:"retCode"` // 0 indicates success
	RetMsg  string `json:"retMsg"`  // message or error description
	Reasons []struct {
		ReasonCode string `json:"reasonCode"` // failure code, if any
		ReasonMsg  string `json:"reasonMsg"`  // failure message, if any
	} `json:"reasons"`
	Time int64 `json:"time"` // server timestamp
}
