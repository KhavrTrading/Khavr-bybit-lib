# Khavr-bybit-lib Documentation

## Installation

```bash
go get github.com/KhavrTrading/Khavr-bybit-lib
```

Import via the top-level package for convenience:

```go
import bybitlib "github.com/KhavrTrading/Khavr-bybit-lib"
```

Or import sub-packages directly:

```go
import (
    "github.com/KhavrTrading/Khavr-bybit-lib/bybit"
    "github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
    "github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws"
    "github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"
)
```

---

## REST Client

### Creating a Client

```go
client := bybitlib.NewBybitClient("apiKey", "apiSecret")
```

### Available Methods

| Method | Endpoint | Description |
|--------|----------|-------------|
| `CreateOrder(req)` | `POST /v5/order/create` | Place a new order |
| `AmendOrder(req)` | `POST /v5/order/amend` | Modify an existing order |
| `SetMarginMode(mode)` | `POST /v5/account/set-margin-mode` | Switch margin mode |
| `SetTradingStop(req)` | `POST /v5/position/trading-stop` | Set TP/SL/trailing stop |
| `GetPositions(req)` | `GET /v5/position/list` | Query open positions |
| `GetClosedPnL(req)` | `GET /v5/position/closed-pnl` | Query closed PnL records |
| `GetClosedOptionsPositions(req)` | `GET /v5/position/get-closed-positions` | Query closed option positions |
| `GetWalletBalance(req)` | `GET /v5/account/wallet-balance` | Query wallet balance |
| `GetRealtimeOrders(req)` | `GET /v5/order/realtime` | Query active/conditional orders |

---

## REST Types

### Order Types

#### `CreateOrderRequest`

```go
req := &bybitlib.CreateOrderRequest{
    Category:    "linear",           // required: "linear", "inverse", "spot", "option"
    Symbol:      "BTCUSDT",          // required
    Side:        "Buy",              // required: "Buy" or "Sell"
    OrderType:   "Limit",            // required: "Limit" or "Market"
    Qty:         "0.001",            // required
    Price:       "50000",            // required for Limit orders
    TimeInForce: "GTC",              // optional: "GTC", "IOC", "FOK", "PostOnly"
    PositionIdx: intPtr(0),          // optional: 0=one-way, 1=buy hedge, 2=sell hedge
    OrderLinkId: "my-custom-id",     // optional: client-generated ID (max 36 chars)
    TakeProfit:  "55000",            // optional
    StopLoss:    "48000",            // optional
    TpslMode:    "Full",             // optional: "Full" or "Partial"
    ReduceOnly:  false,              // optional
}
```

#### `CreateOrderResponse`

```go
resp, err := client.CreateOrder(req)
// resp.RetCode    int    - 0 = success
// resp.RetMsg     string
// resp.Result.OrderId     string - Bybit order ID
// resp.Result.OrderLinkId string - your client ID
// resp.Time       int64
```

#### `AmendOrderRequest`

```go
req := &bybitlib.AmendOrderRequest{
    Category:   "linear",        // required
    Symbol:     "BTCUSDT",       // required
    OrderId:    "order-id",      // required (or OrderLinkId)
    Price:      "51000",         // optional: new price
    Qty:        "0.002",         // optional: new qty
    TakeProfit: "56000",         // optional: new TP ("0" to cancel)
    StopLoss:   "47000",         // optional: new SL ("0" to cancel)
}
```

#### `OrderRealtimeRequest`

```go
req := &bybitlib.OrderRealtimeRequest{
    Category:    "linear",       // required
    Symbol:      "BTCUSDT",      // optional
    OrderId:     "",             // optional: filter by order ID
    OrderLinkId: "",             // optional: filter by client ID
    Limit:       intPtr(50),     // optional: page size [1,50]
    Cursor:      "",             // optional: pagination token
}
```

#### `OrderRealtimeResponse`

```go
resp, err := client.GetRealtimeOrders(req)
// resp.Result.List []OrderRealtimeItem
//   .OrderId, .Symbol, .Side, .OrderType, .OrderStatus
//   .Price, .Qty, .CumExecQty, .AvgPrice
//   .CreatedTime, .UpdatedTime
// resp.Result.NextPageCursor string
```

