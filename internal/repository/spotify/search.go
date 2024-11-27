package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Restrictions struct {
	Reason string `json:"reason"`
}

type Artist struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Artists              []Artist     `json:"artists"`
}

type ExternalIDs struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

type Track struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIDs      ExternalIDs  `json:"external_ids"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       struct{}     `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
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

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type ArtistItem struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       []string     `json:"genres"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Images       []Image      `json:"images"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Artists struct {
	HREF     string       `json:"href"`
	Limit    int          `json:"limit"`
	Next     *string      `json:"next"`
	Offset   int          `json:"offset"`
	Previous *string      `json:"previous"`
	Total    int          `json:"total"`
	Items    []ArtistItem `json:"items"`
}

type AlbumItem struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Artists              []Artist     `json:"artists"`
}

type Albums struct {
	HREF     string      `json:"href"`
	Limit    int         `json:"limit"`
	Next     *string     `json:"next"`
	Offset   int         `json:"offset"`
	Previous *string     `json:"previous"`
	Total    int         `json:"total"`
	Items    []AlbumItem `json:"items"`
}

type Owner struct {
	ExternalURLs ExternalURLs `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
	DisplayName  string       `json:"display_name"`
}

type PlaylistItem struct {
	Collaborative bool         `json:"collaborative"`
	Description   string       `json:"description"`
	ExternalURLs  ExternalURLs `json:"external_urls"`
	Href          string       `json:"href"`
	ID            string       `json:"id"`
	Images        []Image      `json:"images"`
	Name          string       `json:"name"`
	Owner         Owner        `json:"owner"`
	Public        bool         `json:"public"`
	SnapshotID    string       `json:"snapshot_id"`
	Tracks        struct {
		Href  string `json:"href"`
		Total int    `json:"total"`
	} `json:"tracks"`
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type Playlists struct {
	HREF     string         `json:"href"`
	Limit    int            `json:"limit"`
	Next     *string        `json:"next"`
	Offset   int            `json:"offset"`
	Previous *string        `json:"previous"`
	Total    int            `json:"total"`
	Items    []PlaylistItem `json:"items"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type ShowItem struct {
	AvailableMarkets   []string     `json:"available_markets"`
	Copyrights         []Copyright  `json:"copyrights"`
	Description        string       `json:"description"`
	HTMLDescription    string       `json:"html_description"`
	Explicit           bool         `json:"explicit"`
	ExternalURLs       ExternalURLs `json:"external_urls"`
	Href               string       `json:"href"`
	ID                 string       `json:"id"`
	Images             []Image      `json:"images"`
	IsExternallyHosted bool         `json:"is_externally_hosted"`
	Languages          []string     `json:"languages"`
	MediaType          string       `json:"media_type"`
	Name               string       `json:"name"`
	Publisher          string       `json:"publisher"`
	Type               string       `json:"type"`
	URI                string       `json:"uri"`
	TotalEpisodes      int          `json:"total_episodes"`
}

type Shows struct {
	HREF     string     `json:"href"`
	Limit    int        `json:"limit"`
	Next     *string    `json:"next"`
	Offset   int        `json:"offset"`
	Previous *string    `json:"previous"`
	Total    int        `json:"total"`
	Items    []ShowItem `json:"items"`
}

type ResumePoint struct {
	FullyPlayed      bool `json:"fully_played"`
	ResumePositionMs int  `json:"resume_position_ms"`
}

type EpisodeItem struct {
	AudioPreviewURL      string       `json:"audio_preview_url"`
	Description          string       `json:"description"`
	HTMLDescription      string       `json:"html_description"`
	DurationMs           int          `json:"duration_ms"`
	Explicit             bool         `json:"explicit"`
	ExternalURLs         ExternalURLs `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	IsExternallyHosted   bool         `json:"is_externally_hosted"`
	IsPlayable           bool         `json:"is_playable"`
	Language             string       `json:"language"`
	Languages            []string     `json:"languages"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	ResumePoint          ResumePoint  `json:"resume_point"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Restrictions         Restrictions `json:"restrictions"`
}

type Episodes struct {
	HREF     string        `json:"href"`
	Limit    int           `json:"limit"`
	Next     *string       `json:"next"`
	Offset   int           `json:"offset"`
	Previous *string       `json:"previous"`
	Total    int           `json:"total"`
	Items    []EpisodeItem `json:"items"`
}

type Author struct {
	Name string `json:"name"`
}

type Narrator struct {
	Name string `json:"name"`
}

type AudiobookItem struct {
	Authors          []Author     `json:"authors"`
	AvailableMarkets []string     `json:"available_markets"`
	Copyrights       []Copyright  `json:"copyrights"`
	Description      string       `json:"description"`
	HTMLDescription  string       `json:"html_description"`
	Edition          string       `json:"edition"`
	Explicit         bool         `json:"explicit"`
	ExternalURLs     ExternalURLs `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	Images           []Image      `json:"images"`
	Languages        []string     `json:"languages"`
	MediaType        string       `json:"media_type"`
	Name             string       `json:"name"`
	Narrators        []Narrator   `json:"narrators"`
	Publisher        string       `json:"publisher"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	TotalChapters    int          `json:"total_chapters"`
}

type Audiobooks struct {
	HREF     string          `json:"href"`
	Limit    int             `json:"limit"`
	Next     *string         `json:"next"`
	Offset   int             `json:"offset"`
	Previous *string         `json:"previous"`
	Total    int             `json:"total"`
	Items    []AudiobookItem `json:"items"`
}

type SearchResponse struct {
	Tracks     Tracks     `json:"tracks"`
	Artists    Artists    `json:"artists"`
	Albums     Albums     `json:"albums"`
	Playlists  Playlists  `json:"playlists"`
	Shows      Shows      `json:"shows"`
	Episodes   Episodes   `json:"episodes"`
	Audiobooks Audiobooks `json:"audiobooks"`
}

func (o *outbound) Search(ctx context.Context, query string, limit, offset int) (*SearchResponse, error) {
	params := url.Values{}
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))

	basePath := "https://api.spotify.com/v1/search"
	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to create request to spotify")
		return nil, err
	}

	accessToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("failed to get token details")
		return nil, err
	}

	token := fmt.Sprintf("%s %s", tokenType, accessToken)
	req.Header.Set("Authorization", token)

	resp, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get token from spotify")
		return nil, err
	}
	defer resp.Body.Close()

	searchResponse := new(SearchResponse)
	err = json.NewDecoder(resp.Body).Decode(searchResponse)
	if err != nil {
		log.Error().Err(err).Msg("failed to unmarshal response from spotify")
		return nil, err
	}

	return searchResponse, nil
}
