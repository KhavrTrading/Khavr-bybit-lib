// Package bybitlib provides a Go client for the Bybit REST and WebSocket APIs.
//
// Usage:
//
//	client := bybitlib.NewBybitClient("apiKey", "apiSecret")
//	resp, err := client.CreateOrder(&bybitlib.CreateOrderRequest{...})
//
//	wsClient := bybitlib.NewBybitWsClient("apiKey", "apiSecret")
//	wsClient.SetPositionCallback(func(ev bybitlib.PositionEvent) { ... })
package bybitlib

import (
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"
)

// BybitClient is the REST API client for synchronous operations.
type BybitClient = bybit.BybitClient

// NewBybitClient creates a new Bybit REST client with the given credentials.
func NewBybitClient(apiKey, apiSecret string) *BybitClient {
	return bybit.NewBybitClient(apiKey, apiSecret)
}

// ── Order models ──

type CreateOrderRequest = bybit_models.CreateOrderRequest
type CreateOrderResponse = bybit_models.CreateOrderResponse
type CreateOrderResult = bybit_models.CreateOrderResult

type AmendOrderRequest = bybit_models.AmendOrderRequest
type AmendOrderResponse = bybit_models.AmendOrderResponse

type CancelOrderRequest = bybit_models.CancelOrderRequest
type CancelOrderResponse = bybit_models.CancelOrderResponse

type OrderRealtimeRequest = bybit_models.OrderRealtimeRequest
type OrderRealtimeResponse = bybit_models.OrderRealtimeResponse
type OrderRealtimeItem = bybit_models.OrderRealtimeItem

// ── Position models ──

type PositionListRequest = bybit_models.PositionListRequest
type PositionListResponse = bybit_models.PositionListResponse
type PositionListItem = bybit_models.PositionListItem

type TradingStopRequest = bybit_models.TradingStopRequest
type TradingStopResponse = bybit_models.TradingStopResponse

// ── Closed PnL models ──

type ClosedPnlRequest = bybit_models.ClosedPnlRequest
type ClosedPnlResponse = bybit_models.ClosedPnlResponse
type ClosedPnlResult = bybit_models.ClosedPnlResult
type ClosedPnlEntry = bybit_models.ClosedPnlEntry

// ── Closed options positions models ──

type ClosedOptionsPositionsRequest = bybit_models.ClosedOptionsPositionsRequest
type ClosedOptionsPositionsResponse = bybit_models.ClosedOptionsPositionsResponse
type ClosedOptionsPositionsResult = bybit_models.ClosedOptionsPositionsResult
type ClosedOptionsPositionEntry = bybit_models.ClosedOptionsPositionEntry

// ── Account models ──

type SetMarginModeRequest = bybit_models.SetMarginModeRequest
type SetMarginModeResponse = bybit_models.SetMarginModeResponse

type WalletBalanceRequest = bybit_models.WalletBalanceRequest
type WalletBalanceResponse = bybit_models.WalletBalanceResponse
type WalletBalanceData = bybit_models.WalletBalanceData
type WalletAccount = bybit_models.WalletAccount
type WalletCoin = bybit_models.WalletCoin

// ── Spread models ──

type SpreadInstrumentRequest = bybit_models.SpreadInstrumentRequest
type SpreadInstrumentResponse = bybit_models.SpreadInstrumentResponse
type SpreadInstrumentResult = bybit_models.SpreadInstrumentResult
type SpreadInstrumentItem = bybit_models.SpreadInstrumentItem
type SpreadLeg = bybit_models.SpreadLeg

// ── Interfaces ──

type ParamBuilder = bybit_models.ParamBuilder

// ── WebSocket client ──

// BybitWsClient is the private WebSocket client for real-time data streams.
type BybitWsClient = bybit_ws.BybitWsClient

// NewBybitWsClient creates a new Bybit private WebSocket client.
func NewBybitWsClient(apiKey, apiSecret string) *BybitWsClient {
	return bybit_ws.NewBybitWsClient(apiKey, apiSecret)
}

// ── WebSocket event models ──

