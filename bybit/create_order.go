// bybit/create_order.go
package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// CreateOrder sends a new order to Bybit via POST /v5/order/create.
// It returns the API response containing order IDs, or an error.
func (c *BybitClient) CreateOrder(
	req *bybit_models.CreateOrderRequest,
) (*bybit_models.CreateOrderResponse, error) {
	endpoint := "/v5/order/create"
	var resp bybit_models.CreateOrderResponse

	// Execute HTTP POST and unmarshal into resp
	if err := c.doPost(endpoint, req, &resp); err != nil {
		return nil, err
	}

	// Check API-level success code
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("create order failed: %s", resp.RetMsg)
	}

	return &resp, nil
}
