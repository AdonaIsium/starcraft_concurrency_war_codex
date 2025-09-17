package domain

import "time"

// EventType represents engine-emitted events for pub/sub.
type EventType string

const (
    EventWorldTick     EventType = "WORLD_TICK"
    EventSquadMoved    EventType = "SQUAD_MOVED"
    EventCombatResolved EventType = "COMBAT_RESOLVED"
    EventSquadCreated  EventType = "SQUAD_CREATED"
)

// Event is a generic published struct.
// For production, you may use a typed union or dedicated structs per event.
type Event struct {
    ID        ID
    Type      EventType
    At        time.Time
    // Payload: include fields relevant to the event type.
    // Serialize to JSON when using external pub/sub.
    Payload any
}

