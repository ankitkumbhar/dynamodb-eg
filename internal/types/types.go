package types

type Music struct {
	Artist      string `json:"artist"`
	SongTitle   string `json:"song_title"`
	Description string `json:"description"`
	Views       int64  `json:"views"`
}

type GetItemResponse struct {
	Data []Music
}

type StoreItemRequest struct {
	Music
}
