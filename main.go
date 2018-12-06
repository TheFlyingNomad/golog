package main

import (
	housekeeping "github.com/TheFlyingNomad/golog/local-house-keeping"

	gologC "github.com/TheFlyingNomad/golog/contracts"
	gologP "github.com/TheFlyingNomad/golog/persisters"
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

func main() {}
