package fakepkg

import (
	"github.com/marcos10soares/go-swisstools/pkg/log"

	"github.com/pkg/errors"
)

func NestedErrorOneLevel() error {
	err := nestedErrorTwoLevels()

	// both comparisons work
	// log.Info("str comp - err matches ErrLevel2Error: ", err.Error() == ErrLevel2ErrorMsg)
	// log.Info("alternative causer str comp - err matches ErrLevel2Error: ", errors.Cause(err).Error() == ErrLevel2ErrorMsg)

	// do something
	if errors.Cause(err).Error() == ErrLevel2ErrorMsg {
		log.Debug("wrapping errors")
		// return errors.WithMessage(err, ErrLevel1ErrorMsg) // this looses the stack trace

		// return errors.Wrap(errors.New(ErrLevel1ErrorMsg), err.Error())
		// return errors.Wrap(err, ErrLevel1ErrorMsg)

		// this makes ErrLevel2Error identifiable in main, but with no stack trace
		// return errors.WithMessage(err, ErrLevel1ErrorMsg)

		// this makes ErrLevel1Error identifiable in main, with stack trace to this line
		// return errors.Wrap(errors.New(ErrLevel1ErrorMsg), err.Error())

		// return err // this makes ErrLevel2Error identifiable in main, with stack trace

		return errors.New(ErrLevel1ErrorMsg)
	}

	// this makes ErrLevel1Error identifiable in main, with stack trace to this line
	return errors.Wrap(err, ErrLevel1ErrorMsg)
}
