package logger

import (
	"sync"
)

// Severity is a level of severity associated with a log, ascending values are more severe
type Severity int

const (
	// SeverityDebug logs all operations and useful debug information
	SeverityDebug Severity = iota
	// SeverityInfo logs all operations
	SeverityInfo
	// SeverityError logs severe errors occuring during operations
	SeverityError
)

var severity = SeverityError
var mu = sync.RWMutex{}

// SetSeverity sets the global log severity, logs below this severity will not be registered
func SetSeverity(newSeverity Severity) {
	mu.Lock()
	defer mu.Unlock()

	severity = newSeverity
}
