package main

// main is the program entrypoint.
// Implement the following here (without moving logic into handlers):
// - Load config via internal/config.Load (reads env vars like HTTP_ADDR, TICK_MS).
// - Initialize a logger via internal/logging.New or a structured logger (zap/zerolog).
// - Choose a storage implementation: start with storage.NewInMemoryRepository(); later wire a SQL repo.
// - Initialize a pub/sub bus: start with pubsub.NewInMemoryBus(); later Redis/NATS/Kafka.
// - Create an engine via engine.New with Config (TickDuration from config), repo, bus, logger.
// - Start the engine in a goroutine (use context.WithCancel + signal.Notify for graceful shutdown).
//   Packages to use: context, os/signal, syscall, time.
// - Build HTTP router via api.NewRouter with dependencies; serve via net/http.Server.
// - Handle shutdown: on SIGINT/SIGTERM, call Server.Shutdown with a timeout context.
func main() {
    // Intentionally left empty. See comments above for exact steps and packages.
}