### Position Types

#### `PositionListRequest`

```go
req := &bybitlib.PositionListRequest{
    Category:   "linear",        // required: "linear", "inverse", "option"
    Symbol:     "BTCUSDT",       // optional
    SettleCoin: "USDT",          // optional
    Limit:      intPtr(200),     // optional: [1,200]
    Cursor:     "",              // optional
}
```

#### `PositionListItem` (fields in response)

```go
resp, err := client.GetPositions(req)
for _, p := range resp.Result.List {
    p.Symbol             // "BTCUSDT"
    p.Side               // "Buy" or "Sell"
    p.Size               // position size
    p.AvgPrice           // average entry price
    p.MarkPrice          // current mark price
    p.Leverage           // leverage multiplier
    p.UnrealisedPnl      // unrealised PnL
    p.CumRealisedPnl     // cumulative realised PnL
    p.LiqPrice           // liquidation price
    p.TakeProfit         // TP price
    p.StopLoss           // SL price
    p.TrailingStop       // trailing stop distance
    p.PositionIM         // initial margin
    p.PositionMM         // maintenance margin
    p.PositionStatus     // "Normal", "Liq", "Adl"
    p.PositionIdx        // 0=one-way, 1=buy hedge, 2=sell hedge
    p.TradeMode          // 0=cross, 1=isolated
    p.CreatedTime        // timestamp ms
    p.UpdatedTime        // timestamp ms
}
```

#### `TradingStopRequest`

```go
req := &bybitlib.TradingStopRequest{
    Category:    "linear",       // required
    Symbol:      "BTCUSDT",      // required
    TpslMode:    "Full",         // required: "Full" or "Partial"
    PositionIdx: 0,              // required: 0, 1, or 2
    TakeProfit:  "55000",        // optional ("0" to cancel)
    StopLoss:    "48000",        // optional ("0" to cancel)
    TrailingStop:"200",          // optional ("0" to cancel)
    TpTriggerBy: "LastPrice",    // optional: trigger type
    SlTriggerBy: "LastPrice",    // optional: trigger type
}
```

### Account Types

#### `WalletBalanceRequest`

```go
req := &bybitlib.WalletBalanceRequest{
    AccountType: "UNIFIED",      // "UNIFIED" or "CONTRACT"
    Coin:        "BTC,USDT",     // optional: comma-separated
}
```

#### `WalletAccount` / `WalletCoin` (fields in response)

```go
resp, err := client.GetWalletBalance(req)
for _, acct := range resp.Result.List {
    acct.AccountType            // "UNIFIED"
    acct.TotalEquity            // total equity USD
    acct.TotalWalletBalance     // total wallet balance
    acct.TotalAvailableBalance  // available balance
    acct.TotalPerpUPL           // perps unrealised PnL
    acct.TotalInitialMargin     // total IM
    acct.TotalMaintenanceMargin // total MM

    for _, c := range acct.Coin {
        c.Coin                  // "USDT"
        c.WalletBalance         // wallet balance
        c.Equity                // equity
        c.Free                  // available (spot)
        c.Locked                // locked
        c.UnrealisedPnl         // unrealised PnL
        c.CumRealisedPnl        // cumulative realised PnL
        c.BorrowAmount          // borrowed amount
        c.AvailableToWithdraw   // withdrawable
    }
}
```

#### `ClosedPnlRequest`

```go
req := &bybitlib.ClosedPnlRequest{
    Category:  "linear",         // required
    Symbol:    "BTCUSDT",        // optional
    StartTime: 0,                // optional: ms since epoch
    EndTime:   0,                // optional: ms since epoch
    Limit:     50,               // optional: [1,100]
    Cursor:    "",               // optional
}
```

#### `SetMarginModeRequest`

```go
err := client.SetMarginMode("ISOLATED_MARGIN")
// Modes: "ISOLATED_MARGIN", "REGULAR_MARGIN", "PORTFOLIO_MARGIN", "CROSS_MARGIN"
```

