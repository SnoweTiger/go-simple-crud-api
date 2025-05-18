package dto

// Create user schema
type CreateUserDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// User response schema
type UserDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login,omitempty"`
}

// User data schema
type UserDataDTO struct {
	User     UserDTO      `json:"user"`
	Articles []ArticleDTO `json:"articles"`
}

// Change password schema
type ChangePassDTO struct {
	ID          int    `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
