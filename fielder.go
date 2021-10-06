package goplogadapter

import (
	"fmt"
	"time"

	"github.com/axpira/gop/log"
)

type field struct {
	m   map[string]interface{}
	msg string
	err error
}

func newField() *field {
	field := &field{m: make(map[string]interface{}, 500)}
	return field
}

func (f *field) Str(key string, value string) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Int64(key string, value int64) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Bool(key string, value bool) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Float32(key string, value float32) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Float64(key string, value float64) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Msg(msg string) log.FieldBuilder {
	f.msg = msg
	return f
}

func (f *field) Msgf(format string, args ...interface{}) log.FieldBuilder {
	return f.Msg(fmt.Sprintf(format, args...))
}

func (f *field) Any(key string, value interface{}) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Marshal(key string, value interface{}) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Fields(m map[string]interface{}) log.FieldBuilder {
	for k, v := range m {
		f.m[k] = v
	}
	return f
}

func (f *field) Dict(key string, value log.FieldBuilder) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Int(key string, value int) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Int8(key string, value int8) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Int16(key string, value int16) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Int32(key string, value int32) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Uint(key string, value uint) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Uint8(key string, value uint8) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Uint16(key string, value uint16) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Uint32(key string, value uint32) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Uint64(key string, value uint64) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Timef(key string, value time.Time, format string) log.FieldBuilder {
	if !value.IsZero() {
		f.m[key] = value.Format(format)
	}
	return f
}

func (f *field) Time(key string, value time.Time) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Dur(key string, value time.Duration) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Stringer(key string, value fmt.Stringer) log.FieldBuilder {
	if value != nil {
		f.m[key] = value.String()
	}
	return f
}

func (f *field) Error(key string, value error) log.FieldBuilder {
	if value != nil {
		f.m[key] = value
	}
	return f
}

func (f *field) Err(value error) log.FieldBuilder {
	f.err = value
	return f
}

func (f *field) Complex64(key string, value complex64) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Complex128(key string, value complex128) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Bytes(key string, value []byte) log.FieldBuilder {
	f.m[key] = value
	return f
}

func (f *field) Update(l log.Logger) log.Logger {
	lNew := l.(logger).clone()
	lNew.fielder = lNew.fielder.Fields(f.m).(*field)
	return lNew
}

func (f *field) clone() *field {
	newField := newField()
	for key, value := range f.m {
		newField.m[key] = value
	}
	newField.msg = f.msg
	newField.err = f.err
	return newField
}

func (f *field) send(level log.Level, formatter FormatterFunc) {
	formatter(level, f.msg, f.err, f.m)
}