---

## WebSocket Client (Private)

### Creating and Connecting

```go
ws := bybitlib.NewBybitWsClient("apiKey", "apiSecret")

if err := ws.Connect(); err != nil {
    log.Fatal(err)
}
defer ws.Stop()

// Start listening in a goroutine
go ws.ListenLoop()
```

### Channel Managers (Pub/Sub Pattern)

Channel managers distribute WebSocket events to multiple subscribers using buffered Go channels (size 10). If a subscriber's channel is full, the oldest message is drained.

```go
// Create channel managers (wires callbacks automatically)
orderCh    := bybitlib.NewOrderChannel(ws)
execCh     := bybitlib.NewExecutionChannel(ws)
fastExecCh := bybitlib.NewFastExecutionChannel(ws)
positionCh := bybitlib.NewPositionChannel(ws)
walletCh   := bybitlib.NewWalletChannel(ws)

// Add subscribers (returns a buffered chan)
orderSub := orderCh.Subscribe()
posSub   := positionCh.Subscribe()
walletSub := walletCh.Subscribe()

// Subscribe to Bybit topics
orderCh.SubscribeToOrders("order.linear")
positionCh.SubscribeToPositions("position.linear")
walletCh.SubscribeToWallet()

// Read events
go func() {
    for ev := range orderSub {
        for _, o := range ev.Data {
            fmt.Printf("Order: %s %s %s status=%s\n", o.Symbol, o.Side, o.OrderType, o.OrderStatus)
        }
    }
}()

// Unsubscribe when done
orderCh.Unsubscribe(orderSub)
```

### Available Channel Managers

| Manager | Constructor | Subscribe Method | Topic Examples |
|---------|-------------|-----------------|----------------|
| `OrderChannel` | `NewOrderChannel(ws)` | `SubscribeToOrders(topic)` | `"order"`, `"order.linear"`, `"order.spot"` |
| `ExecutionChannel` | `NewExecutionChannel(ws)` | `SubscribeToExecutions(topic)` | `"execution"`, `"execution.linear"` |
| `FastExecutionChannel` | `NewFastExecutionChannel(ws)` | `SubscribeToFastExecutions(topic)` | `"execution.fast"`, `"execution.fast.linear"` |
| `PositionChannel` | `NewPositionChannel(ws)` | `SubscribeToPositions(topic)` | `"position"`, `"position.linear"` |
| `WalletChannel` | `NewWalletChannel(ws)` | `SubscribeToWallet()` | `"wallet"` |

Each manager has: `Subscribe()`, `Unsubscribe(ch)`, `GetSubscriberCount()`, `IsSubscribed()`.

### Topic Naming

- **All-in-one** (all categories): `"position"`, `"order"`, `"execution"`, `"execution.fast"`
- **Categorised** (specific category): `"position.linear"`, `"order.spot"`, `"execution.inverse"`, etc.
- **Standalone**: `"wallet"`

Convenience methods on `BybitWsClient`:

```go
ws.SubscribeAll("linear")    // position.linear + order.linear + execution.linear + execution.fast.linear + wallet
ws.SubscribeAllInOne()       // position + order + execution + execution.fast + wallet
```

### Callback Pattern (Alternative to Channels)

You can use direct callbacks instead of channel managers:

```go
ws.SetOrderCallback(func(ev bybitlib.OrderEvent) { ... })
ws.SetExecutionCallback(func(ev bybitlib.ExecutionEvent) { ... })
ws.SetFastExecutionCallback(func(ev bybitlib.FastExecutionEvent) { ... })
ws.SetPositionCallback(func(ev bybitlib.PositionEvent) { ... })
ws.SetWalletCallback(func(ev bybitlib.WalletEvent) { ... })
```

---

## WebSocket Event Types

### `OrderEvent` / `OrderData`

