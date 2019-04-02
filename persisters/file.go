package persisters

import (
	"fmt"
	"os"

	gologC "github.com/brightappsllc/golog/contracts"
)

type fileLogger struct {
	file         *os.File
	errorOccured bool
	logUntil     int
}

// NewFileLogger -
func NewFileLogger(logUntil int, fileName string) gologC.Logger {
	f, err := os.Create(fileName)
	if err != nil {
		// return nil, err
		fmt.Println(err)
	}

	return &fileLogger{
		file:         f,
		errorOccured: (err != nil),
		logUntil:     logUntil,
	}
}

// NewFileLoggerDefaultName -
func NewFileLoggerDefaultName(logUntil int) gologC.Logger {
	return NewFileLogger(logUntil, fmt.Sprintf("%s.log", os.Args[0]))
}

func (thisRef fileLogger) Log(logEntry gologC.LogEntry) {
	if thisRef.errorOccured {
		return
	}

	if logEntry.LogType == gologC.TypeDisable {
		return
	}

	if logEntry.LogType > thisRef.logUntil &&
		logEntry.LogType != gologC.TypeTrace {
		return
	}

	thisRef.file.WriteString(logEntry.Message + "\n")
}
