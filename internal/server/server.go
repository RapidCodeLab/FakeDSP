package server

import (
	"net"
	"net/http"

	rtb_validator_middlewears "github.com/RapidCodeLab/fakedsp/pkg/rtb-validator-middlewears"
	"github.com/gorilla/mux"
)

const (
	nativePath      = "/openrtb"
	impObjectsLimit = 5 // Imp objects annount limit
)

type ErrorResponse struct {
	Status int
	Error  string
}

type Logger interface{}

type Config interface {
	GetListenAddr() string
	GetListenNetwork() string
}

type server struct {
	logger Logger
	config Config
	http   *http.Server
}

func New(l Logger, c Config) *server {
	return &server{
		logger: l,
		config: c,
	}
}

func (s *server) Start() error {

	r := mux.NewRouter()

	r.HandleFunc(nativePath, NativeHandler).
		Methods(http.MethodPost)

	r.Use(rtb_validator_middlewears.ValidateOpenRTBBidRequestMiddleware)

	s.http = &http.Server{
		Handler: r,
	}

	l, err := net.Listen(s.config.GetListenNetwork(), s.config.GetListenAddr())
	if err != nil {
		return err
	}

	return s.http.Serve(l)

}