```go
ev.ID           // message ID
ev.Topic        // "order" or "order.linear"
ev.CreationTime // timestamp ms
ev.Data         // []OrderData

// OrderData fields:
d.Category, d.Symbol, d.Side, d.OrderType, d.OrderStatus
d.OrderId, d.OrderLinkId
d.Price, d.Qty, d.AvgPrice
d.CumExecQty, d.CumExecValue, d.CumExecFee
d.LeavesQty, d.LeavesValue
d.TakeProfit, d.StopLoss, d.TpslMode
d.StopOrderType, d.TriggerPrice, d.TriggerBy
d.TimeInForce, d.ReduceOnly, d.CloseOnTrigger
d.CreateType, d.CancelType, d.RejectReason
d.ClosedPnl, d.FeeCurrency
d.PositionIdx, d.TriggerDirection
d.CreatedTime, d.UpdatedTime
```

### `ExecutionEvent` / `ExecutionData`

```go
ev.Data // []ExecutionData

d.Category, d.Symbol, d.Side
d.OrderId, d.OrderLinkId
d.ExecId, d.ExecPrice, d.ExecQty, d.ExecValue
d.ExecPnl, d.ExecFee, d.ExecType, d.ExecTime
d.OrderPrice, d.OrderQty, d.OrderType
d.LeavesQty, d.CreateType, d.StopOrderType
d.IsMaker, d.FeeRate
d.MarkPrice, d.IndexPrice, d.UnderlyingPrice
d.ClosedSize, d.Seq
```

### `FastExecutionEvent` / `FastExecutionData`

Low-latency execution stream with fewer fields:

```go
ev.Data // []FastExecutionData

d.Category, d.Symbol, d.Side
d.OrderId, d.OrderLinkId
d.ExecId, d.ExecPrice, d.ExecQty, d.ExecValue
d.ExecFee, d.ExecType, d.ExecTime
d.OrderPrice, d.OrderQty, d.OrderType
d.LeavesQty, d.StopOrderType
d.IsMaker, d.Seq
```

### `PositionEvent` / `PositionData`

```go
ev.Data // []PositionData

d.Category, d.Symbol, d.Side, d.Size
d.EntryPrice, d.MarkPrice, d.Leverage
d.PositionValue, d.PositionBalance
d.PositionIM, d.PositionMM
d.UnrealisedPnl, d.CurRealisedPnl, d.CumRealisedPnl
d.LiqPrice, d.BustPrice
d.TakeProfit, d.StopLoss, d.TrailingStop, d.TpslMode
d.PositionIdx, d.TradeMode, d.AutoAddMargin
d.RiskId, d.RiskLimitValue
d.PositionStatus, d.AdlRankIndicator, d.IsReduceOnly
d.Delta, d.Gamma, d.Vega, d.Theta  // options only
d.CreatedTime, d.UpdatedTime, d.Seq
```

### `WalletEvent` / `WalletData` / `WalletCoinWS`

```go
ev.Data // []WalletData

w.AccountType            // "UNIFIED"
w.TotalEquity, w.TotalWalletBalance
w.TotalMarginBalance, w.TotalAvailableBalance
w.TotalPerpUPL
w.TotalInitialMargin, w.TotalMaintenanceMargin
w.AccountIMRate, w.AccountMMRate, w.AccountLTV
w.Coin // []WalletCoinWS

c.Coin, c.Equity, c.UsdValue
c.WalletBalance, c.Free, c.Locked
c.UnrealisedPnl, c.CumRealisedPnl
c.BorrowAmount, c.AvailableToWithdraw
c.TotalOrderIM, c.TotalPositionIM, c.TotalPositionMM
c.CollateralSwitch, c.MarginCollateral
```

---

## Constants

All constants are available from the top-level `bybitlib` package.

### Categories

| Constant | Value |
|----------|-------|
| `CategoryLinear` | `"linear"` |
| `CategoryInverse` | `"inverse"` |
| `CategorySpot` | `"spot"` |
| `CategoryOption` | `"option"` |

### Trade Sides

| Constant | Value |
|----------|-------|
| `SideBuy` | `"Buy"` |
| `SideSell` | `"Sell"` |

### Order Types

| Constant | Value |
|----------|-------|
| `OrderTypeLimit` | `"Limit"` |
| `OrderTypeMarket` | `"Market"` |

