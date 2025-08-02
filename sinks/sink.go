// Package sinks defines interface and implementation for sink
package sinks

import (
	"log/slog"

	"github.com/himanhsugusain/spinnaker-audit-sink/spinnaker"
)

type Sink interface {
	WriteEvent(spinnaker.Root)
	WriteError(error)
}

type LogSink struct {
	log *slog.Logger
}

func NewLogSink(log *slog.Logger) *LogSink {
	return &LogSink{
		log: log,
	}
}

func (l *LogSink) WriteEvent(r spinnaker.Root) {
	l.log.Info("write", "event", r)
}

func (l *LogSink) WriteError(err error) {
	l.log.Error(err.Error())
}
