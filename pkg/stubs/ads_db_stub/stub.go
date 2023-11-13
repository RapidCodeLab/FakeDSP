package ads_db_stub

type db_stub struct{}

func New() *db_stub {
	return &db_stub{}
}

func (db *db_stub) GetSeat(seatID int) string {
	return ""
}

func (db *db_stub) GetNative(seatID int) string {
	return ""
}

func (db *db_stub) GetBanner(seatID int, width, height int64) (string, error) {
	return "", nil
}

func (db *db_stub) GetVideo(seatID, itemID int) string {
	return ""
}

func (db *db_stub) GetAudio(seatID, itemID int) string {
	return ""
}
