package validate

import (
	"github.com/go-playground/validator/v10"
)

type RoleCreate struct {
	Name          string `json:"name" validate:"required,min=2,max=100"`
	PermissionIDs []uint `json:"permission_ids" validate:"omitempty,dive,min=1"`
}

func (r *RoleCreate) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type AssignRole struct {
	UserID uint `json:"user_id" validate:"required,min=1"`
	RoleID uint `json:"role_id" validate:"required,min=1"`
}

func (ar *AssignRole) Validate() error {
	validate := validator.New()
	return validate.Struct(ar)
}

type AssignPermission struct {
	RoleID       uint `json:"role_id" validate:"required,min=1"`
	PermissionID uint `json:"permission_id" validate:"required,min=1"`
}

func (ap *AssignPermission) Validate() error {
	validate := validator.New()
	return validate.Struct(ap)
}
