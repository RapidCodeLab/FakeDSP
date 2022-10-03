package rtb_validator_middlewears

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

type bidRequestContext string

var BidRequestContextKey = bidRequestContext("bidRequest")
var BidRequestContextErrorKey = bidRequestContext("bidRequestError")

func ValidateOpenRTBBidRequestMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var (
			isValid bool
			req     openrtb2.BidRequest
		)

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()
		err := d.Decode(&req)
		if err != nil {
			ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
		} else if len(req.Imp) < 1 {
			ctx = context.WithValue(ctx, BidRequestContextErrorKey, errors.New("at least one Imp object required"))
		} else if len(req.ID) < 1 {
			ctx = context.WithValue(ctx, BidRequestContextErrorKey, errors.New("bid request unique id required"))
		}

		//Imp objects is required by specificaions
		for _, val := range req.Imp {
			switch true {
			case len(val.ID) < 1:
				ctx = context.WithValue(ctx, BidRequestContextErrorKey, errors.New("imp object unique id required"))
			case val.Native != nil:
				isValid, err = nativeObjectValidator(val.Native)
				if err != nil {
					ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
				}
			case val.Banner != nil:
				isValid, err = bannerObjectValidator(val.Banner)
				if err != nil {
					ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
				}
			case val.Video != nil:
				isValid, err = videoObjectValidator(val.Video)
				if err != nil {
					ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
				}
			case val.Audio != nil:
				isValid, err = audioObjectValidator(val.Audio)
				if err != nil {
					ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
				}
			default:
				ctx = context.WithValue(ctx, BidRequestContextErrorKey, errors.New("unsupported type of imp object"))
			}
		}

		//implement these validators if you need some custom validations, not required by specifications
		if req.Site != nil {
			isValid, err = siteObjectValidator(req.Site)
			if err != nil {
				ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
			}
		}

		if req.App != nil {
			isValid, err = appObjectValidator(req.App)
			if err != nil {
				ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
			}
		}

		if req.Device != nil {
			isValid, err = deviceObjectValidator(req.Device)
			if err != nil {
				ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
			}
		}

		if req.User != nil {
			isValid, err = userObjectValidator(req.User)
			if err != nil {
				ctx = context.WithValue(ctx, BidRequestContextErrorKey, err)
			}
		}

		if isValid {
			ctx = context.WithValue(ctx, BidRequestContextKey, req)
		}

		nr := r.WithContext(ctx)
		h.ServeHTTP(w, nr)
	})
}
