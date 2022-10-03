package rtb_validator_middlewears

import "github.com/mxmCherry/openrtb/v16/openrtb2"

// banner object has no required fileds, implement it on your mind
func bannerObjectValidator(obj *openrtb2.Banner) (bool, error) {

	return true, nil
}
