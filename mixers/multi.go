package mixers

import (
	housekeeping "github.com/codemodify/golog/local-house-keeping"

	gologC "github.com/codemodify/golog/contracts"
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
