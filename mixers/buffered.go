package mixers

import (
	"sync"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

// BufferedLoggerConfig -
type BufferedLoggerConfig struct {
	MaxLogEntries int
}

type bufferedLogger struct {
	loggerToSendTo gologC.Logger
	logEntries     []gologC.LogEntry
	rwMutex        sync.RWMutex
	config         BufferedLoggerConfig
}

// NewBufferedLogger -
func NewBufferedLogger(logger gologC.Logger, config BufferedLoggerConfig) gologC.Logger {
	return &bufferedLogger{
		loggerToSendTo: logger,
		logEntries:     []gologC.LogEntry{},
		rwMutex:        sync.RWMutex{},
		config:         config,
	}
}

func (thisRef *bufferedLogger) Log(logEntry gologC.LogEntry) {
	thisRef.rwMutex.Lock()
	defer thisRef.rwMutex.Unlock()

	thisRef.logEntries = append(thisRef.logEntries, logEntry)

	if len(thisRef.logEntries) > thisRef.config.MaxLogEntries {
		for _, logEntry := range thisRef.logEntries {
			thisRef.loggerToSendTo.Log(logEntry)
		}

		thisRef.logEntries = []gologC.LogEntry{}
	}
}
