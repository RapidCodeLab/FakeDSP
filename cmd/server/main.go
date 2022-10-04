package main

import (
	"log"

	"github.com/RapidCodeLab/fakedsp/internal/server"
	"github.com/RapidCodeLab/fakedsp/pkg/ads_db"
	"github.com/RapidCodeLab/fakedsp/pkg/config"
)

func main() {

	cfg, err := config.GetHTTPServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	adsDB, err := ads_db.New(cfg.GetAdsDatabasePath())
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(nil, cfg, adsDB)

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}

}
