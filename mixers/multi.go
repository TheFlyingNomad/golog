package mixers

import (
	housekeeping "github.com/TheFlyingNomad/golog/local-house-keeping"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

type multiLogger struct {
	loggers []gologC.Logger
}

// NewMultiLogger -
func NewMultiLogger(loggers []gologC.Logger) gologC.Logger {
	var thisRef = &multiLogger{
		loggers: loggers,
	}

	return housekeeping.NewDefaultHelperImplmentation(thisRef)
}

func (thisRef multiLogger) Log(logEntry gologC.LogEntry) {
	for _, logger := range thisRef.loggers {
		logger.Log(logEntry)
	}
}
