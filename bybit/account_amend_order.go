// bybit/account_amend_order.go
package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// AmendOrder sends a modification request to Bybit via POST /v5/order/amend.
// Returns the raw response or an error if it fails.
func (c *BybitClient) AmendOrder(
	req *bybit_models.AmendOrderRequest,
) (*bybit_models.AmendOrderResponse, error) {
	endpoint := "/v5/order/amend"
	var resp bybit_models.AmendOrderResponse

	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("amend order failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
