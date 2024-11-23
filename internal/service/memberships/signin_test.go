package memberships

import (
	"context"
	"testing"

	"github.com/nawafilhusnul/music-catalog/internal/configs"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignIn(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)
	type args struct {
		ctx context.Context
		req *membershipsmodel.SignInRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignInRequest{
					Email:    "test@example.com",
					Password: "qwerty123",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(gomock.Any(), args.req.Email, args.req.Password, uint(0)).Return(&membershipsmodel.User{
					Model: gorm.Model{
						ID: 1,
					},
					Email:    "test@example.com",
					Username: "testuser",
					Password: "$2a$10$RvosoJD.cebdgXS1bB/teuB0.IE4ZIbeHSIFPoy3tC5tnAxVZ12rK",
				}, nil)
			},
		},
		{
			name: "failed when get user detail",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignInRequest{
					Email:    "test@example.com",
					Password: "qwerty123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(gomock.Any(), args.req.Email, args.req.Password, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when user not found",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignInRequest{
					Email:    "test@example.com",
					Password: "qwerty123",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(gomock.Any(), args.req.Email, args.req.Password, uint(0)).Return(nil, nil)
			},
		},
		{
			name: "failed when password not match",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignInRequest{
					Email:    "test@example.com",
					Password: "-----",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(gomock.Any(), args.req.Email, args.req.Password, uint(0)).Return(&membershipsmodel.User{
					Password: "$2a$10$RvosoJD.cebdgXS1bB/teuB0.IE4ZIbeHSIFPoy3tC5tnAxVZ12rK",
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockFn != nil {
				tt.mockFn(tt.args)
			}
			s := &service{
				repo: mockRepo,
				cfg: &configs.Config{
					Service: configs.Service{
						SecretJWT: "secretmock",
					},
				},
			}
			got, err := s.SignIn(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.SignIn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.NotEmpty(t, got)
				return
			}

			assert.Empty(t, got)
		})
	}
}
