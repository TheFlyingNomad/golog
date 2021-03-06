package formatters

import (
	gologC "github.com/brightappsllc/golog/contracts"
)

// LogTypeToString -
var logTypeToString = map[int]string{
	gologC.TypePanic:   "FWORD",
	gologC.TypeFatal:   "OSHIT",
	gologC.TypeError:   "ERROR",
	gologC.TypeWarning: "OOOPS",
	gologC.TypeInfo:    "INFOO",
	gologC.TypeDebug:   "DEBUG",
	gologC.TypeTrace:   "TRACE",
}
