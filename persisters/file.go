package persisters

import (
	"os"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

type fileLogger struct {
	file *os.File
}

// NewFileLogger -
func NewFileLogger(fileName string) (gologC.Logger, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return &fileLogger{
		file: f,
	}, nil
}

func (thisRef fileLogger) Log(logEntry gologC.LogEntry) {
	thisRef.file.WriteString(logEntry.Message + "\n")
}
