package housekeeping

import (
	"time"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

func (thisRef defaultHelperImplmentation) LogPanicWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypePanic,
		Tag:     tag,
		Level:   level,
		Message: message,
	})
}
func (thisRef defaultHelperImplmentation) LogFatalWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypeFatal,
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
func (thisRef defaultHelperImplmentation) LogWarningWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypeWarning,
		Tag:     tag,
		Level:   level,
		Message: message,
	})
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
func (thisRef defaultHelperImplmentation) LogDebugWithTagAndLevel(tag string, level int, message string) {
	thisRef.Log(gologC.LogEntry{
		Time:    time.Now(),
		LogType: gologC.TypeDebug,
		Tag:     tag,
		Level:   level,
		Message: message,
	})
}
