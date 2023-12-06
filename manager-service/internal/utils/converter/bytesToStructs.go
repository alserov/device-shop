package converter

import (
	"encoding/json"
	"github.com/alserov/device-shop/manager-service/internal/models"
)

type Converter struct {
}

func NewConverter() *Converter {
	return &Converter{}
}

func (*Converter) TxResponseToBytes(uuid string, status uint32, message string) []byte {
	bytes, _ := json.Marshal(models.TxResponse{
		Message: message,
		Status:  status,
		Uuid:    uuid,
	})

	return bytes
}
