package dto

type RegisterRequest struct {
	Login string `json:"login"`
	Email string `json:"email"`
	Password string `json:"password"`
}
