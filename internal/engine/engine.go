package engine

import (
    "context"
    "time"

    "github.com/yourname/starcraft-concurrency-war/internal/domain"
    "github.com/yourname/starcraft-concurrency-war/internal/logging"
    "github.com/yourname/starcraft-concurrency-war/internal/pubsub"
    "github.com/yourname/starcraft-concurrency-war/internal/storage"
)

// Config controls tick rate and future engine options.
type Config struct {
    TickDuration time.Duration // how long each tick lasts (e.g., 200ms)
}

// Engine exposes methods for submitting commands and querying snapshots.
// Concurrency contract:
// - All state mutations happen inside the engine goroutine.
// - External callers submit commands through a channel-safe method.
// - Reads can be served via a snapshot that is updated per tick.
type Engine interface {
    Start(ctx context.Context) error
    Stop()

    Submit(cmd domain.Command) error
    Snapshot() domain.Snapshot
}

// engine holds internal state and dependencies for the game loop.
// Fields to include:
// - cfg: Config
// - repo: storage.Repository
// - bus: pubsub.Bus
// - logger: logging.Logger
// - snapshot: domain.Snapshot (latest public snapshot)
// - commandCh: chan domain.Command (incoming commands)
// - stopCh: chan struct{} (shutdown signal)
// - mu: sync.RWMutex to guard snapshot access
type engine struct{}

// New constructs an engine.
// What to implement:
// - Initialize in-memory world state (via repo or defaults).
// - Build an initial snapshot.
// - Prepare channels (buffer sizes depend on expected throughput).
func New(cfg Config, repo storage.Repository, bus pubsub.Bus, logger logging.Logger) Engine {
    // Implement engine construction here.
    // Steps:
    // - Initialize internal fields (cfg, repo, bus, logger).
    // - Create buffered channels for commandCh and stopCh.
    // - Initialize an initial snapshot (empty or from repo.LoadWorld).
    // - Consider loading world state via repo and building a starting snapshot.
    // Return an Engine implementation (e.g., &engine{...}).
    return nil
}

// Start runs the main tick loop.
// Steps to implement:
// 1) Load/initialize world state (from repo, or create a new World with size).
// 2) Create a time.Ticker with cfg.TickDuration.
// 3) On each tick:
//    - Drain pending commands and apply them to the world state (validate auth/ownership).
//    - Advance simulation: movement, scouting visibility, combat resolution.
//    - Persist state periodically (e.g., every N ticks) via repo.
//    - Publish events (tick, moved, combat, squad created) via bus.
//    - Update the public snapshot (use e.mu to protect).
// 4) Handle ctx.Done or Stop() to gracefully exit.
func (e *engine) Start(ctx context.Context) error {
    // Implement the engine tick loop here.
    // Packages to use:
    // - time: time.NewTicker for fixed tick duration (cfg.TickDuration).
    // - context: select on ctx.Done() for cancellation.
    // Steps each tick:
    // 1) Drain commands from commandCh and validate them (ownership, params).
    // 2) Apply commands to world state (move, scout, queue attacks).
    // 3) Advance simulation systems: movement resolution, combat calculation, fog-of-war.
    // 4) Persist world state periodically via repo.SaveWorld (e.g., every N ticks).
    // 5) Publish domain events via bus.Publish (EventWorldTick, EventSquadMoved, etc.).
    // 6) Update e.snapshot under a write lock.
    // 7) Handle graceful shutdown when stopCh or ctx.Done() fires.
    return nil
}

// Stop requests the engine to shut down.
func (e *engine) Stop() {
    // Implement shutdown signaling here.
    // Steps:
    // - Attempt a non-blocking send on stopCh to request loop exit.
    // - Optionally log if already stopping.
}

// Submit enqueues a command into the engine.
// Steps to implement:
// - Validate command shape (non-nil target, etc.) quickly if needed.
// - Non-blocking send: consider select with default to avoid blocking callers.
// - Return an error if the buffer is full; caller can retry or fail.
func (e *engine) Submit(cmd domain.Command) error {
    // Implement non-blocking command enqueue here.
    // Steps:
    // - Validate minimal shape (e.g., required fields based on Type).
    // - Use select with default to try sending on commandCh.
    // - If full, return an error so caller can retry or fail.
    return nil
}

// Snapshot returns a copy of the latest snapshot.
// Steps to implement:
// - Use e.mu.RLock to read the current snapshot.
// - Return a deep copy if external mutation is a concern.
func (e *engine) Snapshot() domain.Snapshot {
    // Implement snapshot access here.
    // Steps:
    // - Acquire a read lock on e.mu.
    // - Return a copy of e.snapshot to prevent external mutation.
    // - Consider deep copy for nested slices/maps.
    return domain.Snapshot{}
}
