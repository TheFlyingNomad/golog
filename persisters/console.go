package persisters

import (
	"fmt"

	gologC "github.com/TheFlyingNomad/golog/contracts"
	coloredLogs "github.com/logrusorgru/aurora"
)

type consoleLogger struct{}

// NewConsoleLogger -
func NewConsoleLogger() gologC.Logger {
	return &consoleLogger{}
}

func (thisRef consoleLogger) Log(logEntry gologC.LogEntry) {
	if logEntry.LogType == gologC.TypeWarning {
		fmt.Println(coloredLogs.Magenta(logEntry.Message))
	} else if logEntry.LogType == gologC.TypeError {
		fmt.Println(coloredLogs.Red(logEntry.Message))
	} else {
		fmt.Println(logEntry.Message)
	}
}
