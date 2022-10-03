package rtb_validator_middlewears

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mxmCherry/openrtb/v16/native1/request"
	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func nativeObjectValidator(obj *openrtb2.Native) (bool, error) {

	var r request.Request

	err := json.Unmarshal([]byte(obj.Request), &r)
	if err != nil {
		return false, fmt.Errorf("wrong native request: %+v", err)
	}

	if len(r.Assets) < 1 {
		return false, errors.New("at least one asset object required")
	}

	for _, val := range r.Assets {
		switch true {
		case val.ID < 1:
			return false, errors.New("asset id must be greather than zero")
		case val.Title != nil:
			if val.Title.Len < 25 {
				return false, errors.New("title asset len must not be less than 25")
			}
		case val.Video != nil:
			if len(val.Video.MIMEs) < 1 {
				return false, errors.New("video asset at least one mimes value required")
			}
			if len(val.Video.Protocols) < 1 {
				return false, errors.New("video asset at least one protocols value required")
			}
			if val.Video.MaxDuration < 1 || val.Video.MinDuration < 1 || val.Video.MinDuration > val.Video.MaxDuration {
				return false, errors.New("video asset minduration or maxduration values not applicable")
			}
		case val.Data != nil:
			if val.Data.Type < 1 || (val.Data.Type > 12 && val.Data.Type < 500) {
				return false, errors.New("data asset type must be from 1 to 12 or greather than 500")
			}
		default:
		}
	}

	return true, nil
}
