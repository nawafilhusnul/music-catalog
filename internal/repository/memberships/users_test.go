package memberships

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	membershipsmodel "github.com/nawafilhusnul/music-catalog/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	assert.NoError(t, err)

	type args struct {
		ctx   context.Context
		model *membershipsmodel.User
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
				model: &membershipsmodel.User{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.UpdatedBy,
						args.model.CreatedBy,
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
				model: &membershipsmodel.User{
					Email:    "test@example.com",
					Username: "testuser",
					Password: "testpassword",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs(
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						sqlmock.AnyArg(),
						args.model.Email,
						args.model.Username,
						args.model.Password,
						args.model.UpdatedBy,
						args.model.CreatedBy,
					).
					WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			if err := r.CreateUser(tt.args.ctx, tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	now := time.Now()

	type args struct {
		ctx      context.Context
		email    string
		username string
		id       uint
	}
	tests := []struct {
		name    string
		args    args
		want    *membershipsmodel.User
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				email:    "test@example.com",
				username: "testuser",
			},
			want: &membershipsmodel.User{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				Email:    "test@example.com",
				Username: "testuser",
				Password: "testpassword",
			},
			wantErr: false,
			mockFn: func(args args) {

				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
					WithArgs(args.email, args.username, args.id, 1).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "created_at", "updated_at", "email", "username", "password"}).
							AddRow(1, now, now, "test@example.com", "testuser", "testpassword"),
					)

			},
		},
		{
			name: "error",
			args: args{
				ctx:      context.Background(),
				email:    "test@example.com",
				username: "testuser",
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {

				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
					WithArgs(args.email, args.username, args.id, 1).
					WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)

			r := &repository{
				db: gormDB,
			}

			got, err := r.GetUser(tt.args.ctx, tt.args.email, tt.args.username, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetUser() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
