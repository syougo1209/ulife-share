package dto

type UserResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
}

type UserRequest struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Introduction string `json:"introduction"`
}
