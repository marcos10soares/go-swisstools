package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

// errors from pkg/errors have StackTrace
type errStack interface {
	StackTrace() errors.StackTrace
}

// stackTrace returns the stack trace with given error by the formatter.
// If the error is not traceable, nil is returned.
func stackTrace(err error) errors.StackTrace {
	if tracer, ok := err.(errStack); ok {
		stack := tracer.StackTrace()

		return stack
	}

	return nil
}

func getStackStrAndCauseFromErr(err error) (stackStr string, cause string) {
	stack := stackTrace(errors.Cause(err))

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

	return stackStr, cause
}
