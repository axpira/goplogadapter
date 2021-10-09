package goplogadapter

import (
	"errors"
	"testing"
	"time"

	"github.com/axpira/gop/log"
	"github.com/stretchr/testify/assert"
)

func TestFields(t *testing.T) {
	now := time.Now()
	dur, _ := time.ParseDuration("1d2h3m4s")
	formatterFunc := func(level log.Level, msg string, err error, m map[string]interface{}) {
		assert.Nil(t, err)
		assert.Equal(t, "Hello World", msg)
		assert.Equal(t, map[string]interface{}{
			"key_str":        "value_str",
			"key_int64":      int64(-64),
			"key_bool":       true,
			"key_float32":    float32(3.2),
			"key_float64":    float64(6.4),
			"key_any":        "anything",
			"key_map_str":    "value",
			"key_map_int64":  int64(-84),
			"key_map_bool":   false,
			"key_int":        int(-42),
			"key_int8":       int8(-8),
			"key_int16":      int16(-16),
			"key_int32":      int32(-32),
			"key_uint":       uint(42),
			"key_uint8":      uint8(8),
			"key_uint16":     uint16(16),
			"key_uint32":     uint32(32),
			"key_uint64":     uint64(64),
			"key_timef":      now.Format(time.RFC850),
			"key_time":       now,
			"key_dur":        dur,
			"key_error":      errors.New("unknown error"),
			"key_complex64":  complex(float32(3), float32(2)),
			"key_complex128": complex(float64(6), float64(4)),
			"key_bytes":      []byte("hello world"),
		}, m)
	}
	gotField := newField().
		Str("key_str", "value_str").
		Int64("key_int64", -64).
		Bool("key_bool", true).
		Float32("key_float32", float32(3.2)).
		Float64("key_float64", float64(6.4)).
		Interface("key_any", "anything").
		Fields(map[string]interface{}{
			"key_map_str":   "value",
			"key_map_int64": int64(-84),
			"key_map_bool":  false,
		}).
		Int("key_int", int(-42)).
		Int8("key_int8", int8(-8)).
		Int16("key_int16", int16(-16)).
		Int32("key_int32", int32(-32)).
		Uint("key_uint", uint(42)).
		Uint8("key_uint8", uint8(8)).
		Uint16("key_uint16", uint16(16)).
		Uint32("key_uint32", uint32(32)).
		Uint64("key_uint64", uint64(64)).
		Timef("key_timef", now, time.RFC850).
		Time("key_time", now).
		Dur("key_dur", dur).
		Error("key_error", errors.New("unknown error")).
		Complex64("key_complex64", complex(float32(3), float32(2))).
		Complex128("key_complex128", complex(float64(6), float64(4))).
		Bytes("key_bytes", []byte("hello world"))

	gotField.Msg("Hello World").(*field).send(log.NoLevel, formatterFunc)
	gotField.Msgf("Hello %s", "World").(*field).send(log.NoLevel, formatterFunc)
}

func TestFieldUpdate(t *testing.T) {
	now := time.Now()
	dur, _ := time.ParseDuration("1d2h3m4s")
	formatterFunc := func(level log.Level, msg string, err error, m map[string]interface{}) {
		assert.Nil(t, err)
		assert.Equal(t, "Hello World", msg)
		assert.Equal(t, map[string]interface{}{
			"key_str":        "value_str",
			"key_int64":      int64(-64),
			"key_bool":       true,
			"key_float32":    float32(3.2),
			"key_float64":    float64(6.4),
			"key_any":        "anything",
			"key_map_str":    "value",
			"key_map_int64":  int64(-84),
			"key_map_bool":   false,
			"key_int":        int(-42),
			"key_int8":       int8(-8),
			"key_int16":      int16(-16),
			"key_int32":      int32(-32),
			"key_uint":       uint(42),
			"key_uint8":      uint8(8),
			"key_uint16":     uint16(16),
			"key_uint32":     uint32(32),
			"key_uint64":     uint64(64),
			"key_timef":      now.Format(time.RFC850),
			"key_time":       now,
			"key_dur":        dur,
			"key_error":      errors.New("unknown error"),
			"key_complex64":  complex(float32(3), float32(2)),
			"key_complex128": complex(float64(6), float64(4)),
			"key_bytes":      []byte("hello world"),
		}, m)
	}
	gotField := newField().
		Str("key_str", "value_str").
		Int64("key_int64", -64).
		Bool("key_bool", true).
		Float32("key_float32", float32(3.2)).
		Float64("key_float64", float64(6.4)).
		Interface("key_any", "anything").
		Fields(map[string]interface{}{
			"key_map_str":   "value",
			"key_map_int64": int64(-84),
			"key_map_bool":  false,
		}).
		Int("key_int", int(-42)).
		Int8("key_int8", int8(-8)).
		Int16("key_int16", int16(-16)).
		Int32("key_int32", int32(-32)).
		Uint("key_uint", uint(42)).
		Uint8("key_uint8", uint8(8)).
		Uint16("key_uint16", uint16(16)).
		Uint32("key_uint32", uint32(32)).
		Uint64("key_uint64", uint64(64)).
		Timef("key_timef", now, time.RFC850).
		Time("key_time", now).
		Dur("key_dur", dur).
		Error("key_error", errors.New("unknown error")).
		Complex64("key_complex64", complex(float32(3), float32(2))).
		Complex128("key_complex128", complex(float64(6), float64(4))).
		Bytes("key_bytes", []byte("hello world"))

	gotField.Msg("Hello World").(*field).send(log.NoLevel, formatterFunc)
	gotField.Msgf("Hello %s", "World").(*field).send(log.NoLevel, formatterFunc)
}
