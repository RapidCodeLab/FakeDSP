package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	rtb_validator_middlewears "github.com/RapidCodeLab/fakedsp/pkg/rtb-validator-middlewears"
	"github.com/google/uuid"
	"github.com/mxmCherry/openrtb/v16/native1/request"
	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func NativeHandler(w http.ResponseWriter, r *http.Request, ads AdsDB) {
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
		_, err = w.Write(errorMsgJSON)
		if err != nil {
			fmt.Printf("response write error: %+v", err)
		}
		return
	}

	val := r.Context().Value(rtb_validator_middlewears.BidRequestContextKey).(openrtb2.BidRequest)

	bids := make([]openrtb2.Bid, 0, len(val.Imp)*4)

	// For each seat in demand

	seats := 2
	seatBids := make([]openrtb2.SeatBid, 0, seats)

	// One Bid object for every Native, Banner, Video, Audio object
	// in every Imp object identified with mtype && impid
	for i, v := range val.Imp {

		if i > impObjectsLimit {
			continue
		}

		if v.Banner != nil {
			a, err := ads.GetBanner(0, *v.Banner.W, *v.Banner.H)
			if err != nil {
				continue
			}
			bid := openrtb2.Bid{
				ID:    uuid.NewString(),
				ImpID: v.ID,
				MType: openrtb2.MarkupBanner,
				AdM:   a,
				Price: randomPrice(),
			}
			bids = append(bids, bid)
		}

		if v.Native != nil {
			native := request.Request{}
			err := json.Unmarshal([]byte(v.Native.Request), &native)
			if err != nil {
				continue
			}

			for idx := 0; idx < int(native.PlcmtCnt); idx++ {
				a := ads.GetNative(0)
				bid := openrtb2.Bid{
					ID:    uuid.NewString(),
					ImpID: v.ID,
					MType: openrtb2.MarkupNative,
					AdM:   a,
					Price: randomPrice(),
				}
				bids = append(bids, bid)
			}
		}

		if v.Video != nil {
			vast := ads.GetVideo(0, i)

			bid := openrtb2.Bid{
				ID:    uuid.NewString(),
				ImpID: v.ID,
				MType: openrtb2.MarkupVideo,
				AdM:   vast,
				Price: randomPrice(),
			}
			bids = append(bids, bid)
		}

		if v.Audio != nil {
			bid := openrtb2.Bid{
				ID:    uuid.NewString(),
				ImpID: v.ID,
				MType: openrtb2.MarkupAudio,
				Price: randomPrice(),
			}
			bids = append(bids, bid)
		}
	}

	seatBid := openrtb2.SeatBid{
		Seat: ads.GetSeat(0),
		Bid:  bids,
	}

	seatBids = append(seatBids, seatBid)

	br := openrtb2.BidResponse{}
	br.ID = val.ID
	br.SeatBid = seatBids

	brJSON, err := json.Marshal(br)
	if err != nil {
		errorMsg := ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  "unexpected jsom error",
		}
		errorMsgJSON, err := json.Marshal(errorMsg)
		if err != nil {
			fmt.Printf("error marshaling errorMsg: %+v", err)
		}

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(errorMsgJSON)
		if err != nil {
			fmt.Printf("response write error: %+v", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(brJSON)
	if err != nil {
		fmt.Printf("response write error: %+v", err)
	}
}