### Time in Force

| Constant | Value |
|----------|-------|
| `TimeInForceGTC` | `"GTC"` |
| `TimeInForceIOC` | `"IOC"` |
| `TimeInForceFOK` | `"FOK"` |
| `TimeInForcePostOnly` | `"PostOnly"` |

### Margin Modes

| Constant | Value |
|----------|-------|
| `MarginModeIsolated` | `"ISOLATED_MARGIN"` |
| `MarginModeRegular` | `"REGULAR_MARGIN"` |
| `MarginModePortfolio` | `"PORTFOLIO_MARGIN"` |
| `MarginModeCross` | `"CROSS_MARGIN"` |

### Position Index (Hedge Mode)

| Constant | Value |
|----------|-------|
| `PositionIdxOneWay` | `0` |
| `PositionIdxBuySide` | `1` |
| `PositionIdxSellSide` | `2` |

### Trigger Types

| Constant | Value |
|----------|-------|
| `TriggerByLastPrice` | `"LastPrice"` |
| `TriggerByIndexPrice` | `"IndexPrice"` |
| `TriggerByMarkPrice` | `"MarkPrice"` |

### TP/SL Modes

| Constant | Value |
|----------|-------|
| `TpslModeFull` | `"Full"` |
| `TpslModePartial` | `"Partial"` |

### Account Types

| Constant | Value |
|----------|-------|
| `AccountTypeUnified` | `"UNIFIED"` |
| `AccountTypeContract` | `"CONTRACT"` |

### Order Status

| Constant | Value |
|----------|-------|
| `OrderStatusNew` | `"New"` |
| `OrderStatusPartiallyFilled` | `"PartiallyFilled"` |
| `OrderStatusFilled` | `"Filled"` |
| `OrderStatusCancelled` | `"Cancelled"` |
| `OrderStatusRejected` | `"Rejected"` |
| `OrderStatusDeactivated` | `"Deactivated"` |

### STP Modes

| Constant | Value |
|----------|-------|
| `STPModeNone` | `"none"` |
| `STPModeCancelMaker` | `"CancelMaker"` |
| `STPModeCancelTaker` | `"CancelTaker"` |
| `STPModeCancelBoth` | `"CancelBoth"` |

---

## Project Structure

```
Khavr-bybit-lib/
├── bybitlib.go                  # Top-level re-exports (REST + WS)
├── go.mod
├── bybit/                       # REST client package
│   ├── client.go                # BybitClient, NewBybitClient()
│   ├── do_request.go            # doGet(), doPost() with HMAC auth
│   ├── utils.go                 # BuildGetParams(), GetCurrentTime()
│   ├── create_order.go          # CreateOrder()
│   ├── account_amend_order.go   # AmendOrder()
│   ├── account_set_margin_mode.go
│   ├── client_get_positions.go  # GetPositions()
│   ├── GetClosedPnL.go          # GetClosedPnL()
│   ├── GetClosePos.go           # GetClosedOptionsPositions()
│   ├── GetWalletBalance.go      # GetWalletBalance()
│   ├── get_realtime_orders.go   # GetRealtimeOrders()
│   ├── trading_stopLoss.go      # SetTradingStop()
│   └── bybit_models/            # Request/response types
├── bybit_ws/                    # Private WebSocket client package
│   ├── ws_client.go             # BybitWsClient (connect, auth, listen, reconnect)
│   ├── ws_types.go              # Protocol types (auth, subscribe)
│   ├── order_channel.go         # OrderChannel pub/sub manager
│   ├── execution_channel.go     # ExecutionChannel pub/sub manager
│   ├── fast_execution_channel.go
│   ├── position_channel.go      # PositionChannel pub/sub manager
│   ├── wallet_channel.go        # WalletChannel pub/sub manager
│   └── bybit_ws_models/         # WebSocket event types
│       ├── order_event.go
│       ├── execution_event.go
│       ├── position_event.go
│       └── wallet_event.go
└── cmd/ws_test/                 # Test binary (not part of library)
```
