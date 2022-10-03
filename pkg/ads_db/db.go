package ads_db

import (
	"encoding/json"
	"os"
)

type adsDB struct {
	seats []seat
}

type seat struct {
	seats   []string
	natives []native
	banners []banner
	videos  []video
	audios  []audio
}

type native struct {
	Title string
	Text  string
	Image string
	Link  string
}

type banner struct {
	Image string
	Link  string
}

type video struct {
	URI  string
	Link string
}

type audio struct {
	URI  string
	Link string
}

func New(path string) (*adsDB, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var v []seat

	d := json.NewDecoder(f)
	d.DisallowUnknownFields()
	err = d.Decode(&v)

	if err != nil {
		return nil, err
	}

	return &adsDB{
		seats: v,
	}, nil
}

func (db *adsDB) GetSeat(seatID, itemID int) string {
	return db.seats[seatID].seats[itemID]
}

func (db *adsDB) GetNative(seatID, itemID int) native {
	return db.seats[seatID].natives[itemID]
}

func (db *adsDB) GetBanner(seatID, itemID int) banner {
	return db.seats[seatID].banners[itemID]
}

func (db *adsDB) GetVideo(seatID, itemID int) video {
	return db.seats[seatID].videos[itemID]
}

func (db *adsDB) GetAudio(seatID, itemID int) audio {
	return db.seats[seatID].audios[itemID]
}
