package pubsub

import "context"

// Topic lets you categorize events (e.g., "world", "combat", "squad").
type Topic string

// Message is a generic wrapper for payloads.
type Message struct {
    Topic   Topic
    Payload any
}

// Bus defines pub/sub operations.
// Implementations: InMemory (here), or Redis/NATS/Kafka.
type Bus interface {
    Publish(ctx context.Context, msg Message) error
    Subscribe(ctx context.Context, topic Topic) (<-chan Message, func(), error)
}
