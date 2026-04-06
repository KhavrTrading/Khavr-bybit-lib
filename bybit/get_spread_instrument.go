package bybit

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

// GetSpreadInstrument queries spread instrument info via GET /v5/spread/instrument.
func (c *BybitClient) GetSpreadInstrument(
	req *bybit_models.SpreadInstrumentRequest,
) (*bybit_models.SpreadInstrumentResponse, error) {
	endpoint := "/v5/spread/instrument"
	var resp bybit_models.SpreadInstrumentResponse

	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}
	if resp.RetCode != 0 {
		return nil, fmt.Errorf("get spread instrument failed: %s", resp.RetMsg)
	}
	return &resp, nil
}
