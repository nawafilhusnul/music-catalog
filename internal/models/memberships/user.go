package memberships

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"column:email;not null;unique"`
	Username  string `gorm:"column:username;not null;unique"`
	Password  string `gorm:"column:password;not null"`
	UpdatedBy uint   `gorm:"column:updated_by;not null"`
	CreatedBy uint   `gorm:"column:created_by;not null"`
	DeletedBy uint   `gorm:"column:deleted_by;default:null"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}
