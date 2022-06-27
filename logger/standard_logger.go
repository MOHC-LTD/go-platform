package logger

import (
	"fmt"
	"log"
	"os"
)

//StandardLogger holds properties that help logs go to system out
type StandardLogger struct {
	log           Level
	loggers       map[Level]*log.Logger
	callStackSkip int
}

//NewStandardLogger create set of standard out loggers with different prefixes
func NewStandardLogger(logLevel Level) Logger {
	logMap := map[Level]*log.Logger{}

	logMap[DEBUG] = log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile)
	logMap[INFO] = log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
	logMap[WARN] = log.New(os.Stdout, "WARN ", log.Ldate|log.Ltime|log.Lshortfile)
	logMap[ERROR] = log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	return StandardLogger{logLevel, logMap, 3}
}

//Debug prints a message to standard out if the log level is debug or below
func (l StandardLogger) Debug(message string) {
	if l.log <= DEBUG {
		l.print(l.loggers[DEBUG], message)
	}
}

//Debugf prints a formatted message to standard out if the log level is debug or below
func (l StandardLogger) Debugf(message string, args ...interface{}) {
	if l.log <= DEBUG {
		l.printf(l.loggers[DEBUG], message, args...)
	}
}

//Info prints a message to standard out if the log level is info or below
func (l StandardLogger) Info(message string) {
	if l.log <= INFO {
		l.print(l.loggers[INFO], message)
	}
}

//Infof prints a formatted message to standard out if the log level is info or below
func (l StandardLogger) Infof(message string, args ...interface{}) {
	if l.log <= INFO {
		l.printf(l.loggers[INFO], message, args...)
	}
}

//Warn prints a message to standard out if the log level is warn or below
func (l StandardLogger) Warn(message string) {
	if l.log <= WARN {
		l.print(l.loggers[WARN], message)
	}
}

//Warnf prints a formatted message to standard out if the log level is warn or below
func (l StandardLogger) Warnf(message string, args ...interface{}) {
	if l.log <= WARN {
		l.printf(l.loggers[WARN], message, args...)
	}
}

//Error prints a message to standard out if the log level is error or below
func (l StandardLogger) Error(message string) {
	if l.log <= ERROR {
		l.print(l.loggers[ERROR], message)
	}
}

//Errorf prints a formatted message to standard out if the log level is error or below
func (l StandardLogger) Errorf(message string, args ...interface{}) {
	if l.log <= ERROR {
		l.printf(l.loggers[ERROR], message, args...)
	}
}

//FatalF prints a formatted message before calling os.Exit(1)
func (l StandardLogger) Fatalf(message string, args ...interface{}) {
	l.printf(l.loggers[ERROR], message, args...)
	os.Exit(1)
}

func (l StandardLogger) print(logger *log.Logger, message string) {
	_ = logger.Output(l.callStackSkip, message)
}

func (l StandardLogger) printf(logger *log.Logger, format string, args ...interface{}) {
	_ = logger.Output(l.callStackSkip, fmt.Sprintf(format, args...))
}
