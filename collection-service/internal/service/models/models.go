package models

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type ChangeCollectionReq struct {
	UserUUID   string
	DeviceUUID string
}
