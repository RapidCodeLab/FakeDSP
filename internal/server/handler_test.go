package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	rtb_validator_middlewears "github.com/RapidCodeLab/fakedsp/pkg/rtb-validator-middlewears"
	"github.com/mxmCherry/openrtb/v16/openrtb2"
)

func TestNativeHandlerResponseOK(t *testing.T) {
	br := openrtb2.BidRequest{}
	req, err := http.NewRequest(http.MethodPost, nativePath, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), rtb_validator_middlewears.BidRequestContextKey, br)
	nr := req.WithContext(ctx)

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			NativeHandler(w, r, nil)
		})
	handler.ServeHTTP(res, nr)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
