// Package bybitlib provides a Go client for the Bybit REST API.
//
// Usage:
//
//	client := bybitlib.NewBybitClient("apiKey", "apiSecret")
//	resp, err := client.CreateOrder(&bybitlib.CreateOrderRequest{...})
package bybitlib

import (
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
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

// ── Interfaces ──

type ParamBuilder = bybit_models.ParamBuilder

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
