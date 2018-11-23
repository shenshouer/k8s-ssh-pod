package log

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(`%{color}%{time:2006-01-02T15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func init() {
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backend1Leveled, backend2Formatter)
}

// IsEnabledFor returns true if the logger is enabled for the given level.
func IsEnabledFor(level logging.Level) bool {
	return log.IsEnabledFor(level)
}

// Fatal is equivalent to l.Critical(fmt.Sprint()) followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf is equivalent to l.Critical followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panic is equivalent to l.Critical(fmt.Sprint()) followed by a call to panic().
func Panic(args ...interface{}) {
	log.Panic(args...)
}

// Panicf is equivalent to l.Critical followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// Critical logs a message using CRITICAL as log level.
func Critical(args ...interface{}) {
	log.Critical(args...)
}

// Criticalf logs a message using CRITICAL as log level.
func Criticalf(format string, args ...interface{}) {
	log.Criticalf(format, args...)
}

// Error logs a message using ERROR as log level.
func Error(args ...interface{}) {
	log.Error(args...)
}

// Errorf logs a message using ERROR as log level.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Warning logs a message using WARNING as log level.
func Warning(args ...interface{}) {
	log.Warning(args...)
}

// Warningf logs a message using WARNING as log level.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}

// Notice logs a message using NOTICE as log level.
func Notice(args ...interface{}) {
	log.Notice(args...)
}

// Noticef logs a message using NOTICE as log level.
func Noticef(format string, args ...interface{}) {
	log.Noticef(format, args...)
}

// Info logs a message using INFO as log level.
func Info(args ...interface{}) {
	log.Info(args...)
}

// Infof logs a message using INFO as log level.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Debug logs a message using DEBUG as log level.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Debugf logs a message using DEBUG as log level.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}
