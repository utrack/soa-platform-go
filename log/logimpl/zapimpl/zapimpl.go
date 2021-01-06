/*Package zapimpl implements logimpl.RootLogger using uber/zap logger.

logimpl.Global is set to zap's DevelopmentConfig logger during init.*/
package zapimpl

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/utrack/soa-platform-go/log/logimpl"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RootLogger struct {
	impl
	atom zap.AtomicLevel
}

var _ logimpl.Logger = impl{}
var _ logimpl.RootLogger = &RootLogger{}

type impl struct {
	sl *zap.SugaredLogger
}

func init() {
	cfg := zap.NewDevelopmentConfig()
	rl, err := New(cfg)
	if err != nil {
		panic(err)
	}
	logimpl.SetGlobal(rl)
}

func New(cfg zap.Config) (*RootLogger, error) {

	atom := zap.NewAtomicLevel()
	cfg.Level = atom

	logger, err := cfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "can't build logger")
	}
	ret := &RootLogger{
		atom: atom,
		impl: impl{
			sl: logger.Sugar(),
		},
	}
	return ret, nil
}

var mapLevels = map[logimpl.Level]zapcore.Level{
	logimpl.LevelDebug: zapcore.DebugLevel,
	logimpl.LevelInfo:  zapcore.InfoLevel,
	logimpl.LevelWarn:  zapcore.WarnLevel,
	logimpl.LevelError: zapcore.ErrorLevel,
	logimpl.LevelFatal: zapcore.FatalLevel,
}

func (r *RootLogger) SetLevel(il logimpl.Level) {
	l, ok := mapLevels[il]
	if !ok {
		panic(fmt.Sprintf("unknown level %v", il))
	}
	r.atom.SetLevel(l)
}

func (i impl) Debugw(msg string, args ...interface{}) {
	i.sl.Debugw(msg, args...)
}
func (i impl) Infow(msg string, args ...interface{}) {
	i.sl.Infow(msg, args...)
}
func (i impl) Warnw(msg string, args ...interface{}) {
	i.sl.Warnw(msg, args...)
}
func (i impl) Errorw(msg string, args ...interface{}) {
	i.sl.Errorw(msg, args...)
}
func (i impl) Fatalw(msg string, args ...interface{}) {
	i.sl.Fatalw(msg, args...)
}
func (i impl) With(kvs ...interface{}) logimpl.Logger {
	return impl{
		sl: i.sl.With(kvs...),
	}
}
