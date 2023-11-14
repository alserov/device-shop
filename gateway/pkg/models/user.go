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

type AddToCollectionReq struct {
	DeviceUUID string `json:"deviceUUID,omitempty" validate:"required"`
	UserUUID   string `json:"userUUID,omitempty" validate:"required"`
}
