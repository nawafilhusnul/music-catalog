package memberships

import (
	"context"
	"errors"

	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	"github.com/nawafilhusnul/music-catalog/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignIn(ctx context.Context, req *membershipsmodel.SignInRequest) (*membershipsmodel.SignInResponse, error) {
	userDetail, err := s.repo.GetUser(ctx, req.Email, req.Password, 0)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Msg("failed to get user detail from database")
		return nil, err
	}

	if userDetail == nil {
		return nil, errors.New("email and password not match")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("email and password not match")
	}

	accessToken, err := jwt.CreateToken(int64(userDetail.ID), userDetail.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create access token")
		return nil, err
	}

	return &membershipsmodel.SignInResponse{
		AccessToken: accessToken,
	}, nil
}