type OrderEvent = bybit_ws_models.OrderEvent
type OrderData = bybit_ws_models.OrderData

type ExecutionEvent = bybit_ws_models.ExecutionEvent
type ExecutionData = bybit_ws_models.ExecutionData

type FastExecutionEvent = bybit_ws_models.FastExecutionEvent
type FastExecutionData = bybit_ws_models.FastExecutionData

type PositionEvent = bybit_ws_models.PositionEvent
type PositionData = bybit_ws_models.PositionData

type WalletEvent = bybit_ws_models.WalletEvent
type WalletData_WS = bybit_ws_models.WalletData
type WalletCoinWS = bybit_ws_models.WalletCoinWS

// ── WebSocket channel managers (pub/sub) ──

type OrderChannel = bybit_ws.OrderChannel
type ExecutionChannel = bybit_ws.ExecutionChannel
type FastExecutionChannel = bybit_ws.FastExecutionChannel
type PositionChannel = bybit_ws.PositionChannel
type WalletChannel = bybit_ws.WalletChannel

var NewOrderChannel = bybit_ws.NewOrderChannel
var NewExecutionChannel = bybit_ws.NewExecutionChannel
var NewFastExecutionChannel = bybit_ws.NewFastExecutionChannel
var NewPositionChannel = bybit_ws.NewPositionChannel
var NewWalletChannel = bybit_ws.NewWalletChannel

// ── Constants ──

const (
	// Categories
	CategoryLinear  = bybit_models.CategoryLinear
	CategoryInverse = bybit_models.CategoryInverse
	CategorySpot    = bybit_models.CategorySpot
	CategoryOption  = bybit_models.CategoryOption

	// Margin modes
	MarginModeIsolated  = bybit_models.MarginModeIsolated
	MarginModeRegular   = bybit_models.MarginModeRegular
	MarginModePortfolio = bybit_models.MarginModePortfolio
	MarginModeCross     = bybit_models.MarginModeCross

	// Trade sides
	SideBuy  = bybit_models.SideBuy
	SideSell = bybit_models.SideSell

	// Order types
	OrderTypeLimit  = bybit_models.OrderTypeLimit
	OrderTypeMarket = bybit_models.OrderTypeMarket

	// Time in force
	TimeInForceGTC      = bybit_models.TimeInForceGTC
	TimeInForceIOC      = bybit_models.TimeInForceIOC
	TimeInForceFOK      = bybit_models.TimeInForceFOK
	TimeInForcePostOnly = bybit_models.TimeInForcePostOnly

	// Position index
	PositionIdxOneWay   = bybit_models.PositionIdxOneWay
	PositionIdxBuySide  = bybit_models.PositionIdxBuySide
	PositionIdxSellSide = bybit_models.PositionIdxSellSide

	// Trigger price types
	TriggerByLastPrice  = bybit_models.TriggerByLastPrice
	TriggerByIndexPrice = bybit_models.TriggerByIndexPrice
	TriggerByMarkPrice  = bybit_models.TriggerByMarkPrice

	// TP/SL modes
	TpslModeFull    = bybit_models.TpslModeFull
	TpslModePartial = bybit_models.TpslModePartial

	// Account types
	AccountTypeUnified  = bybit_models.AccountTypeUnified
	AccountTypeContract = bybit_models.AccountTypeContract

	// Order status
	OrderStatusNew             = bybit_models.OrderStatusNew
	OrderStatusPartiallyFilled = bybit_models.OrderStatusPartiallyFilled
	OrderStatusFilled          = bybit_models.OrderStatusFilled
	OrderStatusCancelled       = bybit_models.OrderStatusCancelled
	OrderStatusRejected        = bybit_models.OrderStatusRejected
	OrderStatusDeactivated     = bybit_models.OrderStatusDeactivated

	// STP modes
	STPModeNone        = bybit_models.STPModeNone
	STPModeCancelMaker = bybit_models.STPModeCancelMaker
	STPModeCancelTaker = bybit_models.STPModeCancelTaker
	STPModeCancelBoth  = bybit_models.STPModeCancelBoth
)
