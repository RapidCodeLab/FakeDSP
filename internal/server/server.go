package server

import (
	"math/rand"
	"net"
	"net/http"

	rtb_validator_middlewears "github.com/RapidCodeLab/fakedsp/pkg/rtb-validator-middlewears"
	"github.com/gorilla/mux"
)

const (
	nativePath      = "/openrtb"
	impObjectsLimit = 5 // Imp objects ammount limit
)

type ErrorResponse struct {
	Status int
	Error  string
}

type AdsDB interface {
	GetSeat(seatID int) string
	GetNative(seatID int) string
	GetBanner(seatID int, width, height int64) (string, error)
	GetVideo(seatID, itemID int) string
	GetAudio(seatID, itemID int) string
}

type Logger interface{}

type Config interface {
	GetListenAddr() string
	GetListenNetwork() string
	GetAdsDatabasePath() string
}

type server struct {
	logger Logger
	config Config
	http   *http.Server
	adsDB  AdsDB
}

func New(l Logger, c Config, db AdsDB) *server {
	return &server{
		logger: l,
		config: c,
		adsDB:  db,
	}
}

func (s *server) Start() error {
	r := mux.NewRouter()

	r.HandleFunc(nativePath, func(w http.ResponseWriter, r *http.Request) {
		NativeHandler(w, r, s.adsDB)
	}).
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

func randomPrice() float64 {
	return 0.5 + rand.Float64()*(1.5-0.5)
}
