package models

type LoginStruct struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWTToken string `json:"jwt"`
}

type SignupStruct struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type SignupResponse struct {
	JWTToken string `json:"jwt"`
}
