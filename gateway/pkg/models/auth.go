package models

type SignupReq struct {
	Username string `json:"username,omitempty" validate:"min=3,max=50,required"`
	Password string `json:"password,omitempty" validate:"min=5,max=100,required"`
	Email    string `json:"email,omitempty" validate:"required"`
}

type LoginReq struct {
	Username string `json:"username,omitempty" validate:"min=3,max=50,required"`
	Password string `json:"password,omitempty" validate:"min=5,max=100,required"`
}
