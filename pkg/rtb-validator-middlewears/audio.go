package rtb_validator_middlewears

import (
	"errors"

	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func audioObjectValidator(obj *openrtb2.Audio) (bool, error) {
	if len(obj.MIMEs) < 1 {
		return false, errors.New("audio object at least one mimes value required")
	}
	return true, nil
}
