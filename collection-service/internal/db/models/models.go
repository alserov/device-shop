package models

type Device struct {
	UUID         string
	Title        string
	Description  string
	Price        float32
	Manufacturer string
	Amount       uint32
}

type DeviceFromCollection struct {
	Device Device `bson:"device"`
}
