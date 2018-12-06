package formatters

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

type simpleFormatterLogger struct {
	loggerToSendTo gologC.Logger
}

// NewSimpleFormatterLogger -
func NewSimpleFormatterLogger(logger gologC.Logger) gologC.Logger {
	return &simpleFormatterLogger{
		loggerToSendTo: logger,
	}
}

func (thisRef simpleFormatterLogger) Log(logEntry gologC.LogEntry) {
	var formattedTime = logEntry.Time.UTC().Format(time.RFC3339Nano)
	var formattedTimeLen = len(formattedTime)
	if formattedTimeLen < 30 {
		var spacesCount = 30 - formattedTimeLen

		var newV = fmt.Sprintf("%"+strconv.Itoa(spacesCount+1)+"v", "Z")
		newV = strings.Replace(newV, " ", "0", spacesCount)

		formattedTime = strings.Replace(
			formattedTime,
			"Z",
			newV,
			1,
		)
	}

	logEntry.Message = fmt.Sprintf(
		"%s %s %s %"+strconv.Itoa(logEntry.Level*4)+"v %s",
		formattedTime,
		logEntry.LogType,
		logEntry.Tag,
		"",
		logEntry.Message,
	)

	thisRef.loggerToSendTo.Log(logEntry)
}
