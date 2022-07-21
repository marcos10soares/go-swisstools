/*
This package is an example for log usage
*/
package main

import (
	"fmt"
	"swisstools/app/log/fakepkg"
	"swisstools/app/log/fakepkg/nestedpkg"
	"swisstools/pkg/log"

	"github.com/pkg/errors"
)

const ErrFakeErrorMsg = "this is a fake error"

// var ErrFakeError = errors.New("this is a fake error")

func main() {
	log.SetGlobalLevel(log.DebugLevel)

	log.Println("msg", "test")

	newLogger := log.NewWith("service", "awesome")

	log.Debug("this is a debug msg")

	newLogger.Info("awesome service")

	log.Info("this is a info msg")

	log.Error(errors.New(ErrFakeErrorMsg), "could not open file")

	err := fakepkg.NestedErrorOneLevel()
	log.Error(err, "error in main because x")

	// this won't work, because for this to work
	// it would return the line where the error is declared,
	// instead of where it's thrown

	// this can lead to panics if the error was not wrapped!
	// unwrappedError := errors.Unwrap(err)

	// log.Info("err matches ErrLevel2Error: ", errors.Is(err, errors.New(fakepkg.ErrLevel2ErrorMsg)))
	// log.Info("unwrappedError matches ErrLevel2Error: ", errors.Is(unwrappedError, errors.New(fakepkg.ErrLevel2ErrorMsg)))
	// log.Info("err matches ErrLevel1Error: ", errors.Is(err, errors.New(fakepkg.ErrLevel1ErrorMsg)))
	// log.Info("unwrappedError matches ErrLevel1Error: ", errors.Is(unwrappedError, errors.New(fakepkg.ErrLevel1ErrorMsg)))

	// alternative solution 🤷🏻‍♂️ - do not work
	// log.Info("alternative - err matches ErrLevel2Error: ", err.Error() == fakepkg.ErrLevel2ErrorMsg)
	// log.Info("alternative - unwrappedError matches ErrLevel2Error: ", unwrappedError.Error() == fakepkg.ErrLevel2ErrorMsg)
	// log.Info("alternative - err matches ErrLevel1Error: ", err.Error() == fakepkg.ErrLevel1ErrorMsg)
	// log.Info("alternative - unwrappedError matches ErrLevel1Error: ", unwrappedError.Error() == fakepkg.ErrLevel1ErrorMsg)

	// proper way to check
	log.Info("causer str comp - err matches fakepkg.ErrLevel1Error: ", errors.Cause(err).Error() == fakepkg.ErrLevel1ErrorMsg)
	log.Info("causer str comp - err matches fakepkg.ErrLevel2Error: ", errors.Cause(err).Error() == fakepkg.ErrLevel2ErrorMsg)
	log.Info("causer str comp - err matches nestedpkg.ErrFakeErrorMsg: ", errors.Cause(err).Error() == nestedpkg.ErrFakeErrorMsg)

	log.Info("cause", errors.Cause(err))

	log.Error(err, "err")

	log.Error(errors.Cause(err), "cause-error")

	// ideally
	// log.Error(errors.Cause(err), "cause-error", err.Error())

	log.Error(errors.New(ErrFakeErrorMsg), "new-fake-error-msg")

	log.Error(fmt.Errorf("this is a basic error with no causer or stack"), "basic-err")
}
