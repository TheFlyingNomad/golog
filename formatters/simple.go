package formatters

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	housekeeping "github.com/TheFlyingNomad/golog/local-house-keeping"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

type simpleFormatterLogger struct {
	loggerToSendTo gologC.Logger
}

// NewSimpleFormatterLogger -
func NewSimpleFormatterLogger(logger gologC.Logger) gologC.EasyLogger {
	var formatterLogger = &simpleFormatterLogger{
		loggerToSendTo: logger,
	}

	return housekeeping.NewDefaultHelperImplmentation(
		formatterLogger,
	)
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

	var formatting = "%s | %s"
	if len(strings.TrimSpace(logEntry.Tag)) > 0 {
		formatting = formatting + " | %s"
	} else {
		formatting = formatting + " |"
	}

	if logEntry.Level > 0 {
		formatting = formatting + fmt.Sprintf(" %"+strconv.Itoa(logEntry.Level*4)+"v", "")
	}
	formatting = formatting + " %s"

	if len(strings.TrimSpace(logEntry.Tag)) > 0 {
		logEntry.Message = fmt.Sprintf(
			formatting,
			formattedTime,
			logTypeToString[logEntry.LogType],
			logEntry.Tag,
			logEntry.Message,
		)
	} else {
		logEntry.Message = fmt.Sprintf(
			formatting,
			formattedTime,
			logTypeToString[logEntry.LogType],
			logEntry.Message,
		)
	}

	thisRef.loggerToSendTo.Log(logEntry)
}
