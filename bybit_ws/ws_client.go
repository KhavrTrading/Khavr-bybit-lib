package bybit_ws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

// BybitWsClient handles Bybit private WebSocket connections.
type BybitWsClient struct {
	// Connection
	WsConn    *websocket.Conn
	BaseWsURL string
	stopChan  chan struct{}
	connMu    sync.Mutex

	// Authentication
	apiKey    string
	apiSecret string
	loggedIn  bool

	// Callbacks
	orderCallback         func(bybit_ws_models.OrderEvent)
	executionCallback     func(bybit_ws_models.ExecutionEvent)
	fastExecutionCallback func(bybit_ws_models.FastExecutionEvent)
	positionCallback      func(bybit_ws_models.PositionEvent)
	walletCallback        func(bybit_ws_models.WalletEvent)

	// Subscriptions tracking (for reconnect)
	activeTopics []string
	subMu        sync.RWMutex
}

// NewBybitWsClient creates a new private WebSocket client.
func NewBybitWsClient(apiKey, apiSecret string) *BybitWsClient {
	return &BybitWsClient{
		BaseWsURL:    "wss://stream.bybit.com/v5/private",
		stopChan:     make(chan struct{}),
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		activeTopics: make([]string, 0),
	}
}

// Connect establishes the WebSocket connection and authenticates.
func (c *BybitWsClient) Connect() error {
	dialer := *websocket.DefaultDialer
	dialer.EnableCompression = true
	dialer.ReadBufferSize = 16 * 1024
	dialer.WriteBufferSize = 16 * 1024

	var err error
	c.WsConn, _, err = dialer.Dial(c.BaseWsURL, nil)
	if err != nil {
		return fmt.Errorf("failed to connect to Bybit private WebSocket: %w", err)
	}

	log.Info("[Bybit:Private] Connected to WebSocket")

	if err := c.sendAuth(); err != nil {
		c.Close()
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	return nil
}

// Close closes the WebSocket connection.
func (c *BybitWsClient) Close() error {
	if c.WsConn == nil {
		return nil
	}
	c.loggedIn = false
	log.Info("[Bybit:Private] Closing WebSocket connection")
	return c.WsConn.Close()
}

// Stop signals the client to stop.
func (c *BybitWsClient) Stop() {
	select {
	case <-c.stopChan:
	default:
		close(c.stopChan)
	}
}

// ── Callback setters ──

func (c *BybitWsClient) SetOrderCallback(cb func(bybit_ws_models.OrderEvent)) {
	c.orderCallback = cb
}

func (c *BybitWsClient) SetExecutionCallback(cb func(bybit_ws_models.ExecutionEvent)) {
	c.executionCallback = cb
}

func (c *BybitWsClient) SetFastExecutionCallback(cb func(bybit_ws_models.FastExecutionEvent)) {
	c.fastExecutionCallback = cb
}

func (c *BybitWsClient) SetPositionCallback(cb func(bybit_ws_models.PositionEvent)) {
	c.positionCallback = cb
}

func (c *BybitWsClient) SetWalletCallback(cb func(bybit_ws_models.WalletEvent)) {
	c.walletCallback = cb
}

// ── Subscribe methods ──

// SubscribeOrder subscribes to the order topic.
// Use "order" for all categories, or "order.linear", "order.spot", etc.
func (c *BybitWsClient) SubscribeOrder(topic string) error {
	return c.addAndSubscribe(topic)
}

// SubscribeExecution subscribes to the execution topic.
// Use "execution" for all categories, or "execution.linear", etc.
func (c *BybitWsClient) SubscribeExecution(topic string) error {
	return c.addAndSubscribe(topic)
}

// SubscribeFastExecution subscribes to the fast execution topic.
// Use "execution.fast" for all, or "execution.fast.linear", etc.
func (c *BybitWsClient) SubscribeFastExecution(topic string) error {
	return c.addAndSubscribe(topic)
}

// SubscribePosition subscribes to the position topic.
// Use "position" for all categories, or "position.linear", etc.
func (c *BybitWsClient) SubscribePosition(topic string) error {
	return c.addAndSubscribe(topic)
}

// SubscribeWallet subscribes to the wallet topic.
func (c *BybitWsClient) SubscribeWallet() error {
	return c.addAndSubscribe("wallet")
}

// SubscribeAll subscribes to all private topics for a given category.
// e.g. SubscribeAll("linear") subscribes to position.linear, order.linear, execution.linear, execution.fast.linear, wallet.
func (c *BybitWsClient) SubscribeAll(category string) error {
	topics := []string{
		"position." + category,
		"order." + category,
		"execution." + category,
		"execution.fast." + category,
		"wallet",
	}
	return c.subscribeTopics(topics)
}

// SubscribeAllInOne subscribes to all-in-one topics (all categories).
func (c *BybitWsClient) SubscribeAllInOne() error {
	topics := []string{"position", "order", "execution", "execution.fast", "wallet"}
	return c.subscribeTopics(topics)
}

// ── Auth ──

// sendAuth sends the authentication message to Bybit.
// Signature: hex(HMAC-SHA256(apiSecret, "GET/realtime" + expires))
func (c *BybitWsClient) sendAuth() error {
	expires := time.Now().UnixMilli() + 5000

	payload := fmt.Sprintf("GET/realtime%d", expires)
	h := hmac.New(sha256.New, []byte(c.apiSecret))
	h.Write([]byte(payload))
	sign := hex.EncodeToString(h.Sum(nil))

	authReq := wsAuthRequest{
		Op:   "auth",
		Args: []any{c.apiKey, expires, sign},
	}

	reqBytes, err := json.Marshal(authReq)
	if err != nil {
		return fmt.Errorf("marshal auth request: %w", err)
	}

	c.connMu.Lock()
	err = c.WsConn.WriteMessage(websocket.TextMessage, reqBytes)
	c.connMu.Unlock()
	if err != nil {
		return fmt.Errorf("send auth request: %w", err)
	}

	log.Info("[Bybit:Private] Sent auth request")
	return c.waitForAuthResponse(10 * time.Second)
}

func (c *BybitWsClient) waitForAuthResponse(timeout time.Duration) error {
	c.WsConn.SetReadDeadline(time.Now().Add(timeout))
	defer c.WsConn.SetReadDeadline(time.Time{})

	_, msg, err := c.WsConn.ReadMessage()
	if err != nil {
		return fmt.Errorf("read auth response: %w", err)
	}

	var resp wsOpResponse
	if err := json.Unmarshal(msg, &resp); err != nil {
		return fmt.Errorf("unmarshal auth response: %w", err)
	}

	if resp.Op == "auth" && resp.Success {
		c.loggedIn = true
		log.Info("[Bybit:Private] Auth successful")
		return nil
	}

	return fmt.Errorf("auth failed: op=%s success=%v msg=%s", resp.Op, resp.Success, resp.RetMsg)
}

// ── Subscribe internals ──

func (c *BybitWsClient) addAndSubscribe(topic string) error {
	if !c.loggedIn {
		return fmt.Errorf("not authenticated")
	}

	c.subMu.Lock()
	c.activeTopics = append(c.activeTopics, topic)
	c.subMu.Unlock()

	return c.subscribe([]string{topic})
}

func (c *BybitWsClient) subscribeTopics(topics []string) error {
	if !c.loggedIn {
		return fmt.Errorf("not authenticated")
	}

	c.subMu.Lock()
	c.activeTopics = append(c.activeTopics, topics...)
	c.subMu.Unlock()

	return c.subscribe(topics)
}

func (c *BybitWsClient) subscribe(topics []string) error {
	subReq := wsSubscribeRequest{
		Op:   "subscribe",
		Args: topics,
	}

	reqBytes, err := json.Marshal(subReq)
	if err != nil {
		return fmt.Errorf("marshal subscribe request: %w", err)
	}

	c.connMu.Lock()
	err = c.WsConn.WriteMessage(websocket.TextMessage, reqBytes)
	c.connMu.Unlock()
	if err != nil {
		return fmt.Errorf("send subscribe request: %w", err)
	}

	for _, t := range topics {
		log.Infof("[Bybit:Private] Subscribed to %s", t)
	}
	return nil
}

func (c *BybitWsClient) resubscribeAll() error {
	c.subMu.RLock()
	topics := make([]string, len(c.activeTopics))
	copy(topics, c.activeTopics)
	c.subMu.RUnlock()

	if len(topics) == 0 {
		return nil
	}

	log.Infof("[Bybit:Private] Resubscribing to %d topics", len(topics))
	return c.subscribe(topics)
}

// ── Ping ──

func (c *BybitWsClient) sendPing() error {
	c.connMu.Lock()
	defer c.connMu.Unlock()

	pingBytes, _ := json.Marshal(map[string]string{"op": "ping"})
	if err := c.WsConn.WriteMessage(websocket.TextMessage, pingBytes); err != nil {
		return fmt.Errorf("failed to send ping: %w", err)
	}

	log.Debug("[Bybit:Private] Sent ping")
	return nil
}

// ── Listen loop ──

// ListenLoop reads messages and dispatches them to callbacks.
// This blocks until Stop() is called or the connection is permanently lost.
func (c *BybitWsClient) ListenLoop() {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("[Bybit:Private] ws.panic: %v", r)
			c.Stop()
			c.Close()
		}
	}()

	pingTicker := time.NewTicker(20 * time.Second)
	defer pingTicker.Stop()

	go func() {
		for {
			select {
			case <-c.stopChan:
				return
			case <-pingTicker.C:
				if err := c.sendPing(); err != nil {
					log.Warnf("[Bybit:Private] ws.ping.error: %v", err)
				}
			}
		}
	}()

	for {
		select {
		case <-c.stopChan:
			log.Info("[Bybit:Private] Stopping listen loop")
			c.Close()
			return

		default:
			c.WsConn.SetReadDeadline(time.Now().Add(60 * time.Second))
			_, msg, err := c.WsConn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Warnf("[Bybit:Private] ws.read.error: %s", err)
				}

				log.Info("[Bybit:Private] Attempting reconnect...")
				time.Sleep(2 * time.Second)
				if err := c.reconnect(); err != nil {
					log.Errorf("[Bybit:Private] reconnect failed: %v", err)
				}
				continue
			}

			c.dispatchMessage(msg)
		}
	}
}

