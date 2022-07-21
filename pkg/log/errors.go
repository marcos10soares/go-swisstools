package log

import (
	"github.com/pkg/errors"
)

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
