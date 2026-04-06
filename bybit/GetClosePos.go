package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// GetClosedOptionsPositions queries the user's closed options positions via GET /v5/position/get-closed-positions
func (c *BybitClient) GetClosedOptionsPositions(
	req *bybit_models.ClosedOptionsPositionsRequest,
) (*bybit_models.ClosedOptionsPositionsResponse, error) {
	endpoint := "/v5/position/get-closed-positions"
	var resp bybit_models.ClosedOptionsPositionsResponse

	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("get closed options positions failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
