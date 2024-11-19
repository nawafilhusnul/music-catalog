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

func Test_service_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockRepo := NewMockrepository(ctrlMock)

	type args struct {
		ctx context.Context
		req *membershipsmodel.SignUpRequest
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
				req: &membershipsmodel.SignUpRequest{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.ctx, args.req.Email, args.req.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(args.ctx, gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when get user",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignUpRequest{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.ctx, args.req.Email, args.req.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when create user",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignUpRequest{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.ctx, args.req.Email, args.req.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(args.ctx, gomock.Any()).Return(assert.AnError)
			},
		},
		{
			name: "failed because user already exists",
			args: args{
				ctx: context.Background(),
				req: &membershipsmodel.SignUpRequest{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.ctx, args.req.Email, args.req.Username, uint(0)).Return(&membershipsmodel.User{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			s := &service{
				cfg:  &configs.Config{},
				repo: mockRepo,
			}
			if err := s.SignUp(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
