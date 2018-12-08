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

	LogPanicWithFields(fields Fields)
	LogFatalWithFields(fields Fields)
	LogErrorWithFields(fields Fields)
	LogWarningWithFields(fields Fields)
	LogInfoWithFields(fields Fields)
	LogDebugWithFields(fields Fields)
}
