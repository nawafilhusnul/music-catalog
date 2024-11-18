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
	DeletedBy uint   `gorm:"column:deleted_by"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
