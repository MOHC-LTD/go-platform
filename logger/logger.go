package logger

//Logger logs messages with a level
type Logger interface {
	// Debug log a message at a debug level
	Debug(message string)
	// Debugf log a message at a debug level with formatting
	Debugf(format string, args ...interface{})
	// Info log a message at a info level
	Info(message string)
	// Infof log a message at a info level with formatting
	Infof(format string, args ...interface{})
	// Warn log a message at a warn level
	Warn(message string)
	// Warnf log a message at a warn level with formatting
	Warnf(format string, args ...interface{})
	// Error log a message at a error level
	Error(message string)
	// Errorf log a message at a error level with formatting
	Errorf(format string, args ...interface{})
	// Fatalf log a message and exits
	Fatalf(format string, args ...interface{})
}

// Level represents the log level
type Level int64

// Constants showing level types incremented by number
const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	OFF
)

// BuildLogger constructs a logger based off of the level required
func BuildLogger(logLevel string) Logger {
	var log Logger

	switch logLevel {
	case "DEBUG":
		log = NewStandardLogger(DEBUG)
		log.Debug("Application Started Log Level Debug")
	case "INFO":
		log = NewStandardLogger(INFO)
		log.Info("Application Started Log Level Info")
	case "WARN":
		log = NewStandardLogger(WARN)
		log.Warn("Application Started Log Level Warn")
	case "ERROR":
		log = NewStandardLogger(ERROR)
		log.Error("Application Started Log Level Error")
	case "OFF":
		log = NewStandardLogger(OFF)
		log.Info("Application Started Logging Is off")
	default:
		log = NewStandardLogger(INFO)
		log.Info("Application Started Log Level Info")
	}

	return log
}
