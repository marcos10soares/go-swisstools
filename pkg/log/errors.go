package log

import (
	"fmt"
	"path"
	"runtime"

	"github.com/pkg/errors"
)

// errors from pkg/errors can have wrapped messages
type errCauser interface {
	Cause() error
}

// errors from pkg/errors have StackTrace
type errStack interface {
	StackTrace() errors.StackTrace
}

// StackTrace returns the stack trace with given error by the formatter.
// If the error is not traceable, nil is returned.
func StackTrace(err error) errors.StackTrace {
	if tracer, ok := err.(errStack); ok {
		stack := tracer.StackTrace()
		return stack
	}
	return nil
}

// Cause returns the original error of a traceable error.
// If the error is not traceable, return itself.
func Cause(err error) error {
	if tracer, ok := err.(errCauser); ok {
		return tracer.Cause()
	}
	return err
}

func GetCauseCallerFromErr(err error) string {
	stack := StackTrace(errors.Cause(err))

	// we want to get the file path also
	if stack != nil {
		if len(stack) == 0 {
			return ""
		}

		// we only want the most recent frame of stack
		frame := stack[0]
		functionPointer := runtime.FuncForPC(uintptr(frame))
		if functionPointer != nil {
			file, line := functionPointer.FileLine(uintptr(frame))
			return fmt.Sprintf("%s, %s:%d", functionPointer.Name(), path.Base(file), line)
		}
	}

	return ""
}
