package logrus

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/axpira/gop/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockLogger struct {
	mock.Mock
}

func TestLogrus(t *testing.T) {
	out := new(strings.Builder)
	logrus.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp: true,
	})
	logrus.SetOutput(out)
	logrus.SetLevel(logrus.TraceLevel)

	tests := []struct {
		name  string
		level log.Level
	}{
		{"trace", log.TraceLevel},
		{"debug", log.DebugLevel},
		{"info", log.InfoLevel},
		{"warning", log.WarnLevel},
		{"error", log.ErrorLevel},
	}
	for _, tc := range tests {
		out.Reset()
		LogrusFormatter(tc.level, "Hello World", errors.New("unknown error"), map[string]interface{}{
			"key_str": "value",
		})
		assert.JSONEq(t, fmt.Sprintf(`{"error":"unknown error", "key_str": "value", "level":"%s", "msg":"Hello World"}`, tc.name), out.String())
	}
}
