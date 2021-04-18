package dto

type LoginRequest struct {
	Account  string `form:"account,required" json:"account,required" query:"account,required"`
	Password string `form:"password,required" json:"password,required" query:"password,required"`
}

type LoginResponse struct {
	UserId int64 `form:"userId,required" json:"userId,required" query:"userId,required"`
}

type RegisterRequest struct {
	Account  string  `form:"account,required" json:"account,required" query:"account,required"`
	Password string  `form:"password,required" json:"password,required" query:"password,required"`
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" query:"nickname,omitempty"`
}

type RegisterResponse struct {
	UserId int64 `form:"userId,required" json:"userId,required" query:"userId,required"`
}
