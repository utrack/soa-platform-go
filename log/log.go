/*Package log provides global logging functions for use within the app.

See log/logimpl for a description of logger's implementation.*/
package log

import (
	"context"

	// to set up default zap's developer mode implementation
	_ "github.com/utrack/soa-platform-go/log/logimpl/zapimpl"
)

// With adds new metadata to a context.
// Any call using child context will print these kvs
// along with the message itself.
//
// Parent context does not get changed, so you can create
// many different child contexts from a single parent.
func With(ctx context.Context, kvs ...interface{}) context.Context {
	l := fromCtx(ctx)
	l = l.With(kvs...)
	return toCtx(ctx, l)
}

// type Level = logimpl.Level

// const (
// 	LevelDebug Level = logimpl.LevelDebug
// 	LevelInfo  Level = logimpl.LevelInfo
// 	LevelWarn  Level = logimpl.LevelWarn
// 	LevelError Level = logimpl.LevelError
// 	LevelFatal Level = logimpl.LevelFatal
// )

// WithLevel changes logging level for this context.
// func WithLevel(ctx context.Context, l Level) context.Context {
// 	return toCtx(ctx, fromCtx(ctx).WithLevel(l))
// }

func Debug(ctx context.Context, msg string, kvs ...interface{}) {
	fromCtx(ctx).Debugw(msg, kvs...)
}

func Info(ctx context.Context, msg string, kvs ...interface{}) {
	fromCtx(ctx).Infow(msg, kvs...)
}

func Warn(ctx context.Context, msg string, kvs ...interface{}) {
	fromCtx(ctx).Warnw(msg, kvs...)
}

func Warne(ctx context.Context, msg string, err error, kvs ...interface{}) {
	Warn(ctx, msg, append(kvs, "error", err)...)
}

func Error(ctx context.Context, msg string, err error, kvs ...interface{}) {
	Errorn(ctx, msg, append(kvs, "error", err)...)
}

func Errorn(ctx context.Context, msg string, kvs ...interface{}) {
	fromCtx(ctx).Errorw(msg, kvs...)
}

func Fatal(ctx context.Context, msg string, kvs ...interface{}) {
	fromCtx(ctx).Fatalw(msg, kvs...)
}

func Fatale(ctx context.Context, msg string, err error, kvs ...interface{}) {
	Fatal(ctx, msg, append(kvs, "error", err)...)
}
