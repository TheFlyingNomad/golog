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

func (thisRef fileLogger) Log(logEntry gologC.LogEntry) {
	if thisRef.errorOccured {
		return
	}

	if logEntry.LogType > thisRef.logUntil {
		return
	}

	thisRef.file.WriteString(logEntry.Message + "\n")
}
