package exercises

// Logger is configured through the functional options pattern below.
type Logger struct {
	Prefix   string
	Level    string
	UseColor bool
}

// Option mutates a Logger during construction. Each WithX returns a closure over
// its argument.
type Option func(*Logger)

// WithPrefix sets the Logger's prefix.
//
// TODO: return an Option that sets l.Prefix = prefix.
func WithPrefix(prefix string) Option {
	return func(l *Logger) {
		l.Prefix = prefix
	}
}

// WithLevel sets the Logger's level.
//
// TODO: return an Option that sets l.Level = level.
func WithLevel(level string) Option {
	return func(l *Logger) {
		l.Level = level
	}
}

// WithColor sets whether the Logger uses colour.
//
// TODO: return an Option that sets l.UseColor = on.
func WithColor(on bool) Option {
	return func(l *Logger) {
		l.UseColor = on
	}
}

// NewLogger builds a Logger from sensible defaults (Prefix "[LOG]", Level "info",
// UseColor false), then applies each option in order.
//
// TODO: start from the defaults, apply every opt, and return the logger.
func NewLogger(opts ...Option) *Logger {
	logger := &Logger{
		Prefix:   "[LOG]",
		Level:    "info",
		UseColor: false,
	}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}
