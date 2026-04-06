# Khavr-bybit-lib

Go client library for the Bybit REST API.

## Installation

```bash
go get github.com/KhavrTrading/Khavr-bybit-lib
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	bybitlib "github.com/KhavrTrading/Khavr-bybit-lib"
)

func main() {
	client := bybitlib.NewBybitClient("your-api-key", "your-api-secret")

	// Place an order
	resp, err := client.CreateOrder(&bybitlib.CreateOrderRequest{
		Category:  "linear",
		Symbol:    "BTCUSDT",
		Side:      "Buy",
		OrderType: "Limit",
		Qty:       "0.001",
		Price:     "50000",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Order ID:", resp.Result.OrderId)
}
```

## Supported Endpoints

| Method | Endpoint | Function |
|--------|----------|----------|
| POST | `/v5/order/create` | `CreateOrder` |
| POST | `/v5/order/amend` | `AmendOrder` |
| POST | `/v5/account/set-margin-mode` | `SetMarginMode` |
| POST | `/v5/position/trading-stop` | `SetTradingStop` |
| GET | `/v5/position/list` | `GetPositions` |
| GET | `/v5/position/closed-pnl` | `GetClosedPnL` |
| GET | `/v5/position/get-closed-positions` | `GetClosedOptionsPositions` |
| GET | `/v5/account/wallet-balance` | `GetWalletBalance` |
| GET | `/v5/order/realtime` | `GetRealtimeOrders` |

## License

[MIT](LICENSE)
