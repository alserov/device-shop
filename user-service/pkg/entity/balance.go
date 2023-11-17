package entity

type TopUpBalanceReq struct {
	Cash     float32 `json:"cash,omitempty" validate:"required,gt=0.0"`
	UserUUID string  `json:"userUUID,omitempty" validate:"required"`
}
