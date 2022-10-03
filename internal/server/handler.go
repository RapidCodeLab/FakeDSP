package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	rtb_validator_middlewears "github.com/RapidCodeLab/fakedsp/pkg/rtb-validator-middlewears"
	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func NativeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value(rtb_validator_middlewears.BidRequestContextKey) == nil &&
		r.Context().Value(rtb_validator_middlewears.BidRequestContextErrorKey) != nil {

		errorMsg := ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  r.Context().Value(rtb_validator_middlewears.BidRequestContextErrorKey).(error).Error(),
		}

		errorMsgJSON, err := json.Marshal(errorMsg)
		if err != nil {
			fmt.Printf("error marshaling errorMsg: %+v", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorMsgJSON)
		return
	}

	val := r.Context().Value(rtb_validator_middlewears.BidRequestContextKey).(openrtb2.BidRequest)

	bids := make([]openrtb2.Bid, 0, len(val.Imp)*4)

	//For each seat in demand

	seats := 2
	seatBids := make([]openrtb2.SeatBid, 0, seats)

	//One Bid object for every Native, Banner, Video, Audio object
	//in every Imp object identified with mtype && impid
	for i, v := range val.Imp {

		if i > impObjectsLimit {
			continue
		}

		if v.Banner != nil {
			bid := openrtb2.Bid{
				ImpID: v.ID,
				MType: openrtb2.MarkupBanner,
			}
			bids = append(bids, bid)
		}

		if v.Native != nil {
			bid := openrtb2.Bid{
				ImpID: v.ID,
				MType: openrtb2.MarkupNative,
			}
			bids = append(bids, bid)
		}

		if v.Video != nil {
			bid := openrtb2.Bid{
				ImpID: v.ID,
				MType: openrtb2.MarkupVideo,
			}
			bids = append(bids, bid)
		}
		if v.Audio != nil {
			bid := openrtb2.Bid{
				ImpID: v.ID,
				MType: openrtb2.MarkupAudio,
			}
			bids = append(bids, bid)
		}
	}

	seatBid := openrtb2.SeatBid{
		Seat: "agency",
		Bid:  bids,
	}

	seatBids = append(seatBids, seatBid)

	fmt.Printf("Bids: %+v", seatBids)

	w.WriteHeader(http.StatusOK)
}
