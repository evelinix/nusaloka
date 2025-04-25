package dto

type UpdatePasswordRequest struct {
	Password string `json:"password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=NewPassword"`
}

type UpdateAvatarRequest struct {
	Avatar string `json:"avatar"`
}

type AccountResponse struct {
	Email string `json:"email"`
	Avatar string `json:"avatar"`
}
