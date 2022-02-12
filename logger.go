package goplogadapter

import (
	"context"
	"fmt"

	"github.com/axpira/gop/log"
)

type contextKey string

const (
	loggerContextKey contextKey = "goplogadapter"
)

type FormatterFunc func(log.Level, string, error, map[string]interface{})

func WithLevel(level log.Level) log.LoggerOption {
	return log.LoggerOptionFunc(func(l log.Logger) log.Logger {
		l1 := l.(logger)
		l1.level = level
		return l1
	})
}

func WithFormatter(formatter FormatterFunc) log.LoggerOption {
	return log.LoggerOptionFunc(func(l log.Logger) log.Logger {
		l1 := l.(logger)
		l1.Formatter = formatter
		return l1
	})
}

type logger struct {
	Formatter FormatterFunc
	level     log.Level
	fielder   *field
}

func New(opts ...log.LoggerOption) log.Logger {
	return newLogger(newField()).With(opts...)
}

func newLogger(fielder *field) log.Logger {
	l := logger{fielder: fielder}
	return l
}

func (l logger) clone() logger {
	return logger{
		Formatter: l.Formatter,
		level:     l.level,
		fielder:   l.fielder.clone(),
	}
}

func (l logger) Level() log.Level {
	return l.level
}

func (l logger) HasLevel(lv log.Level) bool {
	return lv >= l.level
}

func (l logger) Log(level log.Level, builder log.FieldBuilder) {
	if !l.HasLevel(level) {
		return
	}
	fielder := builder.(*field)
	fielder.Fields(l.fielder.m)
	fielder.send(level, l.Formatter)
}

func (l logger) With(opts ...log.LoggerOption) log.Logger {
	var l1 log.Logger = l.clone()
	for i := range opts {
		l1 = opts[i].Update(l1)
	}
	return l1
}

func (l logger) NewFieldBuilder() log.FieldBuilder {
	return newField()
}

func (l logger) Trc(builder log.FieldBuilder) {
	l.Log(log.TraceLevel, builder)
}

func (l logger) Trace(msg string) {
	l.Log(log.TraceLevel, newField().Msg(msg))
}

func (l logger) Tracef(format string, args ...interface{}) {
	l.Log(log.TraceLevel, newField().Msgf(format, args...))
}

func (l logger) Dbg(builder log.FieldBuilder) {
	l.Log(log.DebugLevel, builder)
}

func (l logger) Debug(msg string) {
	l.Log(log.DebugLevel, newField().Msg(msg))
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.Log(log.DebugLevel, newField().Msgf(format, args...))
}

func (l logger) Inf(builder log.FieldBuilder) {
	l.Log(log.InfoLevel, builder)
}

func (l logger) Info(msg string) {
	l.Log(log.InfoLevel, newField().Msg(msg))
}

func (l logger) Infof(format string, args ...interface{}) {
	l.Log(log.InfoLevel, newField().Msgf(format, args...))
}

func (l logger) Wrn(builder log.FieldBuilder) {
	l.Log(log.WarnLevel, builder)
}

func (l logger) Warn(msg string) {
	l.Log(log.WarnLevel, newField().Msg(msg))
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.Log(log.WarnLevel, newField().Msgf(format, args...))
}

func (l logger) Err(builder log.FieldBuilder) {
	l.Log(log.ErrorLevel, builder)
}

func (l logger) Error(msg string, err error) {
	l.Log(log.ErrorLevel, newField().Err(err).Msg(msg))
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.Log(log.ErrorLevel, newField().Msgf(format, args...))
}

func (l logger) Ftl(builder log.FieldBuilder) {
	l.Log(log.FatalLevel, builder)
}

func (l logger) Fatal(msg string) {
	l.Log(log.FatalLevel, newField().Msg(msg))
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.Log(log.FatalLevel, newField().Msgf(format, args...))
}

func (l logger) Pnc(builder log.FieldBuilder) {
	l.Log(log.PanicLevel, builder)
}

func (l logger) Panic(msg string) {
	l.Log(log.PanicLevel, newField().Msg(msg))
}

func (l logger) Panicf(format string, args ...interface{}) {
	l.Log(log.PanicLevel, newField().Msgf(format, args...))
}

func (l logger) Print(args ...interface{}) {
	l.Log(log.InfoLevel, newField().Msg(fmt.Sprint(args...)))
}

func (l logger) Printf(format string, args ...interface{}) {
	l.Log(log.InfoLevel, newField().Msgf(format, args...))
}

func (l logger) Println(args ...interface{}) {
	l.Log(log.InfoLevel, newField().Msg(fmt.Sprint(args...)))
}

func (l logger) Write(msg []byte) (int, error) {
	l.Log(log.InfoLevel, newField().Msg(string(msg)))
	return 0, nil
}

func (l logger) ToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}

func (l logger) FromCtx(ctx context.Context) log.Logger {
	v := ctx.Value(loggerContextKey)
	if v == nil {
		return l
	}
	return v.(log.Logger)
}
