package memberships

import (
	"context"

	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
)

func (r *repository) CreateUser(ctx context.Context, model *membershipsmodel.User) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *repository) GetUser(ctx context.Context, email, username string, id uint) (*membershipsmodel.User, error) {
	var user membershipsmodel.User

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
