package goplogadapter

import (
	"errors"
	"testing"

	"github.com/axpira/gop/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoggerLevels(t *testing.T) {
	tests := []struct {
		logLevel      log.Level
		allowedLevels []log.Level
		denyLevels    []log.Level
	}{
		{
			logLevel:      log.TraceLevel,
			allowedLevels: []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{},
		},
		{
			logLevel:      log.DebugLevel,
			allowedLevels: []log.Level{log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{log.TraceLevel},
		},
		{
			logLevel:      log.InfoLevel,
			allowedLevels: []log.Level{log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{log.TraceLevel, log.DebugLevel},
		},
		{
			logLevel:      log.WarnLevel,
			allowedLevels: []log.Level{log.WarnLevel, log.ErrorLevel, log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel},
		},
		{
			logLevel:      log.ErrorLevel,
			allowedLevels: []log.Level{log.ErrorLevel, log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel},
		},
		{
			logLevel:      log.FatalLevel,
			allowedLevels: []log.Level{log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel},
		},
		{
			logLevel:      log.PanicLevel,
			allowedLevels: []log.Level{log.PanicLevel},
			denyLevels:    []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel},
		},
		{
			logLevel:      log.DisabledLevel,
			allowedLevels: []log.Level{},
			denyLevels:    []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel, log.PanicLevel},
		},
		{
			logLevel:      log.NoLevel,
			allowedLevels: []log.Level{log.TraceLevel, log.DebugLevel, log.InfoLevel, log.WarnLevel, log.ErrorLevel, log.FatalLevel, log.PanicLevel},
			denyLevels:    []log.Level{},
		},
	}
	for _, tc := range tests {
		called := 0
		log := New(WithLevel(tc.logLevel), WithFormatter(func(level log.Level, msg string, err error, m map[string]interface{}) {
			called++
			assert.Nil(t, err)
			assert.Equal(t, "Hello World", msg)
			assert.Equal(t, make(map[string]interface{}), m)
		}))
		require.Equal(t, tc.logLevel, log.Level())
		for _, level := range tc.allowedLevels {
			called = 0
			log.Log(level, log.NewFieldBuilder().Msg("Hello World"))
			assert.Equalf(t, called, 1, "when log level is %v and send level %v expected to be called send", tc.logLevel, level)

			if fn, ok := loggerFuncs[level]; ok {
				called = 0
				fn(log, log.NewFieldBuilder().Msg("Hello World"), "Hello World")
				assert.Equalf(t, called, 3, "when log level is %v and send level %v expected to be called send", tc.logLevel, level)
			}
		}
		for _, level := range tc.denyLevels {
			called = 0
			log.Log(level, log.NewFieldBuilder().Msg("Hello World"))
			assert.Equalf(t, called, 0, "when log level is %v and send level %v expected not to be called send", tc.logLevel, level)

			if fn, ok := loggerFuncs[level]; ok {
				called = 0
				fn(log, log.NewFieldBuilder().Msg("Hello World"), "Hello World")
				assert.Equalf(t, called, 0, "when log level is %v and send level %v expected to be called send", tc.logLevel, level)
			}
		}
	}
}

var loggerFuncs = map[log.Level]func(log.Logger, log.FieldBuilder, string){
	log.TraceLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Trc(fb)
		l.Trace(msg)
		l.Tracef(msg)
	},
	log.DebugLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Dbg(fb)
		l.Debug(msg)
		l.Debugf(msg)
	},
	log.InfoLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Inf(fb)
		l.Info(msg)
		l.Infof(msg)
	},
	log.WarnLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Wrn(fb)
		l.Warn(msg)
		l.Warnf(msg)
	},
	log.ErrorLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Err(fb)
		l.Error(msg, nil)
		l.Errorf(msg)
	},
	log.FatalLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Ftl(fb)
		l.Fatal(msg)
		l.Fatalf(msg)
	},
	log.PanicLevel: func(l log.Logger, fb log.FieldBuilder, msg string) {
		l.Pnc(fb)
		l.Panic(msg)
		l.Panicf(msg)
	},
}

func TestUpdate(t *testing.T) {
	l := New(WithLevel(log.InfoLevel), WithFormatter(func(level log.Level, msg string, err error, m map[string]interface{}) {
		assert.Equal(t, level, log.InfoLevel)
		assert.Equal(t, msg, "hello world")
		assert.Equal(t, m, map[string]interface{}{
			"key_str": "value",
		})
		assert.Equal(t, err, errors.New("unknown error"))
	}))

	l1 := l.With(l.NewFieldBuilder().Str("key_str", "value"))
	l1.Inf(l1.NewFieldBuilder().Msg("hello world").Err(errors.New("unknown error")))

	l2 := l1.With(WithFormatter(func(level log.Level, msg string, err error, m map[string]interface{}) {
		assert.Equal(t, level, log.InfoLevel)
		assert.Equal(t, msg, "Hello")
		assert.Equal(t, m, map[string]interface{}{
			"key_str":  "value",
			"key_uint": uint(99),
			"key_int":  42,
		})
		assert.Equal(t, err, nil)
	}))
	l2 = l2.With(l.NewFieldBuilder().Int("key_int", 42))
	l2.Inf(l.NewFieldBuilder().Msg("Hello").Uint("key_uint", uint(99)))
}
