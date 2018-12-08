package contracts

import "time"

// TypePanic -
const (
	TypePanic = iota
	TypeFatal
	TypeError
	TypeWarning
	TypeInfo
	TypeDebug
)

// Fields -
type Fields map[string]interface{}

var logTypeToString = map[int]string{
	TypePanic:   "F__K",
	TypeFatal:   "OSHT",
	TypeError:   "ERRR",
	TypeWarning: "OOPS",
	TypeInfo:    "INFO",
	TypeDebug:   "DEBU",
}

// LogEntry -
type LogEntry struct {
	Time    time.Time // time.Now()
	LogType int       // TypePanic ... TypeDebug
	Tag     string    // "This-Is-A-Tag"
	Level   int       // Ex: parentMethod - level 0, childMethod() - level 1
	Message string    //
}

// Logger -
type Logger interface {
	Log(logEntry LogEntry)
}

// EasyLogger -
type EasyLogger interface {
	Logger

	KeepOnlyLogs(logTypes []int)

	LogPanicWithTagAndLevel(tag string, level int, message string)
	LogFatalWithTagAndLevel(tag string, level int, message string)
	LogErrorWithTagAndLevel(tag string, level int, message string)
	LogWarningWithTagAndLevel(tag string, level int, message string)
	LogInfoWithTagAndLevel(tag string, level int, message string)
	LogDebugWithTagAndLevel(tag string, level int, message string)

	LogPanic(message string)
	LogFatal(message string)
	LogError(message string)
	LogWarning(message string)
	LogInfo(message string)
	LogDebug(message string)

	LogPanicWithFields(message string, fields Fields)
	LogFatalWithFields(message string, fields Fields)
	LogErrorWithFields(message string, fields Fields)
	LogWarningWithFields(message string, fields Fields)
	LogInfoWithFields(message string, fields Fields)
	LogDebugWithFields(message string, fields Fields)
}
