package message

import "sync"

var _ interface {
	Queue
	Iterator
} = (*inMemoryQueue)(nil)

// inMemoryQueue implements a queue that stores messages in memory. Messages
// can be enqueued concurrently but the returned iterator is not safe for
// concurrent access.
type inMemoryQueue struct {
	mu   sync.Mutex
	msgs []Message

	latchedMsg Message
}

// NewInMemoryQueue creates a new in-memory queue instance. This function can
// serve as a QueueFactory.
func NewInMemoryQueue() Queue {
	return new(inMemoryQueue)
}

// Close implements Queue.
func (q *inMemoryQueue) Close() error {
	return nil
}

// Enqueue implements Queue.
func (q *inMemoryQueue) Enqueue(msg Message) error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.msgs = append(q.msgs, msg)
	return nil
}

// PendingMessages implements Queue.
func (q *inMemoryQueue) PendingMessages() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	pending := len(q.msgs) != 0
	return pending
}

// DiscardMessages implements Queue.
func (q *inMemoryQueue) DiscardMessages() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.msgs = q.msgs[:0]
	q.latchedMsg = nil
	return nil
}

// Messages implements Queue.
func (q *inMemoryQueue) Messages() Iterator {
	return q
}

// Next implements Iterator.
func (q *inMemoryQueue) Next() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	qLen := len(q.msgs)
	if qLen == 0 {
		return false
	}

	// Dequeue message from the tail of the queue.
	q.latchedMsg = q.msgs[qLen-1]
	q.msgs = q.msgs[:qLen-1]
	return true
}

// Message implements Iterator.
func (q *inMemoryQueue) Message() Message {
	q.mu.Lock()
	defer q.mu.Unlock()

	msg := q.latchedMsg
	return msg
}

// Error implements Iterator.
func (q *inMemoryQueue) Error() error {
	return nil
}
