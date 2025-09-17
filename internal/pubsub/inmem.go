package pubsub

import "context"

// inMemoryBus is a simple channel-based fan-out.
// This is suitable for local dev and learning purposes.
// inMemoryBus should maintain subscriber lists per topic with buffered channels.
// Fields to include:
// - mu: sync.RWMutex for protecting maps/slices
// - subs: map[Topic][]chan Message to hold subscriber channels
type inMemoryBus struct{}

// NewInMemoryBus constructs a simple in-memory bus.
func NewInMemoryBus() Bus {
    // Implement construction here.
    // Steps:
    // - Initialize the subs map.
    // - Return &inMemoryBus{...} as Bus.
    return nil
}

// Publish sends a message to all topic subscribers.
// Steps to implement:
// - Lock and copy the slice of channels.
// - Non-blocking send (select with default) to avoid publisher stalls.
// - Consider dropping messages or using buffered channels.
func (b *inMemoryBus) Publish(ctx context.Context, msg Message) error {
    // Implement fan-out publish here.
    // Steps:
    // - Read-lock subs for the topic and copy the slice.
    // - Iterate and attempt non-blocking send to each buffered subscriber channel.
    // - Optionally drop or track dropped messages if a subscriber is slow.
    return nil
}

// Subscribe registers a subscriber to a topic.
// Returns a receive-only channel and an unsubscribe function.
func (b *inMemoryBus) Subscribe(ctx context.Context, topic Topic) (<-chan Message, func(), error) {
    // Implement subscription logic here.
    // Steps:
    // - Create a buffered channel for messages.
    // - Append it to the subs list for the topic under a write lock.
    // - Return the receive-only channel and an unsubscribe function that
    //   removes the channel and closes it. Also hook ctx.Done() to auto-unsubscribe.
    return nil, nil, nil
}
