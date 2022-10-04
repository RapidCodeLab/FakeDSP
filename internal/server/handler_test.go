package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNativeHandlerResponseOK(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, nativePath, nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			NativeHandler(w, r, nil)
		})
	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
