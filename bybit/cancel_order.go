package bybit

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// CancelOrder cancels an unfilled or partially filled order via POST /v5/order/cancel.
func (c *BybitClient) CancelOrder(
	req *bybit_models.CancelOrderRequest,
) (*bybit_models.CancelOrderResponse, error) {
	endpoint := "/v5/order/cancel"
	var resp bybit_models.CancelOrderResponse

	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("cancel order failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
