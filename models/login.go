package models

type LoginStruct struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWTToken string `json:"jwt"`
}
