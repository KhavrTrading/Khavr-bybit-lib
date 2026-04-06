package bybit

import (
	"fmt"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// GetPositions queries open positions via GET /v5/position/list.
func (c *BybitClient) GetPositions(
	req *bybit_models.PositionListRequest,
) (*bybit_models.PositionListResponse, error) {
	endpoint := "/v5/position/list"
	var resp bybit_models.PositionListResponse

	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("get positions failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
