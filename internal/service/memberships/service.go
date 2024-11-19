package memberships

import (
	"context"

	"github.com/nawafilhusnul/music-catalog/internal/configs"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

//go:generate mockgen -source=service.go -destination=service_mock.go -package=memberships
type repository interface {
	CreateUser(ctx context.Context, model *membershipsmodel.User) error
	GetUser(ctx context.Context, email, username string, id uint) (*membershipsmodel.User, error)
}

type service struct {
	cfg  *configs.Config
	repo repository
}

func NewService(cfg *configs.Config, repo repository) *service {
	return &service{cfg: cfg, repo: repo}
}
