package bybit_ws

import (
	"sync"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"

	log "github.com/sirupsen/logrus"
)

// WalletChannel manages wallet balance updates from Bybit private WebSocket
// and distributes them to multiple subscribers via channels.
type WalletChannel struct {
	client *BybitWsClient

	subscribers []chan bybit_ws_models.WalletEvent
	subMu       sync.RWMutex

	subscribed bool
	subStateMu sync.Mutex
}

// NewWalletChannel creates a new wallet channel manager.
func NewWalletChannel(client *BybitWsClient) *WalletChannel {
	wc := &WalletChannel{
		client:      client,
		subscribers: make([]chan bybit_ws_models.WalletEvent, 0),
	}
	client.SetWalletCallback(wc.handleUpdate)
	return wc
}

// Subscribe returns a buffered channel that will receive wallet events.
func (wc *WalletChannel) Subscribe() chan bybit_ws_models.WalletEvent {
	ch := make(chan bybit_ws_models.WalletEvent, 10)

	wc.subMu.Lock()
	wc.subscribers = append(wc.subscribers, ch)
	wc.subMu.Unlock()

	log.Infof("[WalletChannel] New subscriber added, total: %d", len(wc.subscribers))
	return ch
}

// Unsubscribe removes a subscriber channel.
func (wc *WalletChannel) Unsubscribe(ch chan bybit_ws_models.WalletEvent) {
	wc.subMu.Lock()
	defer wc.subMu.Unlock()

	for i, sub := range wc.subscribers {
		if sub == ch {
			close(ch)
			wc.subscribers = append(wc.subscribers[:i], wc.subscribers[i+1:]...)
			log.Infof("[WalletChannel] Subscriber removed, total: %d", len(wc.subscribers))
			return
		}
	}
}

// SubscribeToWallet subscribes to wallet updates on the WebSocket.
func (wc *WalletChannel) SubscribeToWallet() error {
	wc.subStateMu.Lock()
	defer wc.subStateMu.Unlock()

	if wc.subscribed {
		log.Warnf("[WalletChannel] Already subscribed to wallet channel")
		return nil
	}

	if err := wc.client.SubscribeWallet(); err != nil {
		return err
	}

	wc.subscribed = true
	log.Infof("[WalletChannel] Subscribed to wallet")
	return nil
}

func (wc *WalletChannel) handleUpdate(update bybit_ws_models.WalletEvent) {
	wc.subMu.RLock()
	defer wc.subMu.RUnlock()

	for _, sub := range wc.subscribers {
		select {
		case sub <- update:
		default:
			select {
			case <-sub:
			default:
			}
			sub <- update
			log.Warnf("[WalletChannel] Subscriber channel was full, drained oldest message")
		}
	}
}

// GetSubscriberCount returns the number of active subscribers.
func (wc *WalletChannel) GetSubscriberCount() int {
	wc.subMu.RLock()
	defer wc.subMu.RUnlock()
	return len(wc.subscribers)
}

// IsSubscribed returns whether we're subscribed to the wallet channel.
func (wc *WalletChannel) IsSubscribed() bool {
	wc.subStateMu.Lock()
	defer wc.subStateMu.Unlock()
	return wc.subscribed
}
