package persisters

import (
	"fmt"
	"sort"
	"sync"
	"time"

	gologC "github.com/brightappsllc/golog/contracts"
)

type inMemoryGroupedAndSortedLogger struct {
	logLinesByKey         map[string][]gologC.LogEntry
	logLinesByKeyWithTime map[string]time.Time
	loggerToSendTo        gologC.Logger
	rwMutex               sync.RWMutex
}

// NewInMemoryGroupedAndSortedLogger -
func NewInMemoryGroupedAndSortedLogger(logger gologC.Logger) gologC.Logger {
	return &inMemoryGroupedAndSortedLogger{
		logLinesByKey:         map[string][]gologC.LogEntry{},
		logLinesByKeyWithTime: map[string]time.Time{},
		loggerToSendTo:        logger,
		rwMutex:               sync.RWMutex{},
	}
}

func (thisRef *inMemoryGroupedAndSortedLogger) Log(logEntry gologC.LogEntry) {
	thisRef.rwMutex.Lock()
	defer thisRef.rwMutex.Unlock()

	thisRef.logLinesByKey[logEntry.Tag] = append(thisRef.logLinesByKey[logEntry.Tag], logEntry)

	// Remember the earliest log-line
	var ok bool
	if _, ok = thisRef.logLinesByKeyWithTime[logEntry.Tag]; !ok {
		thisRef.logLinesByKeyWithTime[logEntry.Tag] = logEntry.Time
	} else {
		var storedTime = thisRef.logLinesByKeyWithTime[logEntry.Tag]
		if storedTime.After(logEntry.Time) {
			thisRef.logLinesByKeyWithTime[logEntry.Tag] = logEntry.Time
		}
	}
}

func (thisRef *inMemoryGroupedAndSortedLogger) Flush() {
	thisRef.rwMutex.RLock()
	defer thisRef.rwMutex.RUnlock()

	// Sort by time
	var allTimes = []time.Time{}
	var timeToLogEntryTag = map[int64]string{}
	for key, value := range thisRef.logLinesByKeyWithTime {
		allTimes = append(allTimes, value)
		timeToLogEntryTag[value.UnixNano()] = key
	}
	sort.Slice(
		allTimes,
		func(i, j int) bool {
			return allTimes[i].Before(allTimes[j])
		},
	)

	for index := range allTimes {
		var logEntryTag = timeToLogEntryTag[allTimes[index].UnixNano()]
		var arrayOfLogEntries = thisRef.logLinesByKey[logEntryTag]

		sort.Slice(
			arrayOfLogEntries,
			func(i, j int) bool {
				return arrayOfLogEntries[i].Time.Before(arrayOfLogEntries[j].Time)
			},
		)

		for i := range arrayOfLogEntries {
			arrayOfLogEntries[i].Tag = fmt.Sprintf("%10s", arrayOfLogEntries[i].Tag)
			thisRef.loggerToSendTo.Log(arrayOfLogEntries[i])
		}
	}
}
