/*
Package log provides a abstraction on loggers

supports both global usage and dependency injection usage
*/
package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var defaultLogger *zerolog.Logger

type logger struct {
	l zerolog.Logger
}

// Logger implements logger methods.
type Logger interface {
	Println(msgs ...interface{})
	Debug(msgs ...interface{})
	Info(msgs ...interface{})
	Error(err error, msgs ...interface{})
}

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	out := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	zlogger := zerolog.New(out).With().Timestamp().Logger()
	defaultLogger = &zlogger
}

// SetGlobalLevel sets the log levels to be printed
//
// equal or above levels will be printed
func SetGlobalLevel(l level) {
	switch l {
	case DebugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case InfoLevel:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WarnLevel:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ErrorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case FatalLevel:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case PanicLevel:
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case NoLevel:
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case Disabled:
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
}

// Println prints a log without level
func Println(msgs ...interface{}) {
	defaultLogger.Log().Caller(1).Msg(fmt.Sprintf("%v", msgs))
}

// Debug prints debug log
func Debug(msgs ...interface{}) {
	defaultLogger.Debug().Caller(1).Msg(fmt.Sprintf("%v", msgs))
}

// Info prints information log
func Info(msgs ...interface{}) {
	defaultLogger.Info().Caller(1).Msg(fmt.Sprintf("%v", msgs))
}

// Error prints error log
func Error(err error, msgs ...interface{}) {
	// cause := GetCauseCallerFromErr(err)
	cause := ""
	stackStr := ""
	stack := StackTrace(errors.Cause(err))

	if pc, _, _, ok := runtime.Caller(1); ok {
		// we get the caller of the log, which is the highest level
		caller := runtime.FuncForPC(pc)
		filepath, line := caller.FileLine(pc)
		filename := path.Base(filepath)

		for i, frame := range stack {
			currentFrameCaller := strings.Split(fmt.Sprintf("%s", frame), ":")[0]
			stackStr += fmt.Sprintf("%s ", currentFrameCaller)

			if i == 0 {
				cause = fmt.Sprintf("%s, %s:%d", runtime.FuncForPC(uintptr(frame)).Name(), filename, line)
			}

			if currentFrameCaller == filename {
				break
			}
		}

		stackStr = strings.TrimSpace(stackStr)
	}

	if cause == "" {
		defaultLogger.Error().Caller(1).Err(err).Msg(fmt.Sprintf("%v", msgs))

	} else if stackStr == "" {
		defaultLogger.Error().Caller(1).Err(err).Str("cause", cause).Msg(fmt.Sprintf("%v", msgs))
	} else {
		defaultLogger.Error().Caller(1).Err(err).Str("stack", stackStr).Str("cause", cause).Msg(fmt.Sprintf("%v", msgs))
	}

}

// With adds a key,value pair to the global logger logs
func With(key, val string) {
	l := defaultLogger.With().Str(key, val).Logger()
	defaultLogger = &l
}

// NewWith creates a child logger with a key,value context
func NewWith(key, val string) Logger {
	return logger{
		l: defaultLogger.With().Str(key, val).Logger(),
	}
}

// Println prints a log without level
func (l logger) Println(msgs ...interface{}) {
	l.l.Log().Caller(1).Msg(fmt.Sprintf("%v", msgs))
}

// Debug prints debug log
func (l logger) Debug(msgs ...interface{}) {
	l.l.Debug().Caller(1).Msg(fmt.Sprintf("%v", msgs))
}

// Info prints information log
func (l logger) Info(msgs ...interface{}) {
	l.l.Info().Caller(1).Msg(fmt.Sprintf("%v", msgs))
}

func (l logger) Error(err error, msgs ...interface{}) {
	l.l.Error().Caller(1).Err(err).Msg(fmt.Sprintf("%v", msgs))
}
