/*Package logimpl describes a generic logger implementation.


This package does not have any log implementation enabled by default.
Import log/logimpl/zapimpl to hook up zap's development logger, or set up your own
via logimpl.SetGlobal.*/
package logimpl

// RootLogger is a root-level logger that was not wrapped via
// With() or WithLevel() calls.
type RootLogger interface {
	Logger
	SetLevel(Level)
}

// Logger is a generic logger with additional key-value fields attached to every log entry.
type Logger interface {
	Debugw(string, ...interface{})
	Infow(string, ...interface{})
	Warnw(string, ...interface{})
	Errorw(string, ...interface{})
	Fatalw(string, ...interface{})

	// With creates new Logger that inherits this one, but adds new
	// key-value fields to every log entry that's emitted.
	//
	// Current Logger is not mutated in any way.
	With(kvs ...interface{}) Logger

	// WithLevel changes logging level for a child logger.
	//	WithLevel(Level) Logger TODO hard to implement w/ zap w/o swapping zapcore
}

var global RootLogger

// Global returns a global logger.
func Global() RootLogger {
	if global == nil {
		panic("global logger is not set")
	}
	return global
}

func SetGlobal(l RootLogger) {
	global = l
}

type Level uint

const (
	LevelDebug Level = 0
	LevelInfo  Level = 1
	LevelWarn  Level = 2
	LevelError Level = 3
	LevelFatal Level = 4
)
