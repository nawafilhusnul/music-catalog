package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	membershipmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

type service interface {
	SignUp(ctx context.Context, req *membershipmodel.SignUpRequest) error
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
