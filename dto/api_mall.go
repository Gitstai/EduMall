package dto

type LoginRequest struct {
	Account  string `form:"account,required" json:"account,required" query:"account,required"`
	Password string `form:"password,required" json:"password,required" query:"password,required"`
	UserType *int32 `form:"userType,omitempty" json:"userType,omitempty" query:"userType,omitempty"`
}

type LoginResponse struct {
	UserId  int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	Code    *int32 `form:"code,omitempty" json:"code,omitempty" query:"code,omitempty"`
	Message *int32 `form:"message,omitempty" json:"message,omitempty" query:"message,omitempty"`
}

type RegisterRequest struct {
	Account  string  `form:"account,required" json:"account,required" query:"account,required"`
	Password string  `form:"password,required" json:"password,required" query:"password,required"`
	Nickname *string `form:"nickname,omitempty" json:"nickname,omitempty" query:"nickname,omitempty"`
}

type RegisterResponse struct {
	UserId  int64  `form:"userId,required" json:"userId,required" query:"userId,required"`
	Code    *int32 `form:"code,omitempty" json:"code,omitempty" query:"code,omitempty"`
	Message *int32 `form:"message,omitempty" json:"message,omitempty" query:"message,omitempty"`
}
