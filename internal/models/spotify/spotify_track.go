package spotify

type Album struct {
	AlbumType            string   `json:"album_type"`
	TotalTracks          int      `json:"total_tracks"`
	AvailableMarkets     []string `json:"available_markets"`
	ExternalURLs         string   `json:"external_urls"`
	Href                 string   `json:"href"`
	ID                   string   `json:"id"`
	Images               []string `json:"images"`
	Name                 string   `json:"name"`
	ReleaseDate          string   `json:"release_date"`
	ReleaseDatePrecision string   `json:"release_date_precision"`
	Restrictions         string   `json:"restrictions"`
	Type                 string   `json:"type"`
	URI                  string   `json:"uri"`
	Artists              []string `json:"artists"`
}

type ExternalIDs struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type Track struct {
	Album        Album       `json:"album"`
	Artists      []string    `json:"artists"`
	DiscNumber   int         `json:"disc_number"`
	DurationMs   int         `json:"duration_ms"`
	Explicit     bool        `json:"explicit"`
	ExternalIDs  ExternalIDs `json:"external_ids"`
	ExternalURLs string      `json:"external_urls"`
	Href         string      `json:"href"`
	ID           string      `json:"id"`
	IsPlayable   bool        `json:"is_playable"`
	LinkedFrom   struct{}    `json:"linked_from"`
	Restrictions string      `json:"restrictions"`
	Name         string      `json:"name"`
	Popularity   int         `json:"popularity"`
	PreviewURL   string      `json:"preview_url"`
	TrackNumber  int         `json:"track_number"`
	Type         string      `json:"type"`
	URI          string      `json:"uri"`
	IsLocal      bool        `json:"is_local"`
}

type Tracks struct {
	HREF     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     *string `json:"next"`
	Offset   int     `json:"offset"`
	Previous *string `json:"previous"`
	Total    int     `json:"total"`
	Items    []Track `json:"items"`
}
