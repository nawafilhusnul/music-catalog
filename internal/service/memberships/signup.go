package memberships

import (
	"context"
	"errors"

	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(ctx context.Context, req *membershipsmodel.SignUpRequest) error {
	existingUser, err := s.repo.GetUser(ctx, req.Email, req.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("failed to get user")
		return err
	}

	if existingUser != nil {
		return errors.New("user already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("failed to hash password")
		return err
	}

	user := &membershipsmodel.User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(pass),
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return err
	}

	return nil
}
