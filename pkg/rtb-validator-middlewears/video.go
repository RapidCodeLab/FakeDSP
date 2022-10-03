package rtb_validator_middlewears

import (
	"errors"

	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func videoObjectValidator(obj *openrtb2.Video) (bool, error) {

	if len(obj.MIMEs) < 1 {
		return false, errors.New("video object at least one mimes value required")
	}
	if len(obj.Protocols) < 1 {
		return false, errors.New("video asset at least one protocols value required")
	}
	if obj.MaxDuration < 1 || obj.MinDuration < 1 || obj.MinDuration > obj.MaxDuration {
		return false, errors.New("video asset minduration or maxduration values not applicable")
	}

	return true, nil
}
