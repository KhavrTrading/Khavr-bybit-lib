package bybit_ws

import (
	"sync"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws/bybit_ws_models"

	log "github.com/sirupsen/logrus"
)

// ExecutionChannel manages execution (fill) updates from Bybit private WebSocket
// and distributes them to multiple subscribers via channels.
type ExecutionChannel struct {
	client *BybitWsClient

	subscribers []chan bybit_ws_models.ExecutionEvent
	subMu       sync.RWMutex

	topic      string
	subscribed bool
	subStateMu sync.Mutex
}

// NewExecutionChannel creates a new execution channel manager.
func NewExecutionChannel(client *BybitWsClient) *ExecutionChannel {
	ec := &ExecutionChannel{
		client:      client,
		subscribers: make([]chan bybit_ws_models.ExecutionEvent, 0),
	}
	client.SetExecutionCallback(ec.handleUpdate)
	return ec
}

// Subscribe returns a buffered channel that will receive execution events.
func (ec *ExecutionChannel) Subscribe() chan bybit_ws_models.ExecutionEvent {
	ch := make(chan bybit_ws_models.ExecutionEvent, 10)

	ec.subMu.Lock()
	ec.subscribers = append(ec.subscribers, ch)
	ec.subMu.Unlock()

	log.Infof("[ExecutionChannel] New subscriber added, total: %d", len(ec.subscribers))
	return ch
}

// Unsubscribe removes a subscriber channel.
func (ec *ExecutionChannel) Unsubscribe(ch chan bybit_ws_models.ExecutionEvent) {
	ec.subMu.Lock()
	defer ec.subMu.Unlock()

	for i, sub := range ec.subscribers {
		if sub == ch {
			close(ch)
			ec.subscribers = append(ec.subscribers[:i], ec.subscribers[i+1:]...)
			log.Infof("[ExecutionChannel] Subscriber removed, total: %d", len(ec.subscribers))
			return
		}
	}
}

// SubscribeToExecutions subscribes to execution updates on the WebSocket.
// topic: "execution" (all-in-one) or "execution.linear", "execution.spot", etc.
func (ec *ExecutionChannel) SubscribeToExecutions(topic string) error {
	ec.subStateMu.Lock()
	defer ec.subStateMu.Unlock()

	if ec.subscribed {
		log.Warnf("[ExecutionChannel] Already subscribed to execution channel")
		return nil
	}

	ec.topic = topic
	if err := ec.client.SubscribeExecution(topic); err != nil {
		return err
	}

	ec.subscribed = true
	log.Infof("[ExecutionChannel] Subscribed to %s", topic)
	return nil
}

func (ec *ExecutionChannel) handleUpdate(update bybit_ws_models.ExecutionEvent) {
	ec.subMu.RLock()
	defer ec.subMu.RUnlock()

	for _, sub := range ec.subscribers {
		select {
		case sub <- update:
		default:
			select {
			case <-sub:
			default:
			}
			sub <- update
			log.Warnf("[ExecutionChannel] Subscriber channel was full, drained oldest message")
		}
	}
}

// GetSubscriberCount returns the number of active subscribers.
func (ec *ExecutionChannel) GetSubscriberCount() int {
	ec.subMu.RLock()
	defer ec.subMu.RUnlock()
	return len(ec.subscribers)
}

// IsSubscribed returns whether we're subscribed to the execution channel.
func (ec *ExecutionChannel) IsSubscribed() bool {
	ec.subStateMu.Lock()
	defer ec.subStateMu.Unlock()
	return ec.subscribed
}
