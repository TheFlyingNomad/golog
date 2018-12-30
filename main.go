package golog

import (
	housekeeping "github.com/brightappsllc/golog/local-house-keeping"

	gologC "github.com/brightappsllc/golog/contracts"
	gologP "github.com/brightappsllc/golog/persisters"
)

var instance gologC.EasyLogger

// Init -
func Init(logger gologC.EasyLogger) {
	instance = logger
}

// NewConsoleLogger -
func NewConsoleLogger() gologC.EasyLogger {
	return housekeeping.NewDefaultHelperImplmentation(
		gologP.NewConsoleLogger(),
	)
}

// Instance -
func Instance() gologC.EasyLogger {
	return instance
}
