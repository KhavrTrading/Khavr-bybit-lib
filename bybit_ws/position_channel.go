package bybit_ws

import (
	"sync"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"

	log "github.com/sirupsen/logrus"
)

// PositionChannel manages position updates from Bybit private WebSocket
// and distributes them to multiple subscribers via channels.
type PositionChannel struct {
	client *BybitWsClient

	subscribers []chan bybit_ws_models.PositionEvent
	subMu       sync.RWMutex

	topic      string
	subscribed bool
	subStateMu sync.Mutex
}

// NewPositionChannel creates a new position channel manager.
func NewPositionChannel(client *BybitWsClient) *PositionChannel {
	pc := &PositionChannel{
		client:      client,
		subscribers: make([]chan bybit_ws_models.PositionEvent, 0),
	}
	client.SetPositionCallback(pc.handleUpdate)
	return pc
}

// Subscribe returns a buffered channel that will receive position events.
func (pc *PositionChannel) Subscribe() chan bybit_ws_models.PositionEvent {
	ch := make(chan bybit_ws_models.PositionEvent, 10)

	pc.subMu.Lock()
	pc.subscribers = append(pc.subscribers, ch)
	pc.subMu.Unlock()

	log.Infof("[PositionChannel] New subscriber added, total: %d", len(pc.subscribers))
	return ch
}

// Unsubscribe removes a subscriber channel.
func (pc *PositionChannel) Unsubscribe(ch chan bybit_ws_models.PositionEvent) {
	pc.subMu.Lock()
	defer pc.subMu.Unlock()

	for i, sub := range pc.subscribers {
		if sub == ch {
			close(ch)
			pc.subscribers = append(pc.subscribers[:i], pc.subscribers[i+1:]...)
			log.Infof("[PositionChannel] Subscriber removed, total: %d", len(pc.subscribers))
			return
		}
	}
}

// SubscribeToPositions subscribes to position updates on the WebSocket.
// topic: "position" (all-in-one) or "position.linear", "position.inverse", etc.
func (pc *PositionChannel) SubscribeToPositions(topic string) error {
	pc.subStateMu.Lock()
	defer pc.subStateMu.Unlock()

	if pc.subscribed {
		log.Warnf("[PositionChannel] Already subscribed to position channel")
		return nil
	}

	pc.topic = topic
	if err := pc.client.SubscribePosition(topic); err != nil {
		return err
	}

	pc.subscribed = true
	log.Infof("[PositionChannel] Subscribed to %s", topic)
	return nil
}

func (pc *PositionChannel) handleUpdate(update bybit_ws_models.PositionEvent) {
	pc.subMu.RLock()
	defer pc.subMu.RUnlock()

	for _, sub := range pc.subscribers {
		select {
		case sub <- update:
		default:
			select {
			case <-sub:
			default:
			}
			sub <- update
			log.Warnf("[PositionChannel] Subscriber channel was full, drained oldest message")
		}
	}
}

// GetSubscriberCount returns the number of active subscribers.
func (pc *PositionChannel) GetSubscriberCount() int {
	pc.subMu.RLock()
	defer pc.subMu.RUnlock()
	return len(pc.subscribers)
}

// IsSubscribed returns whether we're subscribed to the position channel.
func (pc *PositionChannel) IsSubscribed() bool {
	pc.subStateMu.Lock()
	defer pc.subStateMu.Unlock()
	return pc.subscribed
}
