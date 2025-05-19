package dto

type LoginDTO struct {
	Login    string `json:"login" `
	Password string `json:"password"`
}

type LoginRespDTO struct {
	Login string `json:"login"`
	Token string `json:"token"`
}
