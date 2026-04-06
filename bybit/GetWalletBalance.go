// bybit/get_wallet_balance.go
package bybit

import (
	"fmt"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
)

func (c *BybitClient) GetWalletBalance(req *bybit_models.WalletBalanceRequest) (*bybit_models.WalletBalanceResponse, error) {
	endpoint := "/v5/account/wallet-balance"

	var resp bybit_models.WalletBalanceResponse
	if err := c.doGet(endpoint, req, &resp); err != nil {
		return nil, err
	}

	if resp.RetCode != 0 {
		return nil, fmt.Errorf("get wallet balance failed: %s", resp.RetMsg)
	}

	return &resp, nil
}
