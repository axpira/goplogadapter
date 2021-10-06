package json

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/axpira/gop/log"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	out := new(strings.Builder)
	Output = out

	tests := []struct {
		name  string
		level log.Level
	}{
		{"trace", log.TraceLevel},
		{"debug", log.DebugLevel},
		{"info", log.InfoLevel},
		{"warn", log.WarnLevel},
		{"error", log.ErrorLevel},
	}
	for _, tc := range tests {
		out.Reset()
		JsonFormatter(tc.level, "Hello World", errors.New("unknown error"), map[string]interface{}{})
		assert.JSONEq(t, fmt.Sprintf(`{"err":"unknown error", "level":"%s", "msg":"Hello World"}`, tc.name), out.String())
	}
}
