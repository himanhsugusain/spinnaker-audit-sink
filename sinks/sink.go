// Package sinks defines interface and implementation for sink
package sinks

import (
	"log/slog"
	"os"

	"github.com/himanhsugusain/spinnaker-audit-sink/spinnaker"
)

type Sink interface {
	Key() string
	WriteEvent(spinnaker.Root)
	WriteError(error)
}
type LogSink struct {
	log *slog.Logger
}

func ReplaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Key == "time" {
		return slog.Attr{} // Return an empty Attr to omit it
	}
	return a
}

func GetSinks(sinkList []string) []Sink {
	s := make([]Sink, 0)
	for _, n := range sinkList {
		if n == "logger" {
			s = append(s, newLogSink())
		}
	}
	return s
}

func newLogSink() *LogSink {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		ReplaceAttr: ReplaceAttr,
	}))
	return &LogSink{
		log: log,
	}
}

func (l *LogSink) Key() string {
	return "logger"
}
func (l *LogSink) WriteEvent(r spinnaker.Root) {
	l.log.Info(
		r.Payload.Details.Type,
		"trigger", map[string]string{
			"user": r.Payload.Content.Execution.Trigger.User,
			"type": r.Payload.Content.Execution.Trigger.Type,
		},
		"pipeline", map[string]string{
			"name":             r.Payload.Content.Execution.Name,
			"pipelineConfigId": r.Payload.Content.Execution.PipelineConfigID,
			"application":      r.Payload.Details.Application,
		},
		"status", r.Payload.Content.Execution.Status,
		"execution", map[string]any{
			"id":        r.Payload.Content.Execution.ID,
			"startTime": r.Payload.Content.Execution.StartTime,
			"endTime":   r.Payload.Content.Execution.EndTime,
			"buildTime": r.Payload.Content.Execution.BuildTime,
		},
	)
}

func (l *LogSink) WriteError(err error) {
	l.log.Error(err.Error())
}
