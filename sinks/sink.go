// Package sinks defines interface and implementation for sink
package sinks

import (
	"log/slog"
	"os"

	"github.com/himanhsugusain/spinnaker-audit-sink/spinnaker"
)

type Sink interface {
	WriteEvent(spinnaker.Root)
	WriteError(error)
}

type LogSink struct {
	log *slog.Logger
}

func NewLogSink() *LogSink {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	return &LogSink{
		log: log,
	}
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
		},
		"status", r.Payload.Content.Execution.Status,
		"time", map[string]float64{
			"startTime": r.Payload.Content.Execution.StartTime,
			"endTime":   r.Payload.Content.Execution.EndTime,
			"buildTime": r.Payload.Content.Execution.BuildTime,
		},
	)
}

func (l *LogSink) WriteError(err error) {
	l.log.Error(err.Error())
}
