package housekeeping

import (
	"time"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

type defaultHelperImplmentation struct {
	loggerToSendTo gologC.Logger
}

// NewDefaultHelperImplmentation -
func NewDefaultHelperImplmentation(logger gologC.Logger) gologC.EasyLogger {
	return &defaultHelperImplmentation{
		loggerToSendTo: logger,
	}
}

func (thisRef defaultHelperImplmentation) Log(logEntry gologC.LogEntry) {
	thisRef.loggerToSendTo.Log(logEntry)
}

func (thisRef defaultHelperImplmentation) LogInfoWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypeInfo,
		Tag:     tag,
		Level:   level,
		Message: message,
	})
}

func (thisRef defaultHelperImplmentation) LogWarningWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypeWarning,
		Tag:     tag,
		Level:   level,
		Message: message,
	})
}

func (thisRef defaultHelperImplmentation) LogErrorWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypeError,
		Tag:     tag,
		Level:   level,
		Message: message,
	})
}

func (thisRef defaultHelperImplmentation) LogInfo(message string) {
	thisRef.LogInfoWithTagAndLevel("", 0, message)
}

func (thisRef defaultHelperImplmentation) LogWarning(message string) {
	thisRef.LogWarningWithTagAndLevel("", 0, message)
}

func (thisRef defaultHelperImplmentation) LogError(message string) {
	thisRef.LogErrorWithTagAndLevel("", 0, message)
}
