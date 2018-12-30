# Log Quick
```go
import (
	...
	golog "github.com/brightappsllc/golog"
	...
)

func main() {
	...
	golog.Init(golog.NewConsoleLogger())
	...
	golog.Instance().LogInfo("Info line")
	golog.Instance().LogWarning("Warning line")
	golog.Instance().LogError("Error line")
	...
}
```

# Log Smart
```go
import (
	...
	golog  "github.com/brightappsllc/golog"
	gologC "github.com/brightappsllc/golog/contracts"
	gologF "github.com/brightappsllc/golog/formatters"
	gologM "github.com/brightappsllc/golog/mixers"
	gologP "github.com/brightappsllc/golog/persisters"
	...
)

func main() {
	...
	golog.Init(
		gologF.NewSimpleFormatterLogger(
			gologM.NewMultiLogger(
				[]gologC.Logger{
					gologP.NewConsoleLogger(),
					gologP.NewFileLogger("log.log"),
				},
			),
		),
	)
	...
	golog.Instance().LogInfo("Info line")
	golog.Instance().LogWarning("Warning line")
	golog.Instance().LogError("Error line")
	...
}
```

# Log Async and/or Buffered
```go
import (
	...
	golog  "github.com/brightappsllc/golog"
	gologC "github.com/brightappsllc/golog/contracts"
	gologF "github.com/brightappsllc/golog/formatters"
	gologM "github.com/brightappsllc/golog/mixers"
	gologP "github.com/brightappsllc/golog/persisters"
	...
)

func main() {
	...
	golog.Init(
		gologF.NewSimpleFormatterLogger(
			gologM.NewAsyncLogger(
				gologM.NewBufferedLogger(
					gologM.NewMultiLogger(
						[]gologC.Logger{
							gologP.NewConsoleLogger(),
							gologP.NewFileLogger("log.log"),
						},
					),
					gologM.BufferedLoggerConfig{
						MaxLogEntries: 100
					}
				),
			),
		),
	)
	...
	golog.Instance().LogInfo("Info line")
	golog.Instance().LogWarning("Warning line")
	golog.Instance().LogError("Error line")
	...
}
```


# Special Cases Logger -
#### `NewInMemoryGroupedAndSortedLogger()`
- This logger stores the log entries in-memory.
- Groups the log entries by the LOG-TAG
- Sorts the log entrie in each group by the time
- When you call `Flush()` it will save to the persisters a nice-call-stack-alike log lines

```go
import (
	...
	golog  "github.com/brightappsllc/golog"
	gologF "github.com/brightappsllc/golog/formatters"
	gologP "github.com/brightappsllc/golog/persisters"
	...
)

var logTag = "LOG_TAG" // Ex: parentMethod - level 0, childMethod() - level 1
var logLevel = 0 // Useful for call-stack type of execution

func main() {
	...
	var inMemoryGroupedAndSortedLogger = gologP.NewInMemoryGroupedAndSortedLogger(
		gologF.NewSimpleFormatterLogger(
			gologP.NewFileLogger("log.log"),
		),
	)
	golog.Init(inMemoryGroupedAndSortedLogger)
	...
	golog.Instance().LogInfoWithTagAndLevel(logTag, logLevel, "Info line")
	golog.Instance().LogWarningWithTagAndLevel(logTag, logLevel, "Warning line")
	golog.Instance().LogErrorWithTagAndLevel(logTag, logLevel, "Error line")
	...
	// Dump log entries from memory to the configured pipe-line
	(inMemoryGroupedAndSortedLogger.(*gologP.InMemoryGroupedAndSortedLogger)).Flush()
	...
}
```
