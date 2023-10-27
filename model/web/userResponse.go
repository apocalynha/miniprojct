package web

type UserReponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Role  string `json:"role" form:"role"`
}

type UserLoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
