package tracks

import (
	"context"

	"github.com/gin-gonic/gin"
	spotifymodel "github.com/nawafilhusnul/music-catalog/internal/models/spotify"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=tracks
type service interface {
	Search(ctx context.Context, query string, pageSize, pageIndex int) (*spotifymodel.Tracks, error)
}

type handler struct {
	*gin.Engine
	svc service
}

func NewHandler(svc service, api *gin.Engine) *handler {
	return &handler{svc: svc, Engine: api}
}

func (h *handler) RegisterRoutes() {
	tracksRouter := h.Group("/tracks")

	tracksRouter.GET("/search", h.SearchTracks)
}
