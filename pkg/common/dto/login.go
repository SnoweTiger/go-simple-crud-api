package dto

type LoginDTO struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRespDTO struct {
	Login string `json:"login"`
	Token string `json:"token"`
}
