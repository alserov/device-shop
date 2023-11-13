package models

type AddToCollectionReq struct {
	DeviceUUID string `json:"deviceUUID,omitempty" validate:"required"`
	UserUUID   string `json:"userUUID,omitempty" validate:"required"`
}
