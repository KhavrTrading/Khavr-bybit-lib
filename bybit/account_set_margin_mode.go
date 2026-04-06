// bybit/account_set_margin_mode.go
package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// SetMarginMode switches the account's margin mode.
// mode must be one of: "ISOLATED_MARGIN", "REGULAR_MARGIN", or "PORTFOLIO_MARGIN".
// Returns nil on success, or an error if the API call fails or returns a non-zero code.
func (c *BybitClient) SetMarginMode(mode string) error {
	endpoint := "/v5/account/set-margin-mode"
	req := &bybit_models.SetMarginModeRequest{
		SetMarginMode: mode,
	}
	var resp bybit_models.SetMarginModeResponse

	// Execute HTTP POST and unmarshal into resp
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return err
	}

	// Check API-level success code
	if resp.RetCode != 0 {
		return fmt.Errorf("set margin mode failed: %s", resp.RetMsg)
	}

	// Optionally inspect resp.Reasons for details if needed
	return nil
}
