package entity

type Device struct {
	UUID         string  `json:"uuid,omitempty" bson:"uuid"`
	Title        string  `json:"title,omitempty" bson:"title" validate:"required,min=3"`
	Description  string  `json:"description,omitempty" bson:"description"`
	Price        float32 `json:"price,omitempty" bson:"price" validate:"required,gt=0"`
	Manufacturer string  `json:"manufacturer,omitempty" bson:"manufacturer" validate:"required"`
	Amount       uint32  `bson:"amount,omitempty" bson:"amount" validate:"gt=0"`
}
