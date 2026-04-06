package bybit_ws

import (
	"sync"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"

	log "github.com/sirupsen/logrus"
)

// OrderChannel manages order updates from Bybit private WebSocket
// and distributes them to multiple subscribers via channels.
type OrderChannel struct {
	client *BybitWsClient

	subscribers []chan bybit_ws_models.OrderEvent
	subMu       sync.RWMutex

	topic      string
	subscribed bool
	subStateMu sync.Mutex
}

// NewOrderChannel creates a new order channel manager.
func NewOrderChannel(client *BybitWsClient) *OrderChannel {
	oc := &OrderChannel{
		client:      client,
		subscribers: make([]chan bybit_ws_models.OrderEvent, 0),
	}
	client.SetOrderCallback(oc.handleUpdate)
	return oc
}

// Subscribe returns a buffered channel that will receive order events.
func (oc *OrderChannel) Subscribe() chan bybit_ws_models.OrderEvent {
	ch := make(chan bybit_ws_models.OrderEvent, 10)

	oc.subMu.Lock()
	oc.subscribers = append(oc.subscribers, ch)
	oc.subMu.Unlock()

	log.Infof("[OrderChannel] New subscriber added, total: %d", len(oc.subscribers))
	return ch
}

// Unsubscribe removes a subscriber channel.
func (oc *OrderChannel) Unsubscribe(ch chan bybit_ws_models.OrderEvent) {
	oc.subMu.Lock()
	defer oc.subMu.Unlock()

	for i, sub := range oc.subscribers {
		if sub == ch {
			close(ch)
			oc.subscribers = append(oc.subscribers[:i], oc.subscribers[i+1:]...)
			log.Infof("[OrderChannel] Subscriber removed, total: %d", len(oc.subscribers))
			return
		}
	}
}

// SubscribeToOrders subscribes to order updates on the WebSocket.
// topic: "order" (all-in-one) or "order.linear", "order.spot", etc.
func (oc *OrderChannel) SubscribeToOrders(topic string) error {
	oc.subStateMu.Lock()
	defer oc.subStateMu.Unlock()

	if oc.subscribed {
		log.Warnf("[OrderChannel] Already subscribed to order channel")
		return nil
	}

	oc.topic = topic
	if err := oc.client.SubscribeOrder(topic); err != nil {
		return err
	}

	oc.subscribed = true
	log.Infof("[OrderChannel] Subscribed to %s", topic)
	return nil
}

func (oc *OrderChannel) handleUpdate(update bybit_ws_models.OrderEvent) {
	oc.subMu.RLock()
	defer oc.subMu.RUnlock()

	for _, sub := range oc.subscribers {
		select {
		case sub <- update:
		default:
			select {
			case <-sub:
			default:
			}
			sub <- update
			log.Warnf("[OrderChannel] Subscriber channel was full, drained oldest message")
		}
	}
}

// GetSubscriberCount returns the number of active subscribers.
func (oc *OrderChannel) GetSubscriberCount() int {
	oc.subMu.RLock()
	defer oc.subMu.RUnlock()
	return len(oc.subscribers)
}

// IsSubscribed returns whether we're subscribed to the order channel.
func (oc *OrderChannel) IsSubscribed() bool {
	oc.subStateMu.Lock()
	defer oc.subStateMu.Unlock()
	return oc.subscribed
}
