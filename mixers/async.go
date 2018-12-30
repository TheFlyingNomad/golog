package mixers

import (
	gologC "github.com/brightappsllc/golog/contracts"
)

type asyncLogger struct {
	loggerToSendTo gologC.Logger
}

// NewAsyncLogger -
func NewAsyncLogger(logger gologC.Logger) gologC.Logger {
	return &asyncLogger{
		loggerToSendTo: logger,
	}
}

func (thisRef asyncLogger) Log(logEntry gologC.LogEntry) {
	go func(theLogEntryCopy gologC.LogEntry) {
		thisRef.loggerToSendTo.Log(theLogEntryCopy)
	}(logEntry)
}
