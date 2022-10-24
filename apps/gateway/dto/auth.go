package dto

type Login struct {
	Email string `json:"email" example:"6530000021@student.chula.ac.th" validate:"required,email"`
}

type MeResponse struct {
	User User `json:"user"`
}
