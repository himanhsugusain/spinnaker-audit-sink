package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/himanhsugusain/spinnaker-audit-sink/config"
	"github.com/himanhsugusain/spinnaker-audit-sink/sinks"
	"github.com/himanhsugusain/spinnaker-audit-sink/spinnaker"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}))
	app := NewApp(log)
	log.Error(app.Run().Error())
}

type App struct {
	c     config.Config
	log   *slog.Logger
	mux   *http.ServeMux
	sinks []sinks.Sink
}

func NewApp(l *slog.Logger) *App {
	mux := http.DefaultServeMux
	cfg, err := config.GetConfig()
	if err != nil {
		l.Error("error reading config", "err", err.Error())
	} else {
		l.Debug("config", "app.yaml", cfg)
	}
	return &App{
		c:     cfg,
		mux:   mux,
		log:   l,
		sinks: sinks.GetSinks(cfg.Sinks),
	}

}

func (a *App) Run() error {
	a.mux.HandleFunc("/spinnakerAuditLogs/", a.spinnakerAuditLogs())
	a.mux.HandleFunc("/", a.defaultHandlerFunc())
	a.log.Info("starting server", "port", a.c.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", a.c.Port), a.mux)
}

func (a *App) logRequest(r *http.Request, start time.Time) {
	a.log.Debug("request", "method", r.Method, "path", r.URL.Path, "latency", time.Since(start), "remote", r.RemoteAddr, "user-agent", r.UserAgent())
}

func (a *App) defaultHandlerFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer a.logRequest(r, start)
		http.NotFoundHandler().ServeHTTP(w, r)
	}
}

func (a *App) spinnakerAuditLogs() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer a.logRequest(r, start)
		if err := verifyRequestor(a.c, r); err != nil {
			a.log.Error("failed to authenticate", "error", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		var event spinnaker.Root
		data, err := io.ReadAll(r.Body)
		if err != nil {
			a.log.Error("failed to read request body", "error", err)
			http.Error(w, "failed to read request body", http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal(data, &event); err != nil {
			a.log.Error("failed to parse event", "error", err)
			http.Error(w, "failed to parse event", http.StatusBadRequest)
			return
		} else {
			if a.FilterEvent(&event) {
				for _, s := range a.sinks {
					s.WriteEvent(event)
				}
			} else {
				a.log.Debug("filtered out", "event", event.Payload.Details.Type)
			}
		}
	}
}

func (a *App) FilterEvent(event *spinnaker.Root) bool {
	keep := false
	a.log.Debug("detailType filter", event.Payload.Details.Type, a.c.Filter.DetailsType)
	for _, detailType := range a.c.Filter.DetailsType {
		if strings.HasPrefix(event.Payload.Details.Type, detailType) {
			keep = true
			break
		}
	}
	return keep
}
func verifyRequestor(cfg config.Config, r *http.Request) error {
	auth := r.Header.Get("authorization")

	basicAuth, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(auth, "Basic ", ""))
	if err != nil {
		return err
	}
	parts := strings.Split(string(basicAuth), ":")
	if parts[0] != cfg.UserName || parts[1] != cfg.PassWord {
		return fmt.Errorf("username, password mismatch")
	}
	return nil
}
