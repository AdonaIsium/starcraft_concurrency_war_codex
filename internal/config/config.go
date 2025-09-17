package config

import "time"

// Config holds server, engine, DB, and pubsub configuration.
type Config struct {
    HTTPAddr     string        // e.g., ":8080"; set by env HTTP_ADDR
    TickDuration time.Duration // e.g., 200ms; set by env TICK_MS

    // DB configuration (used if you switch to SQL repository)
    DB DBConfig
}

// DBConfig holds database connection parameters.
type DBConfig struct {
    Driver   string // e.g., "postgres"
    URL      string // full DSN or URL
    MaxConns int
}

// Load reads environment variables and sets defaults.
// What to implement:
// - Read env vars (HTTP_ADDR, TICK_MS, DB_DRIVER, DB_URL, DB_MAX_CONNS).
// - Convert TICK_MS into time.Duration.
// - Provide sane defaults when env vars are missing (":8080", 200ms, etc.).
// - Optional: load from a .env using a library like github.com/joho/godotenv.
func Load() Config {
    // Implement environment loading here.
    // Packages to use:
    // - os: to read env vars like HTTP_ADDR, TICK_MS, DB_DRIVER, DB_URL, DB_MAX_CONNS.
    // - strconv: to parse integers (e.g., TICK_MS to int, DB_MAX_CONNS to int).
    // - time: to convert milliseconds to time.Duration, e.g., time.Duration(ms) * time.Millisecond.
    // Optional:
    // - github.com/joho/godotenv to load a .env file for local development.
    // Steps:
    // 1) Read HTTP_ADDR (default to ":8080").
    // 2) Read TICK_MS (default to 200), convert to duration for TickDuration.
    // 3) Read DB_* values and fill DBConfig fields; set sensible defaults.
    // 4) Return a Config struct populated with these values.
    return Config{}
}
