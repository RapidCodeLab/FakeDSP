package ads_db

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"

	"github.com/haxqer/vast"
)

type AdsDB struct {
	seats []seat
}

type seat struct {
	Name    string
	Natives []native
	Banners []banner
	Videos  []video
	Audios  []audio
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
	URI string
}

func New(path string) (*AdsDB, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var v []seat

	d := json.NewDecoder(f)
	d.DisallowUnknownFields()
	err = d.Decode(&v)

	if err != nil {
		return nil, err
	}

	return &AdsDB{
		seats: v,
	}, nil
}

func (db *AdsDB) GetSeat(seatID int) string {
	return db.seats[seatID].Name
}

func (db *AdsDB) GetNative(seatID, itemID int) string {
	a := db.seats[seatID].Natives[itemID]
	return fmt.Sprintf("<div><a href=\"%s\"><img src=\"%s\"/><br>%s</a><br>%s</div>",
		a.Link, a.Image, a.Title, a.Text)
}

func (db *AdsDB) GetBanner(seatID, itemID int) string {
	a := db.seats[seatID].Banners[itemID]
	return fmt.Sprintf("<a href=\"%s\"><img srec=\"%s\"/></a>",
		a.Link, a.Image)

}

func (db *AdsDB) GetVideo(seatID, itemID int) string {
	return prepareVAST(db.seats[seatID].Videos[itemID])
}

func (db *AdsDB) GetAudio(seatID, itemID int) string {
	return ""
	//return db.seats[seatID].Audios[itemID]
}

func prepareVAST(v video) string {

	o := &vast.VAST{
		Version: "3.0",
		Ads: []vast.Ad{
			{
				ID: "123",
				InLine: &vast.InLine{
					AdSystem: &vast.AdSystem{Name: "DSP"},
					AdTitle:  vast.CDATAString{CDATA: "adTitle"},
					Creatives: []vast.Creative{
						{
							Sequence: 0,
							Linear: &vast.Linear{
								VideoClicks: &vast.VideoClicks{
									ClickThroughs: []vast.VideoClick{
										{
											ID:  "1",
											URI: v.Link,
										},
									},
								},
								MediaFiles: []vast.MediaFile{
									{
										Delivery: "progressive",
										Type:     "video/mp4",
										URI:      v.URI,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	vastXMLText, _ := xml.Marshal(o)

	return string(vastXMLText)

}
