package bybit

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// GetClosedPnL queries the user's closed profit & loss records via GET /v5/position/closed-pnl
func (c *BybitClient) GetClosedPnL(
	req *bybit_models.ClosedPnlRequest,
) (*bybit_models.ClosedPnlResponse, error) {
	endpoint := "/v5/position/closed-pnl"
	var resp bybit_models.ClosedPnlResponse

	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("get closed PnL failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
