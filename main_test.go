package main

import (
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/himanhsugusain/spinnaker-audit-sink/config"
	"github.com/himanhsugusain/spinnaker-audit-sink/sinks"
)

func TestEchoHandler(t *testing.T) {
	data, err := os.ReadFile("test-event.json")
	if err != nil {
		t.Fatalf("Failed to read test-event.json: %v", err)
	}

	req := httptest.NewRequest("POST", "/events", strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("foo:bar"))))

	rr := httptest.NewRecorder()
	ll := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	app := &App{
		c: config.Config{
			Port:     "8080",
			UserName: "foo",
			PassWord: "bar",
		},
		log: ll,
		sinks: []sinks.Sink{
			sinks.NewLogSink(ll),
		},
	}
	handler := http.HandlerFunc(app.spinnakerAuditLogs())
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", rr.Code)
	}
}
