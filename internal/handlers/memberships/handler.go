package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

//go:generate mockgen -source=handler.go -destination=handler_mock.go -package=memberships
type service interface {
	SignUp(ctx context.Context, req *membershipsmodel.SignUpRequest) error
}

type handler struct {
	*gin.Engine
	svc service
}

func NewHandler(svc service, api *gin.Engine) *handler {
	return &handler{svc: svc, Engine: api}
}

func (h *handler) RegisterRoutes() {
	membershipRouter := h.Group("/memberships")

	membershipRouter.POST("/signup", h.SignUp)

}
