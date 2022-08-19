package fakepkg

import (
	"github.com/marcos10soares/go-swisstools/app/log/fakepkg/nestedpkg"

	"github.com/pkg/errors"
)

func nestedErrorTwoLevels() error {
	err := nestedpkg.ReturnNestedErr()

	return errors.Wrap(err, ErrLevel2ErrorMsg)
}
