package storage

import (
    "context"

    "github.com/yourname/starcraft-concurrency-war/internal/domain"
)

// InMemoryRepository stores everything in maps for easy iteration.
// Concurrency note: keep locking simple and coarse; this is just for learning.
type InMemoryRepository struct {
    // Fields to include:
    // - mu: sync.RWMutex to guard maps
    // - players: map[domain.ID]domain.Player
    // - squads: map[domain.ID]domain.Squad
    // - world: *domain.World (latest saved)
}

// NewInMemoryRepository constructs the repo.
func NewInMemoryRepository() *InMemoryRepository {
    // Implement construction here.
    // Steps:
    // - Initialize maps for players and squads.
    // - Return &InMemoryRepository{...}.
    return nil
}

func (r *InMemoryRepository) CreatePlayer(ctx context.Context, p domain.Player) error {
    // Implement in-memory create here.
    // Steps:
    // - Lock for writing.
    // - If CreatedAt zero, set to time.Now().UTC().
    // - Insert into players map.
    // - Return nil.
    return nil
}

func (r *InMemoryRepository) GetPlayer(ctx context.Context, id domain.ID) (domain.Player, error) {
    // Implement read here.
    // Steps:
    // - RLock, lookup in players map.
    // - If not found, return ErrNotFound.
    // - Else, return player.
    return domain.Player{}, nil
}

func (r *InMemoryRepository) CreateSquad(ctx context.Context, s domain.Squad) error {
    // Implement in-memory create for squads here.
    // Steps similar to CreatePlayer; set CreatedAt and insert into squads map.
    return nil
}

func (r *InMemoryRepository) GetSquad(ctx context.Context, id domain.ID) (domain.Squad, error) {
    // Implement read for squads here.
    // Steps:
    // - RLock, lookup in squads map; return ErrNotFound if missing.
    return domain.Squad{}, nil
}

func (r *InMemoryRepository) UpdateSquad(ctx context.Context, s domain.Squad) error {
    // Implement update for squads here.
    // Steps:
    // - Lock, check existence; if missing return ErrNotFound; else overwrite.
    return nil
}

func (r *InMemoryRepository) SaveWorld(ctx context.Context, w domain.World) error {
    // Implement save world here.
    // Steps:
    // - Lock and assign r.world = &w (consider deep copy if needed).
    return nil
}

func (r *InMemoryRepository) LoadWorld(ctx context.Context) (domain.World, error) {
    // Implement load world here.
    // Steps:
    // - RLock; if r.world == nil, decide to return ErrNotFound or a default world.
    // - Else, return a copy of *r.world.
    return domain.World{}, nil
}
