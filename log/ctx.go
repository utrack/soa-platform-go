package log

import (
	"context"

	"github.com/utrack/soa-platform-go/log/logimpl"
)

type typeCtxLoggerKey string

// thought about swapping words Logger and Key here, but this probably will
// flare up some code security checks :^)
const ctxLoggerKey = "kvs"

var global logimpl.Logger

func toCtx(ctx context.Context, l logimpl.Logger) context.Context {
	return context.WithValue(ctx, ctxLoggerKey, l)
}

func fromCtx(ctx context.Context) logimpl.Logger {
	v, _ := ctx.Value(ctxLoggerKey).(logimpl.Logger)
	if v != nil {
		return v
	}
	return logimpl.Global()
}
