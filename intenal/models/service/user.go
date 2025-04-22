package service

type AddUserRequest struct {
	Login       string `json:"login"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type AddUserResponse struct {
	ID int64 `json:"id"`
}
