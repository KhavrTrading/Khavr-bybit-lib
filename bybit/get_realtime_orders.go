package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// GetRealtimeOrders calls GET /v5/order/realtime with the provided parameters
// and returns the raw Bybit response or an error.
func (c *BybitClient) GetRealtimeOrders(
	req *bybit_models.OrderRealtimeRequest,
) (*bybit_models.OrderRealtimeResponse, error) {
	endpoint := "/v5/order/realtime"
	var resp bybit_models.OrderRealtimeResponse

	// Execute HTTP GET with query params
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	// Check API-level success code
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("get realtime orders failed: %s", resp.RetMsg)
	}

	return &resp, nil
}