// ── Dispatch ──

func (c *BybitWsClient) dispatchMessage(msg []byte) {
	// Check for op responses (auth, subscribe, pong)
	var opResp struct {
		Op      string `json:"op"`
		Success bool   `json:"success"`
		RetMsg  string `json:"ret_msg"`
	}
	if json.Unmarshal(msg, &opResp) == nil && opResp.Op != "" {
		switch opResp.Op {
		case "subscribe":
			log.Debugf("[Bybit:Private] Subscription confirmed")
		case "pong":
			log.Debug("[Bybit:Private] Pong received")
		case "auth":
			log.Debug("[Bybit:Private] Auth event received")
		}
		if !opResp.Success && opResp.RetMsg != "" {
			log.Errorf("[Bybit:Private] Error: op=%s msg=%s", opResp.Op, opResp.RetMsg)
		}
		return
	}

	// Parse topic field to determine event type
	var base struct {
		Topic string `json:"topic"`
	}
	if err := json.Unmarshal(msg, &base); err != nil || base.Topic == "" {
		return
	}

	topic := base.Topic

	switch {
	case topic == "wallet":
		if c.walletCallback != nil {
			var ev bybit_ws_models.WalletEvent
			if err := json.Unmarshal(msg, &ev); err == nil {
				c.walletCallback(ev)
			} else {
				log.Warnf("[Bybit:Private] Failed to parse wallet event: %v", err)
			}
		}

	case topic == "execution.fast" || strings.HasPrefix(topic, "execution.fast."):
		if c.fastExecutionCallback != nil {
			var ev bybit_ws_models.FastExecutionEvent
			if err := json.Unmarshal(msg, &ev); err == nil {
				c.fastExecutionCallback(ev)
			} else {
				log.Warnf("[Bybit:Private] Failed to parse fast execution event: %v", err)
			}
		}

	case topic == "execution" || strings.HasPrefix(topic, "execution."):
		if c.executionCallback != nil {
			var ev bybit_ws_models.ExecutionEvent
			if err := json.Unmarshal(msg, &ev); err == nil {
				c.executionCallback(ev)
			} else {
				log.Warnf("[Bybit:Private] Failed to parse execution event: %v", err)
			}
		}

	case topic == "order" || strings.HasPrefix(topic, "order."):
		if c.orderCallback != nil {
			var ev bybit_ws_models.OrderEvent
			if err := json.Unmarshal(msg, &ev); err == nil {
				c.orderCallback(ev)
			} else {
				log.Warnf("[Bybit:Private] Failed to parse order event: %v", err)
			}
		}

	case topic == "position" || strings.HasPrefix(topic, "position."):
		if c.positionCallback != nil {
			var ev bybit_ws_models.PositionEvent
			if err := json.Unmarshal(msg, &ev); err == nil {
				c.positionCallback(ev)
			} else {
				log.Warnf("[Bybit:Private] Failed to parse position event: %v", err)
			}
		}

	default:
		log.Debugf("[Bybit:Private] Unknown topic: %s", topic)
	}
}

// ── Reconnect ──

func (c *BybitWsClient) reconnect() error {
	if err := c.Close(); err != nil {
		log.Debugf("[Bybit:Private] Error closing WebSocket: %v", err)
	}

	time.Sleep(1 * time.Second)

	if err := c.Connect(); err != nil {
		return fmt.Errorf("failed to reconnect: %w", err)
	}

	if err := c.resubscribeAll(); err != nil {
		return fmt.Errorf("failed to resubscribe: %w", err)
	}

	log.Info("[Bybit:Private] Reconnected and resubscribed successfully")
	return nil
}
