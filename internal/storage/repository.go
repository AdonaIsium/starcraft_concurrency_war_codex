package storage

import (
    "context"

    "github.com/yourname/starcraft-concurrency-war/internal/domain"
)

// ErrNotFound should be returned when an entity cannot be found.
// Use errors.Is(err, ErrNotFound) to test.
var ErrNotFound = Err("not found")

// Err is a simple string error type.
type Err string

func (e Err) Error() string { return string(e) }

// Repository defines persistence operations for players, squads, units, and world snapshots.
type Repository interface {
    // Players
    CreatePlayer(ctx context.Context, p domain.Player) error
    GetPlayer(ctx context.Context, id domain.ID) (domain.Player, error)

    // Squads
    CreateSquad(ctx context.Context, s domain.Squad) error
    GetSquad(ctx context.Context, id domain.ID) (domain.Squad, error)
    UpdateSquad(ctx context.Context, s domain.Squad) error

    // World
    SaveWorld(ctx context.Context, w domain.World) error
    LoadWorld(ctx context.Context) (domain.World, error)
}
