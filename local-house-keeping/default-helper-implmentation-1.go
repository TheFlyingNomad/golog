package housekeeping

import (
	gologC "github.com/TheFlyingNomad/golog/contracts"
)

type defaultHelperImplmentation struct {
	loggerToSendTo gologC.Logger
	logTypes       []int
}

// NewDefaultHelperImplmentation -
func NewDefaultHelperImplmentation(logger gologC.Logger) gologC.EasyLogger {
	return &defaultHelperImplmentation{
		loggerToSendTo: logger,
		logTypes: []int{
			gologC.TypePanic,
			gologC.TypeFatal,
			gologC.TypeError,
			gologC.TypeWarning,
			gologC.TypeInfo,
			gologC.TypeDebug,
		},
	}
}

func (thisRef *defaultHelperImplmentation) KeepOnlyLogs(logTypes []int) {
	thisRef.logTypes = logTypes
}

func (thisRef defaultHelperImplmentation) Log(logEntry gologC.LogEntry) {
	thisRef.loggerToSendTo.Log(logEntry)
}
