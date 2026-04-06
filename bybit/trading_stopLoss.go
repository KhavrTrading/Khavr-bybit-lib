// bybit/client_trading_stop.go
package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// SetTradingStop sends a TP/SL/TS request to Bybit.
func (c *BybitClient) SetTradingStop(
	req *bybit_models.TradingStopRequest,
) (*bybit_models.TradingStopResponse, error) {
	endpoint := "/v5/position/trading-stop"
	var resp bybit_models.TradingStopResponse

	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("set trading stop failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
