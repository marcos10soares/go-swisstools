package nestedpkg

import "github.com/pkg/errors"

var ErrFakeErrorMsg = "even more nested error msg"

func ReturnNestedErr() error {
	return errors.New(ErrFakeErrorMsg)
}
