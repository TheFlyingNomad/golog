package contracts

import "time"

// TypeError -
const (
	TypeInfo    = 0
	TypeWarning = 1
	TypeError   = 2
)

var logTypeToString = map[int]string{
	TypeInfo:    "INFO",
	TypeWarning: "OOPS",
	TypeError:   "OSHT",
}

// LogEntry -
type LogEntry struct {
	Time    time.Time // time.Now()
	LogType int       // TypeInfo, TypeWarning, TypeError
	Tag     string    // "This-Is-A-Tag"
	Level   int       // Ex: parentMethod - level 1, childMethod() - level 2
	Message string    //
}

// Logger -
type Logger interface {
	Log(logEntry LogEntry)
}

// EasyLogger -
type EasyLogger interface {
	Logger

	LogInfoWithTagAndLevel(tag string, level int, message string)
	LogWarningWithTagAndLevel(tag string, level int, message string)
	LogErrorWithTagAndLevel(tag string, level int, message string)

	LogInfo(message string)
	LogWarning(message string)
	LogError(message string)
}
