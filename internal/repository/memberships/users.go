package memberships

import (
	"context"

	membershipmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

func (r *repository) CreateUser(ctx context.Context, model *membershipmodel.User) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *repository) GetUser(ctx context.Context, email, username string, id uint) (*membershipmodel.User, error) {
	var user membershipmodel.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		Or("username = ?", username).
		Or("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
