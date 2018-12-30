package persisters

import (
	"fmt"

	gologC "github.com/codemodify/golog/contracts"
	coloredLogs "github.com/logrusorgru/aurora"
)

type consoleLogger struct{}

// NewConsoleLogger -
func NewConsoleLogger() gologC.Logger {
	return &consoleLogger{}
}

func (thisRef consoleLogger) Log(logEntry gologC.LogEntry) {
	// TypePanic = iota
	// TypeFatal
	// TypeError
	// TypeWarning
	// TypeInfo
	// TypeDebug

	if logEntry.LogType < gologC.TypeWarning {
		fmt.Println(coloredLogs.Red(logEntry.Message))
	} else if logEntry.LogType == gologC.TypeWarning {
		fmt.Println(coloredLogs.Magenta(logEntry.Message))
	} else if logEntry.LogType == gologC.TypeInfo {
		fmt.Println(logEntry.Message)
	} else if logEntry.LogType == gologC.TypeDebug {
		fmt.Println(coloredLogs.Green(logEntry.Message))
	}
}
