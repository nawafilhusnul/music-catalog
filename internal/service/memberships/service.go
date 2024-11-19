package memberships

import (
	"context"

	"github.com/nawafilhusnul/music-catalog/internal/configs"
	membershipmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

type repository interface {
	CreateUser(ctx context.Context, model *membershipmodel.User) error
	GetUser(ctx context.Context, email, username string, id uint) (*membershipmodel.User, error)
}

type service struct {
	cfg  *configs.Config
	repo repository
}

func NewService(cfg *configs.Config, repo repository) *service {
	return &service{cfg: cfg, repo: repo}
}
