package tracks

import (
	"context"

	spotifymodel "github.com/nawafilhusnul/music-catalog/internal/models/spotify"
	"github.com/nawafilhusnul/music-catalog/internal/repository/spotify"
	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotifymodel.Tracks, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize
	trackDetails, err := s.spotifyOutbound.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error searching tracks")
		return nil, err
	}

	res := modelToResponse(trackDetails)

	return res, nil
}

func modelToResponse(data *spotify.SearchResponse) *spotifymodel.Tracks {
	if data == nil {
		return nil
	}

	tracks := &spotifymodel.Tracks{
		HREF:     data.Tracks.HREF,
		Limit:    data.Tracks.Limit,
		Offset:   data.Tracks.Offset,
		Previous: data.Tracks.Previous,
		Total:    data.Tracks.Total,
		Items:    make([]spotifymodel.Track, len(data.Tracks.Items)),
	}

	for i, item := range data.Tracks.Items {
		tracks.Items[i] = spotifymodel.Track{
			Album:        mapAlbum(item.Album),
			Artists:      mapArtists(item.Artists),
			DiscNumber:   item.DiscNumber,
			DurationMs:   item.DurationMs,
			Explicit:     item.Explicit,
			ExternalIDs:  spotifymodel.ExternalIDs(item.ExternalIDs),
			ExternalURLs: item.ExternalURLs.Spotify,
			Href:         item.Href,
			ID:           item.ID,
			IsPlayable:   item.IsPlayable,
			LinkedFrom:   item.LinkedFrom,
			Restrictions: item.Restrictions.Reason,
			Name:         item.Name,
			Popularity:   item.Popularity,
			PreviewURL:   item.PreviewURL,
			TrackNumber:  item.TrackNumber,
			Type:         item.Type,
			URI:          item.URI,
			IsLocal:      item.IsLocal,
		}
	}

	return tracks
}

func mapAlbum(album spotify.Album) spotifymodel.Album {
	return spotifymodel.Album{
		AlbumType:            album.AlbumType,
		TotalTracks:          album.TotalTracks,
		AvailableMarkets:     album.AvailableMarkets,
		ExternalURLs:         album.ExternalURLs.Spotify,
		Href:                 album.Href,
		ID:                   album.ID,
		Images:               mapImages(album.Images),
		Name:                 album.Name,
		ReleaseDate:          album.ReleaseDate,
		ReleaseDatePrecision: album.ReleaseDatePrecision,
		Restrictions:         album.Restrictions.Reason,
		Type:                 album.Type,
		URI:                  album.URI,
		Artists:              mapArtists(album.Artists),
	}
}

func mapImages(images []spotify.Image) []string {
	if len(images) == 0 {
		return []string{}
	}

	result := make([]string, len(images))
	for i, image := range images {
		result[i] = image.URL
	}
	return result
}

func mapArtists(artists []spotify.Artist) []string {
	if len(artists) == 0 {
		return []string{}
	}

	result := make([]string, len(artists))
	for i, artist := range artists {
		result[i] = artist.Name
	}
	return result
}
