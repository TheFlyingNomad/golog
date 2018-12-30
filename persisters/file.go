package persisters

import (
	"fmt"
	"os"

	gologC "github.com/codemodify/golog/contracts"
)

type fileLogger struct {
	file         *os.File
	errorOccured bool
}

// NewFileLogger -
func NewFileLogger(fileName string) gologC.Logger {
	f, err := os.Create(fileName)
	if err != nil {
		// return nil, err
		fmt.Println(err)
	}

	return &fileLogger{
		file:         f,
		errorOccured: (err != nil),
	}
}

func (thisRef fileLogger) Log(logEntry gologC.LogEntry) {
	if thisRef.errorOccured {
		return
	}

	thisRef.file.WriteString(logEntry.Message + "\n")
}
