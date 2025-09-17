StarCraft Concurrency War (Go Learning Project)

Overview
- Goal: Learn Go concurrency, HTTP (GET/POST), pub/sub, and database persistence by building a simple multiplayer, tick-based world where players manage squads of units.
- Important: This is a “no implementation” learning scaffold. Every function has a signature and detailed comments only. There is no code inside any function body. The project will not compile until you fill in logic. Use the comments as your step-by-step guide with suggested packages and patterns.

High-Level Architecture
- cmd/server: Program entrypoint that wires config, logger, storage, pubsub, engine, and HTTP API.
- internal/config: Loads configuration (ports, tick rate, DB connection, pubsub choice) from env.
- internal/logging: Simple logger abstraction. You can swap to a structured logger later.
- internal/domain: Core types (Player, Squad, Unit, Map, Command, Events) and interfaces shared across layers.
- internal/engine: Game loop, ticks, concurrency orchestration, command handling, world updates, and event emission.
- internal/storage: Repository interfaces and implementations (in-memory for learning, SQL skeleton for production).
- internal/pubsub: Publisher/Subscriber interfaces and in-memory implementation. You can add Redis/NATS later.
- internal/api: HTTP handlers, routing, request/response DTOs, and validation. Talks to services/engine.
- internal/services: Higher-level orchestration between storage, engine, pubsub.

Suggested Learning Path
1) Read domain models to understand world state and commands.
2) Read engine stubs to see how ticks, channels, and goroutines fit together.
3) Read pubsub and storage interfaces to see boundaries; start with in-memory.
4) Wire and run the server after filling in minimal pieces (e.g., in-memory repo).
5) Implement endpoints incrementally (create player, create squad, issue command, get world).
6) Add tests for the service layer and engine command handling.
7) Swap in SQL repository and a real pub/sub when ready.

Concurrency Model (at a glance)
- One central Engine goroutine drives ticks using a `time.Ticker`.
- Commands are submitted via a channel to the Engine.
- Engine updates world state per tick and emits events to a pub/sub.
- You can later experiment with per-squad goroutines (actors) for more granular concurrency.

HTTP API (starter endpoints)
- POST `/players`: Create a player.
- POST `/squads`: Create a squad for a player.
- POST `/commands`: Issue a command (move/attack/scout) for a squad.
- GET `/world`: Get a snapshot of world state.
- GET `/squads/{id}`: Get a squad by ID.

Database Path
- Start with the `inmem` repository to learn quickly.
- Move to `sql` repository using `database/sql` + `pgx` or `lib/pq` (recommended). Migrations are provided as placeholders in `migrations/` with comments.

Pub/Sub Path
- Start with `pubsub/inmem` (simple channel fan-out).
- Swap to Redis Streams, NATS, or Kafka when ready. Interfaces are already in place.

Filling in the Stubs
- Every function contains comments explaining exactly what to implement, what packages to consider, and any concurrency considerations.
- Keep business logic out of handlers; handlers validate/parse, then call services or engine.

Running the Project (after you fill in logic)
- Ensure Go 1.22+ is installed.
- Optional: Create a `.env` file or export env vars (see `internal/config` for keys).
- Run: `go run ./cmd/server`.

How To Use This Scaffold
- Open each file and locate functions with comments. The comments explain exactly what to do and which packages to use.
- Implement pieces in this order to reduce confusion:
  1) storage/inmem.go (basic CRUD in memory) so that the API/services have something to talk to.
  2) services/game_service.go (validate and orchestrate repository/engine calls).
  3) engine/engine.go (setup channels, implement Start loop, Submit, Snapshot).
  4) api/router.go and api/handlers.go (wire routes and parse requests).
  5) cmd/server/main.go (wire everything together and start the server/engine).
  6) Optionally swap in storage/sql.go with migrations/ when ready to learn SQL.

Packages You’ll Likely Use
- context: cancellation, timeouts in handlers and services.
- time: tick loop with time.Ticker, timestamps.
- net/http: HTTP server, handlers, status codes.
- encoding/json: request/response marshalling.
- os, os/signal, syscall: graceful shutdown from SIGINT/SIGTERM.
- sync: mutexes (protect in-memory maps and snapshots).
- github.com/google/uuid: generate IDs for players, squads, commands.
- database/sql + pgx/lib/pq: when implementing SQL repository.
- Optional: github.com/go-chi/chi for routing, github.com/joho/godotenv for .env, zap/zerolog for logging.

Next Steps / Exercises
- Implement move logic: pathing on the grid with simple rules.
- Add combat resolution per tick: damage, HP, death.
- Add fog-of-war: squads only see nearby tiles.
- Add rate limiting or auth to API.
- Add WebSocket endpoint to stream world events to clients.
- Add background workers for persistence snapshots every N ticks.

Why this structure?
- Clear boundaries: domain, engine, infrastructure (storage/pubsub), delivery (HTTP).
- Testable seams: interfaces let you mock engine/pubsub/repo in tests.
- Incremental learning: start in-memory + simple; then upgrade components.
