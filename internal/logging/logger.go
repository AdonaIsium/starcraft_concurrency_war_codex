package logging

// Logger is a minimal logging abstraction.
// You can replace with a structured logger (zap, zerolog) later.
type Logger interface {
    Infof(format string, args ...any)
    Errorf(format string, args ...any)
}

// New returns a simple stdlib logger.
func New() Logger {
    // Implement a logger provider here.
    // Options:
    // - Use the standard library log package: log.New(os.Stdout, "", log.LstdFlags).
    // - Or integrate a structured logger like uber-go/zap or rs/zerolog.
    // If using zap:
    //   logger, _ := zap.NewProduction(); return logger.Sugar()
    // If using zerolog:
    //   l := zerolog.New(os.Stdout).With().Timestamp().Logger(); return l
    // Return a type that implements Infof and Errorf.
    return nil
}

// Implement a concrete type that satisfies Logger with Infof/Errorf methods.
// For stdlib log, wrap a *log.Logger and format messages accordingly.
