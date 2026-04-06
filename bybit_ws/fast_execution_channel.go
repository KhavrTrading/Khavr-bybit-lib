package bybit_ws

import (
	"sync"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"

	log "github.com/sirupsen/logrus"
)

// FastExecutionChannel manages fast execution updates from Bybit private WebSocket
// and distributes them to multiple subscribers via channels.
type FastExecutionChannel struct {
	client *BybitWsClient

	subscribers []chan bybit_ws_models.FastExecutionEvent
	subMu       sync.RWMutex

	topic      string
	subscribed bool
	subStateMu sync.Mutex
}

// NewFastExecutionChannel creates a new fast execution channel manager.
func NewFastExecutionChannel(client *BybitWsClient) *FastExecutionChannel {
	fc := &FastExecutionChannel{
		client:      client,
		subscribers: make([]chan bybit_ws_models.FastExecutionEvent, 0),
	}
	client.SetFastExecutionCallback(fc.handleUpdate)
	return fc
}

// Subscribe returns a buffered channel that will receive fast execution events.
func (fc *FastExecutionChannel) Subscribe() chan bybit_ws_models.FastExecutionEvent {
	ch := make(chan bybit_ws_models.FastExecutionEvent, 10)

	fc.subMu.Lock()
	fc.subscribers = append(fc.subscribers, ch)
	fc.subMu.Unlock()

	log.Infof("[FastExecutionChannel] New subscriber added, total: %d", len(fc.subscribers))
	return ch
}

// Unsubscribe removes a subscriber channel.
func (fc *FastExecutionChannel) Unsubscribe(ch chan bybit_ws_models.FastExecutionEvent) {
	fc.subMu.Lock()
	defer fc.subMu.Unlock()

	for i, sub := range fc.subscribers {
		if sub == ch {
			close(ch)
			fc.subscribers = append(fc.subscribers[:i], fc.subscribers[i+1:]...)
			log.Infof("[FastExecutionChannel] Subscriber removed, total: %d", len(fc.subscribers))
			return
		}
	}
}

// SubscribeToFastExecutions subscribes to fast execution updates on the WebSocket.
// topic: "execution.fast" (all-in-one) or "execution.fast.linear", etc.
func (fc *FastExecutionChannel) SubscribeToFastExecutions(topic string) error {
	fc.subStateMu.Lock()
	defer fc.subStateMu.Unlock()

	if fc.subscribed {
		log.Warnf("[FastExecutionChannel] Already subscribed to fast execution channel")
		return nil
	}

	fc.topic = topic
	if err := fc.client.SubscribeFastExecution(topic); err != nil {
		return err
	}

	fc.subscribed = true
	log.Infof("[FastExecutionChannel] Subscribed to %s", topic)
	return nil
}

func (fc *FastExecutionChannel) handleUpdate(update bybit_ws_models.FastExecutionEvent) {
	fc.subMu.RLock()
	defer fc.subMu.RUnlock()

	for _, sub := range fc.subscribers {
		select {
		case sub <- update:
		default:
			select {
			case <-sub:
			default:
			}
			sub <- update
			log.Warnf("[FastExecutionChannel] Subscriber channel was full, drained oldest message")
		}
	}
}

// GetSubscriberCount returns the number of active subscribers.
func (fc *FastExecutionChannel) GetSubscriberCount() int {
	fc.subMu.RLock()
	defer fc.subMu.RUnlock()
	return len(fc.subscribers)
}

// IsSubscribed returns whether we're subscribed to the fast execution channel.
func (fc *FastExecutionChannel) IsSubscribed() bool {
	fc.subStateMu.Lock()
	defer fc.subStateMu.Unlock()
	return fc.subscribed
}
